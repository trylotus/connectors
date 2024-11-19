package pancakeswapv2

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/trylotus/connectors/pancakeswapv2/bep20"
	"github.com/trylotus/connectors/pancakeswapv2/factory"
	"github.com/trylotus/connectors/pancakeswapv2/pair"
	"github.com/trylotus/go-connector"
	"github.com/trylotus/go-connector/common"
	"github.com/trylotus/go-connector/log"
	"github.com/trylotus/go-connector/source/evm"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	blockRangeLimit             = 10000
	defaultQueryPageSize        = 2048
	defaultSubscriptionPageSize = 100000
)

var pairCreatedTopic = ethcommon.HexToHash("0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9")

type Source struct {
	client *evm.Client
	store  *Store

	queryPageSize        int64 // Limit number of pairs per query
	subscriptionPageSize int64 // Limit number of pairs per subscription

	factoryAddr ethcommon.Address
	pairs       PoolList

	factoryContract *factory.Factory

	pairCacheLock  *common.LockSet[ethcommon.Address]
	tokenCacheLock *common.LockSet[ethcommon.Address]
}

var _ connector.Source = (*Source)(nil)

func NewSource(client *evm.Client, store *Store, factoryAddr ethcommon.Address, opts ...Option) *Source {
	source := &Source{
		client:               client,
		store:                store,
		factoryAddr:          factoryAddr,
		queryPageSize:        defaultQueryPageSize,
		subscriptionPageSize: defaultSubscriptionPageSize,
		pairCacheLock:        common.NewLockSet[ethcommon.Address](),
		tokenCacheLock:       common.NewLockSet[ethcommon.Address](),
	}

	for _, opt := range opts {
		opt(source)
	}

	return source
}

func (s *Source) BlockNumber(ctx context.Context) (int64, error) {
	n, err := s.client.BlockNumber(ctx)
	if err != nil {
		return 0, err
	}

	return int64(n), nil
}

func (s *Source) Query(ctx context.Context, fromBlock int64, toBlock int64) ([]proto.Message, error) {
	msgs, err := s.queryFactory(ctx, fromBlock, toBlock, nil)
	if err != nil {
		return nil, err
	}

	pairAddrs := s.pairs.Search(toBlock)

	for i := 0; i < len(pairAddrs); i += int(s.queryPageSize) {
		j := i + int(s.queryPageSize)
		if j > len(pairAddrs) {
			j = len(pairAddrs)
		}

		msgs, err = s.queryPairs(ctx, fromBlock, toBlock, pairAddrs[i:j], msgs)
		if err != nil {
			return nil, err
		}

	}

	return msgs, nil
}

func (s *Source) queryFactory(ctx context.Context, fromBlock int64, toBlock int64, results []proto.Message) ([]proto.Message, error) {
	filter := ethereum.FilterQuery{
		Addresses: []ethcommon.Address{s.factoryAddr},
		FromBlock: big.NewInt(fromBlock),
	}

	if toBlock > 0 {
		filter.ToBlock = big.NewInt(toBlock)
	}

	logs, err := s.FilterLogs(ctx, filter)
	if err != nil {
		return nil, err
	}

	for _, vLog := range logs {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			if vLog.Removed {
				continue
			}

			event, err := s.factoryContract.ParsePairCreated(vLog)
			if err != nil {
				log.Error().Err(err).Str("tx", vLog.TxHash.String()).Uint("index", vLog.Index).Msg("Invalid factory log")
				continue
			}

			msg, err := s.ParsePairCreatedEvent(ctx, event)
			if err != nil {
				log.Error().Err(err).Str("tx", vLog.TxHash.String()).Uint("index", vLog.Index).Msgf("Invalid PairCreated event")
				continue
			}

			results = append(results, msg)
		}
	}

	return results, nil
}

func (s *Source) queryPairs(ctx context.Context, fromBlock int64, toBlock int64, pairs []ethcommon.Address, results []proto.Message) ([]proto.Message, error) {
	filter := ethereum.FilterQuery{
		Addresses: pairs,
		FromBlock: big.NewInt(fromBlock),
	}

	if toBlock > 0 {
		filter.ToBlock = big.NewInt(toBlock)
	}

	logs, err := s.FilterLogs(ctx, filter)
	if err != nil {
		return nil, err
	}

	for _, vLog := range logs {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			if vLog.Removed {
				continue
			}

			if bytes.Equal(vLog.Topics[0].Bytes(), pairCreatedTopic.Bytes()) {
				continue
			}

			msg, err := s.ParsePairLog(ctx, vLog)
			if err != nil {
				log.Error().Err(err).Str("tx", vLog.TxHash.String()).Uint("index", vLog.Index).Msg("Invalid pair log")
				continue
			}

			results = append(results, msg)
		}
	}

	return results, nil
}

func (s *Source) Subscribe(ctx context.Context, msgCh chan<- proto.Message, errCh chan<- error) {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		s.subscribeFactory(ctx, msgCh, errCh)
	}()

	pairAddrs := s.pairs.Addresses

	for i := 0; i < len(pairAddrs); i += int(s.subscriptionPageSize) {
		j := i + int(s.subscriptionPageSize)
		if j > len(pairAddrs) {
			j = len(pairAddrs)
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			s.subscribePairs(ctx, pairAddrs[i:j], msgCh, errCh)
		}()
	}

	wg.Wait()
}

func (s *Source) subscribeFactory(ctx context.Context, msgCh chan<- proto.Message, errCh chan<- error) {
	ch := make(chan *factory.FactoryPairCreated, 1)

	sub, err := s.factoryContract.WatchPairCreated(&bind.WatchOpts{Context: ctx}, ch, nil, nil)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to watch pair created")
	}

	defer sub.Unsubscribe()

	for {
		select {
		case err := <-sub.Err():
			log.Error().Err(err).Msg("Error while watching pair created")
			errCh <- err
			return
		case event := <-ch:
			if event.Raw.Removed {
				continue
			}

			log.Info().Str("number", event.Arg3.String()).Str("address", event.Pair.String()).Msg("New pair created")

			pair := &Pair{
				Number:  event.Arg3.Int64(),
				Address: event.Raw.Address,
				Token0:  event.Token0,
				Token1:  event.Token1,
			}
			if err := s.store.AddPair(ctx, pair); err != nil {
				log.Error().Err(err).Int64("number", pair.Number).Str("address", pair.Address.String()).Msg("Failed to add pair to store")
			}

			msg, err := s.ParsePairCreatedEvent(ctx, event)
			if err != nil {
				log.Error().Err(err).Str("tx", event.Raw.TxHash.String()).Uint("index", event.Raw.Index).Msg("Invalid PairCreated event")
				continue
			}

			msgCh <- msg

			// Prevent gaps
			go func() {
				subCtx, cancel := context.WithTimeout(ctx, 3*time.Minute)
				defer cancel()

				msgs, err := s.queryPairs(subCtx, int64(event.Raw.BlockNumber), 0, []ethcommon.Address{event.Pair}, nil)
				if err != nil {
					errCh <- err
					return
				}

				for _, msg := range msgs {
					msgCh <- msg
				}
			}()

			go s.subscribePairs(ctx, []ethcommon.Address{event.Pair}, msgCh, errCh)
		}
	}
}

func (s *Source) subscribePairs(ctx context.Context, pairs []ethcommon.Address, msgCh chan<- proto.Message, errCh chan<- error) {
	logCh := make(chan types.Log, 2048)

	sub, err := s.client.SubscribeFilterLogs(ctx, ethereum.FilterQuery{Addresses: pairs}, logCh)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to subscribe to pair contracts")
	}

	defer sub.Unsubscribe()

	for {
		select {
		case err := <-sub.Err():
			errCh <- err
			return
		case vLog := <-logCh:
			if vLog.Removed {
				continue
			}

			if bytes.Equal(vLog.Topics[0].Bytes(), pairCreatedTopic.Bytes()) {
				continue
			}

			msg, err := s.ParsePairLog(ctx, vLog)
			if err != nil {
				log.Error().Err(err).Str("tx", vLog.TxHash.String()).Uint("index", vLog.Index).Msg("Invalid pair log")
				continue
			}

			msgCh <- msg
		}
	}
}

func (s *Source) ParsePairCreatedEvent(ctx context.Context, event *factory.FactoryPairCreated) (proto.Message, error) {
	retryCtx := common.ContextWithFuncName(ctx, "ParsePairCreatedEvent")
	retryCtx = common.ContextWithOptionalRetry(retryCtx)

	return common.RetryT(retryCtx, func() (proto.Message, error) {
		subCtx, cancel := context.WithTimeout(ctx, 1*time.Minute)
		defer cancel()
		return s.parsePairCreatedEvent(subCtx, event)
	})
}

func (s *Source) parsePairCreatedEvent(ctx context.Context, event *factory.FactoryPairCreated) (proto.Message, error) {
	t, err := s.BlockTime(ctx, event.Raw.BlockHash)
	if err != nil {
		return nil, fmt.Errorf("failed to get timestamp: %w", err)
	}

	token0, err := s.GetToken(ctx, event.Token0)
	if err != nil {
		return nil, fmt.Errorf("failed to get token %s: %w", event.Token0, err)
	}

	token1, err := s.GetToken(ctx, event.Token1)
	if err != nil {
		return nil, fmt.Errorf("failed to get token %s: %w", event.Token1, err)

	}

	return &factory.PairCreated{
		Ts:             &timestamppb.Timestamp{Seconds: int64(t)},
		BlockNumber:    event.Raw.BlockNumber,
		BlockHash:      event.Raw.BlockHash.Bytes(),
		TxHash:         event.Raw.TxHash.Bytes(),
		LogIndex:       uint64(event.Raw.Index),
		Token0:         event.Token0.Bytes(),
		Token1:         event.Token1.Bytes(),
		Pair:           event.Pair.Bytes(),
		Arg3:           event.Arg3.String(),
		Token0Name:     token0.Name,
		Token0Symbol:   token0.Symbol,
		Token0Decimals: uint32(token0.Decimals),
		Token1Name:     token1.Name,
		Token1Symbol:   token1.Symbol,
		Token1Decimals: uint32(token1.Decimals),
	}, nil
}

func (s *Source) ParsePairLog(ctx context.Context, vLog types.Log) (proto.Message, error) {
	retryCtx := common.ContextWithFuncName(ctx, "ParsePairLog")
	retryCtx = common.ContextWithOptionalRetry(retryCtx)

	return common.RetryT(retryCtx, func() (proto.Message, error) {
		subCtx, cancel := context.WithTimeout(ctx, 1*time.Minute)
		defer cancel()
		return s.parsePairLog(subCtx, vLog)
	})
}

func (s *Source) parsePairLog(ctx context.Context, vLog types.Log) (proto.Message, error) {
	t, err := s.BlockTime(ctx, vLog.BlockHash)
	if err != nil {
		return nil, fmt.Errorf("failed to get timestamp: %w", err)
	}

	ts := &timestamppb.Timestamp{Seconds: int64(t)}

	event, err := pair.UnpackLog(vLog)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack pair log: %w", err)
	}

	p, err := s.GetPair(ctx, vLog.Address)
	if err != nil {
		return nil, fmt.Errorf("failed to get pair %s: %w", vLog.Address, err)
	}

	switch event := event.(type) {
	case pair.PairMint:
		token0, err := s.GetToken(ctx, p.Token0)
		if err != nil {
			return nil, err
		}

		token1, err := s.GetToken(ctx, p.Token1)
		if err != nil {
			return nil, err
		}

		amount0 := tokenAmount(event.Amount0, token0.Decimals)
		amount1 := tokenAmount(event.Amount1, token1.Decimals)

		return &pair.Mint{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			BlockHash:   vLog.BlockHash.Bytes(),
			TxHash:      vLog.TxHash.Bytes(),
			LogIndex:    uint64(vLog.Index),
			Pair:        vLog.Address.Bytes(),
			Sender:      event.Sender.Bytes(),
			Amount0:     floatString(amount0, token0.Decimals),
			Amount1:     floatString(amount1, token1.Decimals),
		}, nil
	case pair.PairSwap:
		token0, err := s.GetToken(ctx, p.Token0)
		if err != nil {
			return nil, err
		}

		token1, err := s.GetToken(ctx, p.Token1)
		if err != nil {
			return nil, err
		}

		amount0In := tokenAmount(event.Amount0In, token0.Decimals)
		amount0Out := tokenAmount(event.Amount0Out, token0.Decimals)
		amount1In := tokenAmount(event.Amount1In, token1.Decimals)
		amount1Out := tokenAmount(event.Amount1Out, token1.Decimals)

		return &pair.Swap{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			BlockHash:   vLog.BlockHash.Bytes(),
			TxHash:      vLog.TxHash.Bytes(),
			LogIndex:    uint64(vLog.Index),
			Pair:        vLog.Address.Bytes(),
			Amount0In:   floatString(amount0In, token0.Decimals),
			Amount0Out:  floatString(amount0Out, token0.Decimals),
			Amount1In:   floatString(amount1In, token1.Decimals),
			Amount1Out:  floatString(amount1Out, token1.Decimals),
			Sender:      event.Sender.Bytes(),
			To:          event.To.Bytes(),
		}, nil
	case pair.PairSync:
		token0, err := s.GetToken(ctx, p.Token0)
		if err != nil {
			return nil, err
		}

		token1, err := s.GetToken(ctx, p.Token1)
		if err != nil {
			return nil, err
		}

		reserve0 := tokenAmount(event.Reserve0, token0.Decimals)
		reserve1 := tokenAmount(event.Reserve1, token1.Decimals)

		return &pair.Sync{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			BlockHash:   vLog.BlockHash.Bytes(),
			TxHash:      vLog.TxHash.Bytes(),
			LogIndex:    uint64(vLog.Index),
			Pair:        vLog.Address.Bytes(),
			Reserve0:    floatString(reserve0, token0.Decimals),
			Reserve1:    floatString(reserve1, token1.Decimals),
		}, nil
	case pair.PairTransfer:
		// PancakeSwapV2 LP token decimals is always 18
		value := tokenAmount(event.Value, 18)

		return &pair.Transfer{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			BlockHash:   vLog.BlockHash.Bytes(),
			TxHash:      vLog.TxHash.Bytes(),
			LogIndex:    uint64(vLog.Index),
			Pair:        vLog.Address.Bytes(),
			From:        event.From.Bytes(),
			To:          event.To.Bytes(),
			Value:       floatString(value, 18),
		}, nil
	case pair.PairApproval:
		// PancakeSwapV2 LP token decimals is always 18
		value := tokenAmount(event.Value, 18)

		return &pair.Approval{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			BlockHash:   vLog.BlockHash.Bytes(),
			TxHash:      vLog.TxHash.Bytes(),
			LogIndex:    uint64(vLog.Index),
			Pair:        vLog.Address.Bytes(),
			Owner:       event.Owner.Bytes(),
			Spender:     event.Spender.Bytes(),
			Value:       floatString(value, 18),
		}, nil
	case pair.PairBurn:
		token0, err := s.GetToken(ctx, p.Token0)
		if err != nil {
			return nil, err
		}

		token1, err := s.GetToken(ctx, p.Token1)
		if err != nil {
			return nil, err
		}

		amount0 := tokenAmount(event.Amount0, token0.Decimals)
		amount1 := tokenAmount(event.Amount1, token1.Decimals)

		return &pair.Burn{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			BlockHash:   vLog.BlockHash.Bytes(),
			TxHash:      vLog.TxHash.Bytes(),
			LogIndex:    uint64(vLog.Index),
			Pair:        vLog.Address.Bytes(),
			Sender:      event.Sender.Bytes(),
			Amount0:     floatString(amount0, token0.Decimals),
			Amount1:     floatString(amount1, token1.Decimals),
			To:          event.To.Bytes(),
		}, nil
	default:
		return nil, fmt.Errorf("unhandled event: %s", reflect.TypeOf(event))
	}
}

func (s *Source) BlockTime(ctx context.Context, hash ethcommon.Hash) (uint64, error) {
	retryCtx := common.ContextWithFuncName(ctx, "BlockTime")
	retryCtx = common.ContextWithOptionalRetry(retryCtx)

	return common.RetryT(retryCtx, func() (uint64, error) {
		subCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		return s.client.BlockTime(subCtx, hash)
	})
}

func (s *Source) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	retryCtx := common.ContextWithFuncName(ctx, "FilterLogs")
	retryCtx = common.ContextWithOptionalRetry(retryCtx)

	return common.RetryT(retryCtx, func() ([]types.Log, error) {
		subCtx, cancel := context.WithTimeout(ctx, 3*time.Minute)
		defer cancel()
		return s.client.FilterLogs(subCtx, q)
	})
}

func (s *Source) GetToken(ctx context.Context, address ethcommon.Address) (*Token, error) {
	s.tokenCacheLock.Lock(address)
	defer s.tokenCacheLock.Unlock(address)

	token, err := s.store.GetToken(ctx, address)
	if err != nil {
		return nil, err
	}
	if token != nil {
		return token, nil
	}

	retryCtx := common.ContextWithFuncName(ctx, "GetTokenFromRpc")
	retryCtx = common.ContextWithOptionalRetry(retryCtx)

	token, err = common.RetryT(retryCtx, func() (*Token, error) {
		subCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		return s.getTokenFromRpc(subCtx, address)
	})
	if err != nil {
		return nil, err
	}

	if err := s.store.AddToken(ctx, token); err != nil {
		log.Error().Err(err).Str("address", token.Address.String()).Msg("Failed to add token to store")
	}

	return token, nil
}

func (s *Source) getTokenFromRpc(ctx context.Context, address ethcommon.Address) (*Token, error) {
	tokenContract, err := bep20.NewBep20(address, s.client)
	if err != nil {
		return nil, err
	}

	decimals, err := tokenContract.Decimals(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("failed to get token decimals: %w", err)
	}

	name, err := tokenContract.Name(&bind.CallOpts{Context: ctx})
	if err != nil {
		if strings.HasPrefix(err.Error(), "abi") { // Decode error
			log.Error().Err(err).Str("address", address.String()).Msg("Failed to get token name")
		} else {
			return nil, fmt.Errorf("failed to get token name: %w", err)
		}
	}

	symbol, err := tokenContract.Symbol(&bind.CallOpts{Context: ctx})
	if err != nil {
		if strings.HasPrefix(err.Error(), "abi") { // Decode error
			log.Error().Err(err).Str("address", address.String()).Msg("Failed to get token symbol")
		} else {
			return nil, fmt.Errorf("failed to get token symbol: %w", err)
		}
	}

	return &Token{
		Address:  address,
		Name:     name,
		Symbol:   symbol,
		Decimals: decimals.Int64(),
	}, nil
}

func (s *Source) GetPair(ctx context.Context, address ethcommon.Address) (*Pair, error) {
	s.pairCacheLock.Lock(address)
	defer s.pairCacheLock.Unlock(address)

	p, err := s.store.GetPair(ctx, address)
	if err != nil {
		return nil, err
	}
	if p != nil {
		return p, nil
	}

	retryCtx := common.ContextWithFuncName(ctx, "GetPairFromRpc")
	retryCtx = common.ContextWithOptionalRetry(retryCtx)

	p, err = common.RetryT(retryCtx, func() (*Pair, error) {
		subCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		return s.getPairFromRpc(subCtx, address)
	})
	if err != nil {
		return nil, err
	}

	if err := s.store.AddPair(ctx, p); err != nil {
		log.Error().Err(err).Str("address", p.Address.String()).Msg("Failed to add pair to store")
	}

	return p, nil
}

func (s *Source) getPairFromRpc(ctx context.Context, address ethcommon.Address) (*Pair, error) {
	pairContract, err := pair.NewPair(address, s.client)
	if err != nil {
		return nil, err
	}

	token0, err := pairContract.Token0(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	token1, err := pairContract.Token1(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	p := &Pair{
		Address: address,
		Token0:  token0,
		Token1:  token1,
	}

	return p, nil
}

func (s *Source) Init(ctx context.Context) {
	factoryContract, err := factory.NewFactory(s.factoryAddr, s.client)
	if err != nil {
		log.Fatal().Err(err).Str("contract", s.factoryAddr.String()).Msg("Failed to create factory contract")
	}

	s.factoryContract = factoryContract

	log.Info().Msg("Loading all pairs")

	s.loadPairsFromStore(ctx)

	scannedBlock, err := s.store.GetScannedBlock(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get scanned block number")
	}

	blockNumber, err := s.BlockNumber(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get block number")
	}

	for i := scannedBlock + 1; i <= blockNumber; i += blockRangeLimit {
		j := i + blockRangeLimit - 1
		if j > blockNumber {
			j = blockNumber
		}

		s.loadPairsFromRPC(ctx, uint64(i), uint64(j))
	}

	s.pairs.SortAndRemoveDuplicates()

	log.Info().Int("total", s.pairs.Len()).Msg("Loaded all pairs")
}

func (s *Source) loadPairsFromStore(ctx context.Context) {
	pairCh, errCh := s.store.AllPairs(ctx)

	for pair := range pairCh {
		s.pairs.Add(ethcommon.HexToAddress(pair.Address.String()), pair.BlockNumber)
	}

	for err := range errCh {
		log.Fatal().Err(err).Msg("Failed to load pairs from store")
	}

	log.Info().Int("total", s.pairs.Len()).Msg("Loaded pairs from store")
}

func (s *Source) loadPairsFromRPC(ctx context.Context, from uint64, to uint64) {
	opts := &bind.FilterOpts{
		Context: ctx,
		Start:   from,
		End:     &to,
	}

	var successCount, failedCount int

	it, err := s.factoryContract.FilterPairCreated(opts, nil, nil)
	if err != nil {
		log.Fatal().Err(err).Uint64("from", from).Uint64("to", to).Msg("Failed to load pairs from RPC")
	}

	defer it.Close()

	for it.Next() {
		if err := it.Error(); err != nil {
			log.Fatal().Err(err).Uint64("from", from).Uint64("to", to).Msg("Failed to load pairs from RPC")
		}

		if it.Event.Raw.Removed {
			continue
		}

		if _, err := s.GetToken(ctx, it.Event.Token0); err != nil {
			failedCount++
			log.Error().Err(err).Str("address", it.Event.Token0.String()).Msg("Failed to get token")
			continue
		}

		if _, err := s.GetToken(ctx, it.Event.Token1); err != nil {
			failedCount++
			log.Error().Err(err).Str("address", it.Event.Token1.String()).Msg("Failed to get token")
			continue
		}

		successCount++

		s.pairs.Add(it.Event.Pair, int64(it.Event.Raw.BlockNumber))

		pair := Pair{
			Number:      it.Event.Arg3.Int64(),
			Address:     it.Event.Pair,
			Token0:      it.Event.Token0,
			Token1:      it.Event.Token1,
			BlockNumber: int64(it.Event.Raw.BlockNumber),
		}

		if err := s.store.AddPair(ctx, &pair); err != nil {
			log.Error().Err(err).Str("address", pair.Address.String()).Msg("Failed to add pair to store")
		}
	}

	log.Info().Uint64("from", from).Uint64("to", to).Int("success", successCount).Int("failed", failedCount).Msg("Loaded pairs from RPC")

	if err := s.store.SetScannedBlock(ctx, int64(to)); err != nil {
		log.Error().Err(err).Uint64("number", to).Msg("Failed to set scanned block")
	}
}

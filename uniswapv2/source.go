package uniswapv2

import (
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
	"github.com/trylotus/connectors/uniswapv2/erc20"
	"github.com/trylotus/connectors/uniswapv2/factory"
	"github.com/trylotus/connectors/uniswapv2/pair"
	"github.com/trylotus/go-connector"
	"github.com/trylotus/go-connector/common"
	"github.com/trylotus/go-connector/log"
	"github.com/trylotus/go-connector/source/evm"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	defaultConcurreny           = 10
	defaultQueryPageSize        = 2048
	defaultSubscriptionPageSize = 100000
)

type Source struct {
	client *evm.Client
	store  *Store

	concurrecy           int64 // Limit mumber of queries can be run concurrently
	queryPageSize        int64 // Limit number of pairs per query
	subscriptionPageSize int64 // Limit number of pairs per subscription

	factoryContract     *factory.Factory
	factoryContractAddr ethcommon.Address
	pairs               []ethcommon.Address
	loadAllPairsOnce    sync.Once
}

var _ connector.Source = (*Source)(nil)

func NewSource(client *evm.Client, store *Store, factoryContractAddr string, opts ...Option) connector.Source {
	factoryContract, err := factory.NewFactory(ethcommon.HexToAddress(factoryContractAddr), client)
	if err != nil {
		log.Fatal().Err(err).Str("contract", factoryContractAddr).Msg("Failed to create factory contract")
	}

	source := &Source{
		client:               client,
		store:                store,
		factoryContract:      factoryContract,
		factoryContractAddr:  ethcommon.HexToAddress(factoryContractAddr),
		concurrecy:           defaultConcurreny,
		queryPageSize:        defaultQueryPageSize,
		subscriptionPageSize: defaultSubscriptionPageSize,
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
	var results []proto.Message

	msgs, err := s.queryFactory(ctx, fromBlock, toBlock)
	if err != nil {
		return nil, err
	}

	results = append(results, msgs...)

	pairs := s.AllPairs(ctx)

	for i := 0; i < len(pairs); i += int(s.queryPageSize) {
		j := i + int(s.queryPageSize)
		if j > len(pairs) {
			j = len(pairs)
		}

		msgs, err := s.queryPairs(ctx, fromBlock, toBlock, pairs[i:j])
		if err != nil {
			return nil, err
		}

		results = append(results, msgs...)
	}

	return results, nil
}

func (s *Source) queryFactory(ctx context.Context, fromBlock int64, toBlock int64) ([]proto.Message, error) {
	filter := ethereum.FilterQuery{
		Addresses: []ethcommon.Address{s.factoryContractAddr},
		FromBlock: big.NewInt(fromBlock),
	}

	if toBlock > 0 {
		filter.ToBlock = big.NewInt(toBlock)
	}

	logs, err := s.FilterLogs(ctx, filter)
	if err != nil {
		return nil, err
	}

	msgs := make([]proto.Message, 0, len(logs))

	for _, vLog := range logs {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			event, err := s.factoryContract.ParsePairCreated(vLog)
			if err != nil {
				log.Error().Err(err).Msg("Invalid factory log")
				continue
			}

			msg, err := s.ParsePairCreatedEvent(ctx, event)
			if err != nil {
				log.Error().Err(err).Msg("Invalid PairCreated event")
				continue
			}

			msgs = append(msgs, msg)
		}
	}

	return msgs, nil
}

func (s *Source) queryPairs(ctx context.Context, fromBlock int64, toBlock int64, pairs []ethcommon.Address) ([]proto.Message, error) {
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

	msgs := make([]proto.Message, 0, len(logs))

	for _, vLog := range logs {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			msg, err := s.ParsePairLog(ctx, vLog)
			if err != nil {
				log.Error().Err(err).Msg("Invalid pair log")
				continue
			}

			msgs = append(msgs, msg)
		}
	}

	return msgs, nil
}

func (s *Source) Subscribe(ctx context.Context, msgCh chan<- proto.Message, errCh chan<- error) {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		s.subscribeFactory(ctx, msgCh, errCh)
	}()

	pairs := s.AllPairs(ctx)

	for i := 0; i < len(pairs); i += int(s.subscriptionPageSize) {
		j := i + int(s.subscriptionPageSize)
		if j > len(pairs) {
			j = len(pairs)
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			s.subscribePairs(ctx, pairs[i:j], msgCh, errCh)
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
			log.Info().Str("number", event.Arg3.String()).Str("address", event.Pair.String()).Msg("New pair created")

			pair := &Pair{
				Number:  event.Arg3.Int64(),
				Address: event.Raw.Address.String(),
				Token0:  event.Token0.String(),
				Token1:  event.Token1.String(),
			}
			if err := s.store.AddPair(ctx, pair); err != nil {
				log.Error().Err(err).Int64("number", pair.Number).Str("address", pair.Address).Msg("Failed to add pair to store")
			}

			// Prevent gaps
			go func() {
				subCtx, cancel := context.WithTimeout(ctx, 3*time.Minute)
				defer cancel()

				msgs, err := s.queryPairs(subCtx, int64(event.Raw.BlockNumber), 0, []ethcommon.Address{event.Pair})
				if err != nil {
					errCh <- err
					return
				}

				for _, msg := range msgs {
					msgCh <- msg
				}
			}()

			go s.subscribePairs(ctx, []ethcommon.Address{event.Pair}, msgCh, errCh)

			msg, err := s.ParsePairCreatedEvent(ctx, event)
			if err != nil {
				log.Error().Err(err).Msg("Invalid PairCreated event")
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
	t, err := s.client.BlockTime(ctx, event.Raw.BlockNumber)
	if err != nil {
		return nil, fmt.Errorf("error retrieving timestamp: %w", err)
	}

	token0, err := s.GetToken(ctx, event.Token0)
	if err != nil {
		if err.Error() != "no contract code at given address" {
			return nil, fmt.Errorf("error getting token %s from RPC: %w", event.Token0, err)
		}
	}

	token1, err := s.GetToken(ctx, event.Token1)
	if err != nil {
		if err.Error() != "no contract code at given address" {
			return nil, fmt.Errorf("error getting token %s from RPC: %w", event.Token1, err)
		}
	}

	msg := &factory.PairCreated{
		Ts:          &timestamppb.Timestamp{Seconds: int64(t)},
		BlockNumber: event.Raw.BlockNumber,
		TxHash:      event.Raw.TxHash.Bytes(),
		LogIndex:    uint64(event.Raw.Index),
		Token0:      event.Token0.Bytes(),
		Token1:      event.Token1.Bytes(),
		Pair:        event.Pair.Bytes(),
		Arg3:        event.Arg3.String(),
	}

	if token0 != nil {
		msg.Token0Name = token0.Name
		msg.Token0Symbol = token0.Symbol
		msg.Token0Decimals = uint32(token0.Decimals)
	}

	if token1 != nil {
		msg.Token1Name = token1.Name
		msg.Token1Symbol = token1.Symbol
		msg.Token1Decimals = uint32(token1.Decimals)
	}

	return msg, nil
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
			msg, err := s.ParsePairLog(ctx, vLog)
			if err != nil {
				log.Error().Err(err).Msg("Invalid pair log")
				continue
			}

			msgCh <- msg
		}
	}
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
	t, err := s.client.BlockTime(ctx, vLog.BlockNumber)
	if err != nil {
		return nil, fmt.Errorf("error retrieving timestamp: %w", err)
	}

	ts := &timestamppb.Timestamp{Seconds: int64(t)}

	event, err := pair.UnpackLog(vLog)
	if err != nil {
		return nil, fmt.Errorf("error unpacking pair log: %w", err)
	}

	p, err := s.store.GetPair(ctx, vLog.Address.String())
	if err != nil {
		log.Warn().Err(err).Str("address", vLog.Address.String()).Msg("Failed to get pair from store")
	}

	if p == nil {
		p, err = s.GetPair(ctx, vLog.Address)
		if err != nil {
			return nil, fmt.Errorf("error getting pair %s from RPC: %w", vLog.Address, err)
		}
	}

	switch event := event.(type) {
	case pair.PairMint:
		token0, err := s.GetToken(ctx, ethcommon.HexToAddress(p.Token0))
		if err != nil {
			return nil, err
		}

		token1, err := s.GetToken(ctx, ethcommon.HexToAddress(p.Token0))
		if err != nil {
			return nil, err
		}

		amount0 := tokenAmount(event.Amount0, token0.Decimals)
		amount1 := tokenAmount(event.Amount1, token1.Decimals)

		return &pair.Mint{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			TxHash:      vLog.TxHash.Bytes(),
			LogIndex:    uint64(vLog.Index),
			Pair:        vLog.Address.Bytes(),
			Sender:      event.Sender.Bytes(),
			Amount0:     floatString(amount0),
			Amount1:     floatString(amount1),
		}, nil
	case pair.PairSwap:
		token0, err := s.GetToken(ctx, ethcommon.HexToAddress(p.Token0))
		if err != nil {
			return nil, err
		}

		token1, err := s.GetToken(ctx, ethcommon.HexToAddress(p.Token0))
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
			TxHash:      vLog.TxHash.Bytes(),
			LogIndex:    uint64(vLog.Index),
			Pair:        vLog.Address.Bytes(),
			Amount0In:   floatString(amount0In),
			Amount0Out:  floatString(amount0Out),
			Amount1In:   floatString(amount1In),
			Amount1Out:  floatString(amount1Out),
			Sender:      event.Sender.Bytes(),
			To:          event.To.Bytes(),
		}, nil
	case pair.PairSync:
		token0, err := s.GetToken(ctx, ethcommon.HexToAddress(p.Token0))
		if err != nil {
			return nil, err
		}

		token1, err := s.GetToken(ctx, ethcommon.HexToAddress(p.Token0))
		if err != nil {
			return nil, err
		}

		reserve0 := tokenAmount(event.Reserve0, token0.Decimals)
		reserve1 := tokenAmount(event.Reserve1, token1.Decimals)

		return &pair.Sync{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			TxHash:      vLog.TxHash.Bytes(),
			LogIndex:    uint64(vLog.Index),
			Pair:        vLog.Address.Bytes(),
			Reserve0:    floatString(reserve0),
			Reserve1:    floatString(reserve1),
		}, nil
	case pair.PairTransfer:
		// Uniswapv2 LP token decimals is always 18
		value := tokenAmount(event.Value, 18)

		return &pair.Transfer{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			TxHash:      vLog.TxHash.Bytes(),
			LogIndex:    uint64(vLog.Index),
			Pair:        vLog.Address.Bytes(),
			From:        event.From.Bytes(),
			To:          event.To.Bytes(),
			Value:       floatString(value),
		}, nil
	case pair.PairApproval:
		// Uniswapv2 LP token decimals is always 18
		value := tokenAmount(event.Value, 18)

		return &pair.Approval{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			TxHash:      vLog.TxHash.Bytes(),
			LogIndex:    uint64(vLog.Index),
			Pair:        vLog.Address.Bytes(),
			Owner:       event.Owner.Bytes(),
			Spender:     event.Spender.Bytes(),
			Value:       floatString(value),
		}, nil
	case pair.PairBurn:
		token0, err := s.GetToken(ctx, ethcommon.HexToAddress(p.Token0))
		if err != nil {
			return nil, err
		}

		token1, err := s.GetToken(ctx, ethcommon.HexToAddress(p.Token0))
		if err != nil {
			return nil, err
		}

		amount0 := tokenAmount(event.Amount0, token0.Decimals)
		amount1 := tokenAmount(event.Amount1, token1.Decimals)

		return &pair.Burn{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			TxHash:      vLog.TxHash.Bytes(),
			LogIndex:    uint64(vLog.Index),
			Pair:        vLog.Address.Bytes(),
			Sender:      event.Sender.Bytes(),
			Amount0:     floatString(amount0),
			Amount1:     floatString(amount1),
			To:          event.To.Bytes(),
		}, nil
	default:
		return nil, fmt.Errorf("unhandled event: %s", reflect.TypeOf(event))
	}
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

func (s *Source) AllPairs(ctx context.Context) []ethcommon.Address {
	s.loadAllPairsOnce.Do(func() {
		s.loadAllPairs(ctx)
	})

	return s.pairs
}

func (s *Source) GetToken(ctx context.Context, address ethcommon.Address) (*Token, error) {
	token, err := s.store.GetToken(ctx, address.String())
	if err != nil {
		return nil, err
	}
	if token != nil {
		return token, nil
	}

	retryCtx := common.ContextWithConditionalRetry(ctx, evm.IsRetryableError)
	retryCtx = common.ContextWithFuncName(retryCtx, "GetToken")

	token, err = common.RetryT(retryCtx, func() (*Token, error) {
		subCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		return s.getToken(subCtx, address)
	})
	if err != nil {
		return nil, err
	}

	if err := s.store.AddToken(ctx, token); err != nil {
		log.Error().Err(err).Str("address", token.Address).Msg("Failed to add token to store")
	}

	return token, nil
}

func (s *Source) getToken(ctx context.Context, address ethcommon.Address) (*Token, error) {
	tokenContract, err := erc20.NewErc20(address, s.client)
	if err != nil {
		return nil, err
	}

	name, err := tokenContract.Name(&bind.CallOpts{Context: ctx})
	if err != nil {
		log.Error().Err(err).Str("address", address.String()).Msg("Failed to get token name")
	}

	symbol, err := tokenContract.Symbol(&bind.CallOpts{Context: ctx})
	if err != nil {
		log.Error().Err(err).Str("address", address.String()).Msg("Failed to get token symbol")
	}

	decimals, err := tokenContract.Decimals(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	return &Token{
		Address:  address.String(),
		Name:     name,
		Symbol:   symbol,
		Decimals: decimals,
	}, nil
}

func (s *Source) GetPair(ctx context.Context, address ethcommon.Address) (*Pair, error) {
	retryCtx := common.ContextWithConditionalRetry(ctx, evm.IsRetryableError)
	retryCtx = common.ContextWithFuncName(retryCtx, "GetPair")

	return common.RetryT(retryCtx, func() (*Pair, error) {
		subCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		return s.getPair(subCtx, address)
	})
}

func (s *Source) getPair(ctx context.Context, address ethcommon.Address) (*Pair, error) {
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
		Address: address.String(),
		Token0:  token0.String(),
		Token1:  token1.String(),
	}

	return p, nil
}

func (s *Source) GetPairByNumber(ctx context.Context, number *big.Int) (*Pair, error) {
	retryCtx := common.ContextWithConditionalRetry(ctx, evm.IsRetryableError)
	retryCtx = common.ContextWithFuncName(retryCtx, "GetPairByNumber")

	address, err := common.RetryT(retryCtx, func() (ethcommon.Address, error) {
		subCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		return s.factoryContract.AllPairs(&bind.CallOpts{Context: subCtx}, number)
	})
	if err != nil {
		return nil, err
	}

	pair, err := s.GetPair(ctx, address)
	if err != nil {
		return nil, err
	}

	pair.Number = number.Int64()

	return pair, nil
}

func (s *Source) GetPairs(ctx context.Context, from int64, to int64) <-chan *Pair {
	pairCh := make(chan *Pair, 100)

	go func() {
		defer close(pairCh)

		var wg sync.WaitGroup

		sem := make(chan struct{}, s.concurrecy)

		for i := from; i <= to; i++ {
			sem <- struct{}{}
			wg.Add(1)

			go func(number int64) {
				defer wg.Done()
				defer func() { <-sem }()

				pair, err := s.GetPairByNumber(ctx, big.NewInt(number))
				if err != nil {
					log.Fatal().Err(err).Int64("number", number).Msg("Failed to load pair from RPC")
				}

				pairCh <- pair
			}(i)
		}

		wg.Wait()
	}()

	return pairCh
}

func (s *Source) loadAllPairs(ctx context.Context) {
	pairCount, err := s.factoryContract.AllPairsLength(&bind.CallOpts{Context: ctx})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to query all pairs length")
	}

	log.Info().Int64("total", pairCount.Int64()).Msg("Loading all pairs")

	s.loadPairsFromStore(ctx)
	s.loadPairsFromRPC(ctx, int64(len(s.pairs)), pairCount.Int64()-1)

	log.Info().Int("total", len(s.pairs)).Msg("Successfully loaded all pairs")
}

func (s *Source) loadPairsFromStore(ctx context.Context) {
	log.Info().Msg("Loading pairs from store")

	pairCh, errCh := s.store.AllPairs(ctx)

	var lastPairNumber int64 = -1

	for pair := range pairCh {
		if pair.Number-lastPairNumber > 1 {
			// Filling the gap
			s.loadPairsFromRPC(ctx, lastPairNumber+1, pair.Number-1)
		}

		s.pairs = append(s.pairs, ethcommon.HexToAddress(pair.Address))

		lastPairNumber = pair.Number

		log.Debug().Int64("number", pair.Number).Str("address", pair.Address).Msg("Loaded pair from store")
	}

	for err := range errCh {
		log.Fatal().Err(err).Msg("Failed to load pair from store")
	}

	log.Info().Msg("Successfully loaded pairs from store")
}

func (s *Source) loadPairsFromRPC(ctx context.Context, from int64, to int64) {
	log.Info().Int64("from", from).Int64("to", to).Msg("Loading pairs from RPC")

	pairCh := s.GetPairs(ctx, from, to)

	for pair := range pairCh {
		s.pairs = append(s.pairs, ethcommon.HexToAddress(pair.Address))

		if err := s.store.AddPair(ctx, pair); err != nil {
			log.Error().Err(err).Int64("number", pair.Number).Str("address", pair.Address).Msg("Failed to add pair to store")
		}

		log.Debug().Int64("number", pair.Number).Str("address", pair.Address).Msg("Loaded pair from RPC")
	}

	log.Info().Int64("from", from).Int64("to", to).Msg("Successfully loaded pairs from RPC")
}

func floatString(n *big.Rat) string {
	s := n.FloatString(18)
	s = strings.TrimRight(s, "0")
	return strings.TrimSuffix(s, ".")
}

func tokenAmount(rawAmount *big.Int, decimals uint8) *big.Rat {
	return new(big.Rat).SetFrac(rawAmount, new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil))
}

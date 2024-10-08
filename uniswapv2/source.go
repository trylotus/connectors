package uniswapv2

import (
	"context"
	"fmt"
	"math/big"
	"reflect"
	"strings"
	"sync"

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
	defaultConcurrenyLimit      = 5
	defaultBlockRangeLimit      = 2048
	defaultQueryPageSize        = 2048
	defaultSubscriptionPageSize = 100000
)

type Source struct {
	client *evm.Client
	store  *Store

	concurrecyLimit      int64 // Limit mumber of queries can be run concurrently
	blockRangeLimit      int64 // Limit number of blocks per query
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
		concurrecyLimit:      defaultConcurrenyLimit,
		blockRangeLimit:      defaultBlockRangeLimit,
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

func (s *Source) Query(ctx context.Context, fromBlock int64, toBlock int64, msgCh chan<- proto.Message, errCh chan<- error) {
	if fromBlock > toBlock {
		fromBlock, toBlock = toBlock, fromBlock
	}

	pairs := s.AllPairs(ctx)

	sem := make(chan struct{}, s.concurrecyLimit)

	for chunkHead := fromBlock; chunkHead <= toBlock; chunkHead += s.blockRangeLimit {
		chunkTail := chunkHead + s.blockRangeLimit - 1
		if chunkTail > toBlock {
			chunkTail = toBlock
		}

		log.Info().Int64("from", chunkHead).Int64("to", chunkTail).Msg("Query")

		var wg sync.WaitGroup

		sem <- struct{}{}
		wg.Add(1)

		go func(from int64, to int64) {
			defer wg.Done()
			defer func() { <-sem }()

			s.queryFactory(ctx, from, to, msgCh, errCh)
		}(chunkHead, chunkTail)

		for i := 0; i < len(pairs); i += int(s.queryPageSize) {
			j := i + int(s.queryPageSize)
			if j > len(pairs) {
				j = len(pairs)
			}

			sem <- struct{}{}
			wg.Add(1)

			go func(from int64, to int64, pairs []ethcommon.Address) {
				defer wg.Done()
				defer func() { <-sem }()

				s.queryPairs(ctx, from, to, pairs, msgCh, errCh)
			}(chunkHead, chunkTail, pairs[i:j])
		}

		wg.Wait()
	}
}

func (s *Source) queryFactory(ctx context.Context, fromBlock int64, toBlock int64, msgCh chan<- proto.Message, errCh chan<- error) {
	filter := ethereum.FilterQuery{
		Addresses: []ethcommon.Address{s.factoryContractAddr},
		FromBlock: big.NewInt(fromBlock),
	}

	if toBlock > 0 {
		filter.ToBlock = big.NewInt(toBlock)
	}

	retryCtx := common.ContextWithConditionalRetry(ctx, evm.IsRetryableError)
	retryCtx = common.ContextWithFuncName(retryCtx, "FilterLogs")

	logs, err := common.RetryT(retryCtx, func() ([]types.Log, error) {
		return s.client.FilterLogs(ctx, filter)
	})

	if err != nil {
		errCh <- err
		return
	}

	for _, vLog := range logs {
		event, err := s.factoryContract.ParsePairCreated(vLog)
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse factory log")
			continue
		}

		msg, err := s.parsePairCreatedEvent(ctx, event)
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse factory pair created event")
			continue
		}

		msgCh <- msg
	}
}

func (s *Source) queryPairs(ctx context.Context, fromBlock int64, toBlock int64, pairs []ethcommon.Address, msgCh chan<- proto.Message, errCh chan<- error) {

	filter := ethereum.FilterQuery{
		Addresses: pairs,
		FromBlock: big.NewInt(fromBlock),
	}

	if toBlock > 0 {
		filter.ToBlock = big.NewInt(toBlock)
	}

	retryCtx := common.ContextWithConditionalRetry(ctx, evm.IsRetryableError)
	retryCtx = common.ContextWithFuncName(retryCtx, "FilterLogs")

	logs, err := common.RetryT(retryCtx, func() ([]types.Log, error) {
		return s.client.FilterLogs(ctx, filter)
	})

	if err != nil {
		errCh <- err
		return
	}

	for _, vLog := range logs {
		msg, err := s.parsePairLog(ctx, vLog)
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse pair log")
			continue
		}

		msgCh <- msg
	}
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

			go func() {
				pair := &Pair{
					Number:  event.Arg3.Int64(),
					Address: event.Raw.Address.String(),
					Token0:  event.Token0.String(),
					Token1:  event.Token1.String(),
				}
				if err := s.store.AddPair(ctx, pair); err != nil {
					log.Error().Err(err).Int64("number", pair.Number).Str("address", pair.Address).Msg("Failed to add pair to store")

				}
			}()

			// Prevent gaps
			go s.queryPairs(ctx, int64(event.Raw.BlockNumber), 0, []ethcommon.Address{event.Pair}, msgCh, errCh)

			go s.subscribePairs(ctx, []ethcommon.Address{event.Pair}, msgCh, errCh)

			msg, err := s.parsePairCreatedEvent(ctx, event)
			if err != nil {
				log.Error().Err(err).Msg("Failed to parse factory pair created event")
				continue
			}

			msgCh <- msg
		}
	}
}

func (s *Source) parsePairCreatedEvent(ctx context.Context, event *factory.FactoryPairCreated) (proto.Message, error) {
	t, err := s.client.BlockTime(ctx, event.Raw.BlockNumber)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get block time")
		return nil, err
	}

	tokens, errs := s.GetTokens(ctx, event.Token0, event.Token1)
	if errs[0] != nil {
		log.Error().Err(errs[0]).Str("address", event.Token0.String()).Msg("Failed to get token0")
		return nil, errs[0]
	}
	if errs[1] != nil {
		log.Error().Err(errs[1]).Str("address", event.Token1.String()).Msg("Failed to get token1")
		return nil, errs[1]
	}

	return &factory.PairCreated{
		Ts:             &timestamppb.Timestamp{Seconds: int64(t)},
		BlockNumber:    event.Raw.BlockNumber,
		TxHash:         event.Raw.TxHash.Bytes(),
		LogIndex:       uint64(event.Raw.Index),
		Token0:         event.Token0.Bytes(),
		Token1:         event.Token1.Bytes(),
		Pair:           event.Pair.Bytes(),
		Arg3:           event.Arg3.String(),
		Token0Name:     tokens[0].Name,
		Token0Symbol:   tokens[0].Symbol,
		Token0Decimals: uint32(tokens[0].Decimals),
		Token1Name:     tokens[1].Name,
		Token1Symbol:   tokens[1].Symbol,
		Token1Decimals: uint32(tokens[1].Decimals),
	}, nil
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
			msg, err := s.parsePairLog(ctx, vLog)
			if err != nil {
				log.Error().Err(err).Msg("Failed to parse pair log")
				continue
			}

			msgCh <- msg
		}
	}
}

func (s *Source) parsePairLog(ctx context.Context, vLog types.Log) (proto.Message, error) {
	t, err := s.client.BlockTime(ctx, vLog.BlockNumber)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get block time")
		return nil, err
	}

	ts := &timestamppb.Timestamp{Seconds: int64(t)}

	event, err := pair.UnpackLog(vLog)
	if err != nil {
		log.Error().Err(err).Msg("Failed to unpack pair log")
		return nil, err
	}

	p, err := s.store.GetPair(ctx, vLog.Address.String())
	if err != nil {
		log.Error().Err(err).Str("address", vLog.Address.String()).Msg("Failed to get pair from store")
	}

	if p == nil {
		p, err = s.GetPair(ctx, vLog.Address)
		if err != nil {
			log.Error().Err(err).Str("address", vLog.Address.String()).Msg("Failed to get pair from RPC")
			return nil, err
		}
	}

	switch event := event.(type) {
	case pair.PairMint:
		tokens, errs := s.GetTokens(ctx, ethcommon.HexToAddress(p.Token0), ethcommon.HexToAddress(p.Token1))
		if errs[0] != nil {
			log.Error().Err(errs[0]).Str("address", p.Token0).Msg("Failed to get token0")
			return nil, errs[0]
		}
		if errs[1] != nil {
			log.Error().Err(errs[1]).Str("address", p.Token1).Msg("Failed to get token1")
			return nil, errs[1]
		}

		amount0 := tokenAmount(event.Amount0, tokens[0].Decimals)
		amount1 := tokenAmount(event.Amount1, tokens[1].Decimals)

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
		tokens, errs := s.GetTokens(ctx, ethcommon.HexToAddress(p.Token0), ethcommon.HexToAddress(p.Token1))
		if errs[0] != nil {
			log.Error().Err(errs[0]).Str("address", p.Token0).Msg("Failed to get token0")
			return nil, errs[0]
		}
		if errs[1] != nil {
			log.Error().Err(errs[1]).Str("address", p.Token1).Msg("Failed to get token1")
			return nil, errs[1]
		}

		amount0In := tokenAmount(event.Amount0In, tokens[0].Decimals)
		amount0Out := tokenAmount(event.Amount0Out, tokens[0].Decimals)
		amount1In := tokenAmount(event.Amount1In, tokens[1].Decimals)
		amount1Out := tokenAmount(event.Amount1Out, tokens[1].Decimals)

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
		tokens, errs := s.GetTokens(ctx, ethcommon.HexToAddress(p.Token0), ethcommon.HexToAddress(p.Token1))
		if errs[0] != nil {
			log.Error().Err(errs[0]).Str("address", p.Token0).Msg("Failed to get token0")
			return nil, errs[0]
		}
		if errs[1] != nil {
			log.Error().Err(errs[1]).Str("address", p.Token1).Msg("Failed to get token1")
			return nil, errs[1]
		}

		reserve0 := tokenAmount(event.Reserve0, tokens[0].Decimals)
		reserve1 := tokenAmount(event.Reserve1, tokens[1].Decimals)

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
		tokens, errs := s.GetTokens(ctx, ethcommon.HexToAddress(p.Token0), ethcommon.HexToAddress(p.Token1))
		if errs[0] != nil {
			log.Error().Err(errs[0]).Str("address", p.Token0).Msg("Failed to get token0")
			return nil, errs[0]
		}
		if errs[1] != nil {
			log.Error().Err(errs[1]).Str("address", p.Token1).Msg("Failed to get token1")
			return nil, errs[1]
		}

		amount0 := tokenAmount(event.Amount0, tokens[0].Decimals)
		amount1 := tokenAmount(event.Amount1, tokens[1].Decimals)

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

func (s *Source) AllPairs(ctx context.Context) []ethcommon.Address {
	s.loadAllPairsOnce.Do(func() {
		s.loadAllPairs(ctx)
	})

	return s.pairs
}

func (s *Source) GetTokens(ctx context.Context, addresses ...ethcommon.Address) ([]*Token, []error) {
	tokens := make([]*Token, len(addresses))
	errs := make([]error, len(addresses))

	var wg sync.WaitGroup

	wg.Add(len(addresses))

	for i := range addresses {
		address := addresses[i]

		go func(i int) {
			defer wg.Done()
			tokens[i], errs[i] = s.GetToken(ctx, address)
		}(i)
	}

	wg.Wait()

	return tokens, errs
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
		return s.getToken(ctx, address)
	})
	if err != nil {
		return nil, err
	}

	go func() {
		if err := s.store.AddToken(ctx, token); err != nil {
			log.Error().Err(err).Str("address", token.Address).Msg("Failed to add token to store")
		}
	}()

	return token, nil
}

func (s *Source) getToken(ctx context.Context, address ethcommon.Address) (*Token, error) {
	tokenContract, err := erc20.NewErc20(address, s.client)
	if err != nil {
		return nil, err
	}

	var (
		name             string
		symbol           string
		decimals         uint8
		err1, err2, err3 error
		wg               sync.WaitGroup
	)

	wg.Add(3)

	go func() {
		defer wg.Done()
		name, err1 = tokenContract.Name(&bind.CallOpts{Context: ctx})
	}()
	go func() {
		defer wg.Done()
		symbol, err2 = tokenContract.Symbol(&bind.CallOpts{Context: ctx})
	}()
	go func() {
		defer wg.Done()
		decimals, err3 = tokenContract.Decimals(&bind.CallOpts{Context: ctx})
	}()

	wg.Wait()

	if err1 != nil {
		log.Error().Err(err).Str("address", address.String()).Msg("Failed to get token name")
	}
	if err2 != nil {
		log.Error().Err(err).Str("address", address.String()).Msg("Failed to get token symbol")
	}
	if err3 != nil {
		return nil, err3
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
		return s.getPair(ctx, address)
	})
}

func (s *Source) getPair(ctx context.Context, address ethcommon.Address) (*Pair, error) {
	pairContract, err := pair.NewPair(address, s.client)
	if err != nil {
		return nil, err
	}

	var (
		token0, token1 ethcommon.Address
		err1, err2     error
		wg             sync.WaitGroup
	)

	wg.Add(2)

	go func() {
		defer wg.Done()
		token0, err1 = pairContract.Token0(&bind.CallOpts{Context: ctx})
	}()
	go func() {
		defer wg.Done()
		token1, err2 = pairContract.Token1(&bind.CallOpts{Context: ctx})
	}()

	wg.Wait()

	if err1 != nil {
		return nil, err1
	}
	if err2 != nil {
		return nil, err2
	}

	p := &Pair{
		Address: address.String(),
		Token0:  token0.String(),
		Token1:  token1.String(),
	}

	return p, nil
}

func (s *Source) GetPairByNumber(ctx context.Context, number *big.Int) (*Pair, error) {
	address, err := s.factoryContract.AllPairs(&bind.CallOpts{Context: ctx}, number)
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

		sem := make(chan struct{}, s.concurrecyLimit)

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

package uniswapv3

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
	"github.com/trylotus/connectors/uniswapv3/erc20"
	"github.com/trylotus/connectors/uniswapv3/factory"
	"github.com/trylotus/connectors/uniswapv3/pool"
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
	queryPageSize        int64 // Limit number of pools per query
	subscriptionPageSize int64 // Limit number of pools per subscription

	factoryContract     *factory.Factory
	factoryContractAddr ethcommon.Address
	pools               []ethcommon.Address
	loadAllPoolsOnce    sync.Once
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

	pools := s.AllPools(ctx)

	for chunkHead := fromBlock; chunkHead <= toBlock; chunkHead += s.queryPageSize {
		chunkTail := chunkHead + s.queryPageSize - 1
		if chunkTail > toBlock {
			chunkTail = toBlock
		}

		log.Info().Int64("from", chunkHead).Int64("to", chunkTail).Msg("Query")

		s.queryFactory(ctx, chunkHead, chunkTail, msgCh, errCh)

		for i := 0; i < len(pools); i += int(s.queryPageSize) {
			j := i + int(s.queryPageSize)
			if j > len(pools) {
				j = len(pools)
			}
			s.queryPools(ctx, chunkHead, chunkTail, pools[i:j], msgCh, errCh)
		}
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
	retryCtx = common.ContextWithFuncName(retryCtx, "QueryFactory")

	logs, err := common.RetryT(retryCtx, func() ([]types.Log, error) {
		subCtx, cancel := context.WithTimeout(ctx, 3*time.Minute)
		defer cancel()
		return s.client.FilterLogs(subCtx, filter)
	})

	if err != nil {
		errCh <- err
		return
	}

	for _, vLog := range logs {
		subCtx, cancel := context.WithTimeout(ctx, 1*time.Minute)
		defer cancel()

		msg, err := s.parseFactoryLog(subCtx, vLog, false, nil, nil)
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse factory log")
			continue
		}

		msgCh <- msg
	}
}

func (s *Source) queryPools(ctx context.Context, fromBlock int64, toBlock int64, pools []ethcommon.Address, msgCh chan<- proto.Message, errCh chan<- error) {
	filter := ethereum.FilterQuery{
		Addresses: pools,
		FromBlock: big.NewInt(fromBlock),
	}

	if toBlock > 0 {
		filter.ToBlock = big.NewInt(toBlock)
	}

	retryCtx := common.ContextWithConditionalRetry(ctx, evm.IsRetryableError)
	retryCtx = common.ContextWithFuncName(retryCtx, "QueryPools")

	logs, err := common.RetryT(retryCtx, func() ([]types.Log, error) {
		subCtx, cancel := context.WithTimeout(ctx, 3*time.Minute)
		defer cancel()
		return s.client.FilterLogs(subCtx, filter)
	})

	if err != nil {
		errCh <- err
		return
	}

	for _, vLog := range logs {
		subCtx, cancel := context.WithTimeout(ctx, 1*time.Minute)
		defer cancel()

		msg, err := s.parsePoolLog(subCtx, vLog)
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse pool log")
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

	pools := s.AllPools(ctx)

	for i := 0; i < len(pools); i += int(s.subscriptionPageSize) {
		j := i + int(s.subscriptionPageSize)
		if j > len(pools) {
			j = len(pools)
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			s.subscribePools(ctx, pools[i:j], msgCh, errCh)
		}()
	}

	wg.Wait()
}

func (s *Source) subscribeFactory(ctx context.Context, msgCh chan<- proto.Message, errCh chan<- error) {
	logCh := make(chan types.Log, 2048)

	sub, err := s.client.SubscribeFilterLogs(ctx, ethereum.FilterQuery{Addresses: []ethcommon.Address{s.factoryContractAddr}}, logCh)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to subscribe to factory contract")
	}

	defer sub.Unsubscribe()

	for {
		select {
		case err := <-sub.Err():
			errCh <- err
			return
		case vLog := <-logCh:
			subCtx, cancel := context.WithTimeout(ctx, 1*time.Minute)
			defer cancel()

			msg, err := s.parseFactoryLog(subCtx, vLog, true, msgCh, errCh)
			if err != nil {
				log.Error().Err(err).Msg("Failed to parse factory log")
				continue
			}

			msgCh <- msg
		}
	}
}

func (s *Source) parseFactoryLog(ctx context.Context, vLog types.Log, subscribe bool, msgCh chan<- proto.Message, errCh chan<- error) (proto.Message, error) {
	t, err := s.client.BlockTime(ctx, vLog.BlockNumber)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get block time")
		return nil, err
	}

	ts := &timestamppb.Timestamp{Seconds: int64(t)}

	event, err := factory.UnpackLog(vLog)
	if err != nil {
		log.Error().Err(err).Msg("Failed to unpack factory log")
		return nil, err
	}

	switch event := event.(type) {
	case factory.FactoryFeeAmountEnabled:
		return &factory.FeeAmountEnabled{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			TxHash:      vLog.TxHash.Bytes(),
			LogIndex:    uint64(vLog.Index),
			Fee:         uint32(event.Fee.Uint64()),
			TickSpacing: int32(event.TickSpacing.Int64()),
		}, nil
	case factory.FactoryOwnerChanged:
		return &factory.OwnerChanged{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			TxHash:      vLog.TxHash.Bytes(),
			LogIndex:    uint64(vLog.Index),
			OldOwner:    event.OldOwner.Bytes(),
			NewOwner:    event.NewOwner.Bytes(),
		}, nil
	case factory.FactoryPoolCreated:
		if subscribe {
			// Prevent gaps
			go s.queryPools(ctx, int64(event.Raw.BlockNumber), 0, []ethcommon.Address{event.Pool}, msgCh, errCh)

			go s.subscribePools(ctx, []ethcommon.Address{event.Pool}, msgCh, errCh)
		}

		tokens, errs := s.GetTokens(ctx, event.Token0, event.Token1)
		if errs[0] != nil {
			log.Error().Err(errs[0]).Str("address", event.Token0.String()).Msg("Failed to get token0")
			if errs[0].Error() != "no contract code at given address" {
				return nil, errs[0]
			}
		}
		if errs[1] != nil {
			log.Error().Err(errs[1]).Str("address", event.Token1.String()).Msg("Failed to get token1")
			if errs[1].Error() != "no contract code at given address" {
				return nil, errs[1]
			}
		}

		msg := &factory.PoolCreated{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			TxHash:      vLog.TxHash.Bytes(),
			LogIndex:    uint64(vLog.Index),
			Token0:      event.Token0.Bytes(),
			Token1:      event.Token1.Bytes(),
			Fee:         uint32(event.Fee.Uint64()),
			TickSpacing: int32(event.TickSpacing.Int64()),
			Pool:        event.Pool.Bytes(),
		}

		if tokens[0] != nil {
			msg.Token0Name = tokens[0].Name
			msg.Token0Symbol = tokens[0].Symbol
			msg.Token0Decimals = uint32(tokens[0].Decimals)
		}

		if tokens[1] != nil {
			msg.Token1Name = tokens[1].Name
			msg.Token1Symbol = tokens[1].Symbol
			msg.Token1Decimals = uint32(tokens[1].Decimals)
		}

		return msg, nil
	default:
		return nil, fmt.Errorf("unhandled event: %s", reflect.TypeOf(event))
	}
}

func (s *Source) subscribePools(ctx context.Context, pools []ethcommon.Address, msgCh chan<- proto.Message, errCh chan<- error) {
	logCh := make(chan types.Log, 2048)

	sub, err := s.client.SubscribeFilterLogs(ctx, ethereum.FilterQuery{Addresses: pools}, logCh)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to subscribe to pool contracts")
	}

	defer sub.Unsubscribe()

	for {
		select {
		case err := <-sub.Err():
			errCh <- err
			return
		case vLog := <-logCh:
			subCtx, cancel := context.WithTimeout(ctx, 1*time.Minute)
			defer cancel()

			msg, err := s.parsePoolLog(subCtx, vLog)
			if err != nil {
				log.Error().Err(err).Msg("Failed to parse pool log")
				continue
			}

			msgCh <- msg
		}
	}
}

func (s *Source) parsePoolLog(ctx context.Context, vLog types.Log) (proto.Message, error) {
	t, err := s.client.BlockTime(ctx, vLog.BlockNumber)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get block time")
		return nil, err
	}

	ts := &timestamppb.Timestamp{Seconds: int64(t)}

	event, err := pool.UnpackLog(vLog)
	if err != nil {
		log.Error().Err(err).Msg("Failed to unpack pool log")
		return nil, err
	}

	p, err := s.store.GetPool(ctx, vLog.Address.String())
	if err != nil {
		log.Error().Err(err).Str("address", vLog.Address.String()).Msg("Failed to get pool from store")
	}

	if p == nil {
		p, err = s.GetPool(ctx, vLog.Address)
		if err != nil {
			log.Error().Err(err).Str("address", vLog.Address.String()).Msg("Failed to get pool from RPC")
			return nil, err
		}

		if err := s.store.AddPool(ctx, p); err != nil {
			log.Error().Err(err).Str("address", p.Address).Msg("Failed to get pool to store")
		}
	}

	switch event := event.(type) {
	case pool.PoolBurn:
		tokens, errs := s.GetTokens(ctx, ethcommon.HexToAddress(p.Token0), ethcommon.HexToAddress(p.Token1))
		if errs[0] != nil {
			log.Error().Err(errs[0]).Str("address", p.Token0).Msg("Failed to get token0")
			return nil, errs[0]
		}
		if errs[1] != nil {
			log.Error().Err(errs[1]).Str("address", p.Token1).Msg("Failed to get token1")
			return nil, errs[1]
		}

		// Uniswapv3 LP token decimals is always 18
		amount := tokenAmount(event.Amount, 18)
		amount0 := tokenAmount(event.Amount0, tokens[0].Decimals)
		amount1 := tokenAmount(event.Amount1, tokens[1].Decimals)

		return &pool.Burn{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			TxHash:      vLog.TxHash.Bytes(),
			LogIndex:    uint64(vLog.Index),
			Pool:        vLog.Address.Bytes(),
			Owner:       event.Owner.Bytes(),
			TickLower:   int32(event.TickLower.Int64()),
			TickUpper:   int32(event.TickUpper.Int64()),
			Amount:      floatString(amount),
			Amount0:     floatString(amount0),
			Amount1:     floatString(amount1),
		}, nil
	case pool.PoolCollect:
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

		return &pool.Collect{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			TxHash:      vLog.TxHash.Bytes(),
			LogIndex:    uint64(vLog.Index),
			Pool:        vLog.Address.Bytes(),
			Owner:       event.Owner.Bytes(),
			Recipient:   event.Recipient.Bytes(),
			TickLower:   int32(event.TickLower.Int64()),
			TickUpper:   int32(event.TickUpper.Int64()),
			Amount0:     floatString(amount0),
			Amount1:     floatString(amount1),
		}, nil
	case pool.PoolCollectProtocol:
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

		return &pool.CollectProtocol{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			TxHash:      vLog.TxHash.Bytes(),
			LogIndex:    uint64(vLog.Index),
			Pool:        vLog.Address.Bytes(),
			Sender:      event.Sender.Bytes(),
			Recipient:   event.Recipient.Bytes(),
			Amount0:     floatString(amount0),
			Amount1:     floatString(amount1),
		}, nil
	case pool.PoolFlash:
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
		paid0 := tokenAmount(event.Paid0, tokens[0].Decimals)
		paid1 := tokenAmount(event.Paid1, tokens[1].Decimals)

		return &pool.Flash{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			TxHash:      vLog.TxHash.Bytes(),
			LogIndex:    uint64(vLog.Index),
			Pool:        vLog.Address.Bytes(),
			Sender:      event.Sender.Bytes(),
			Recipient:   event.Recipient.Bytes(),
			Amount0:     floatString(amount0),
			Amount1:     floatString(amount1),
			Paid0:       floatString(paid0),
			Paid1:       floatString(paid1),
		}, nil
	case pool.PoolIncreaseObservationCardinalityNext:
		return &pool.IncreaseObservationCardinalityNext{
			Ts:                            ts,
			BlockNumber:                   vLog.BlockNumber,
			TxHash:                        vLog.TxHash.Bytes(),
			LogIndex:                      uint64(vLog.Index),
			Pool:                          vLog.Address.Bytes(),
			ObservationCardinalityNextOld: uint32(event.ObservationCardinalityNextOld),
			ObservationCardinalityNextNew: uint32(event.ObservationCardinalityNextNew),
		}, nil
	case pool.PoolInitialize:
		return &pool.Initialize{
			Ts:           ts,
			BlockNumber:  vLog.BlockNumber,
			TxHash:       vLog.TxHash.Bytes(),
			LogIndex:     uint64(vLog.Index),
			Pool:         vLog.Address.Bytes(),
			SqrtPriceX96: event.SqrtPriceX96.String(),
			Tick:         int32(event.Tick.Int64()),
		}, nil
	case pool.PoolMint:
		tokens, errs := s.GetTokens(ctx, ethcommon.HexToAddress(p.Token0), ethcommon.HexToAddress(p.Token1))
		if errs[0] != nil {
			log.Error().Err(errs[0]).Str("address", p.Token0).Msg("Failed to get token0")
			return nil, errs[0]
		}
		if errs[1] != nil {
			log.Error().Err(errs[1]).Str("address", p.Token1).Msg("Failed to get token1")
			return nil, errs[1]
		}

		// Uniswapv3 LP token decimals is always 18
		amount := tokenAmount(event.Amount, 18)
		amount0 := tokenAmount(event.Amount0, tokens[0].Decimals)
		amount1 := tokenAmount(event.Amount1, tokens[1].Decimals)

		return &pool.Mint{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			TxHash:      vLog.TxHash.Bytes(),
			LogIndex:    uint64(vLog.Index),
			Pool:        vLog.Address.Bytes(),
			Sender:      event.Sender.Bytes(),
			Owner:       event.Owner.Bytes(),
			TickLower:   int32(event.TickLower.Int64()),
			TickUpper:   int32(event.TickUpper.Int64()),
			Amount:      floatString(amount),
			Amount0:     floatString(amount0),
			Amount1:     floatString(amount1),
		}, nil
	case pool.PoolSetFeeProtocol:
		return &pool.SetFeeProtocol{
			Ts:              ts,
			BlockNumber:     vLog.BlockNumber,
			TxHash:          vLog.TxHash.Bytes(),
			LogIndex:        uint64(vLog.Index),
			Pool:            vLog.Address.Bytes(),
			FeeProtocol0Old: uint32(event.FeeProtocol0Old),
			FeeProtocol1Old: uint32(event.FeeProtocol1Old),
			FeeProtocol0New: uint32(event.FeeProtocol0New),
			FeeProtocol1New: uint32(event.FeeProtocol1New),
		}, nil
	case pool.PoolSwap:
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

		return &pool.Swap{
			Ts:           ts,
			BlockNumber:  vLog.BlockNumber,
			TxHash:       vLog.TxHash.Bytes(),
			LogIndex:     uint64(vLog.Index),
			Pool:         vLog.Address.Bytes(),
			Sender:       event.Sender.Bytes(),
			Recipient:    event.Sender.Bytes(),
			Amount0:      floatString(amount0),
			Amount1:      floatString(amount1),
			SqrtPriceX96: event.SqrtPriceX96.String(),
			Liquidity:    event.Liquidity.String(),
			Tick:         int32(event.Tick.Int64()),
		}, nil
	default:
		return nil, fmt.Errorf("unhandled event: %s", reflect.TypeOf(event))
	}
}

func (s *Source) AllPools(ctx context.Context) []ethcommon.Address {
	s.loadAllPoolsOnce.Do(func() {
		s.loadAllPools(ctx)
	})

	return s.pools
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
		log.Error().Err(err1).Str("address", address.String()).Msg("Failed to get token name")
	}
	if err2 != nil {
		log.Error().Err(err2).Str("address", address.String()).Msg("Failed to get token symbol")
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

func (s *Source) GetPool(ctx context.Context, address ethcommon.Address) (*Pool, error) {
	retryCtx := common.ContextWithConditionalRetry(ctx, evm.IsRetryableError)
	retryCtx = common.ContextWithFuncName(retryCtx, "GetPool")

	return common.RetryT(retryCtx, func() (*Pool, error) {
		subCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		return s.getPool(subCtx, address)
	})
}

func (s *Source) getPool(ctx context.Context, address ethcommon.Address) (*Pool, error) {
	poolContract, err := pool.NewPool(address, s.client)
	if err != nil {
		return nil, err
	}

	var (
		token0, token1         ethcommon.Address
		fee, tickSpacing       *big.Int
		err1, err2, err3, err4 error
		wg                     sync.WaitGroup
	)

	wg.Add(4)

	go func() {
		defer wg.Done()
		token0, err1 = poolContract.Token0(&bind.CallOpts{Context: ctx})
	}()
	go func() {
		defer wg.Done()
		token1, err2 = poolContract.Token1(&bind.CallOpts{Context: ctx})
	}()
	go func() {
		defer wg.Done()
		fee, err3 = poolContract.Fee(&bind.CallOpts{Context: ctx})
	}()
	go func() {
		defer wg.Done()
		tickSpacing, err4 = poolContract.TickSpacing(&bind.CallOpts{Context: ctx})
	}()

	if err1 != nil {
		return nil, err1
	}
	if err2 != nil {
		return nil, err2
	}
	if err3 != nil {
		return nil, err3
	}
	if err4 != nil {
		return nil, err4
	}

	return &Pool{
		Address:     address.String(),
		Token0:      token0.String(),
		Token1:      token1.String(),
		Fee:         fee.Int64(),
		TickSpacing: tickSpacing.Int64(),
	}, nil
}

func (s *Source) loadAllPools(ctx context.Context) {
	log.Info().Msg("Loading all pools")

	s.loadPoolsFromStore(ctx)

	scannedBlock, err := s.store.GetScannedBlock(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get scanned block number")
	}

	blockNumber, err := s.client.BlockNumber(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get block number")
	}

	for i := uint64(scannedBlock); i <= blockNumber; i += uint64(s.blockRangeLimit) {
		j := i + uint64(s.blockRangeLimit) - 1
		if j > blockNumber {
			j = blockNumber
		}

		s.loadPoolsFromRPC(ctx, i, j)
	}

	log.Info().Int("total", len(s.pools)).Msg("Successfully loaded all pools")
}

func (s *Source) loadPoolsFromStore(ctx context.Context) {
	log.Info().Msg("Loading pools from store")

	poolCh, errCh := s.store.AllPools(ctx)

	for pool := range poolCh {
		s.pools = append(s.pools, ethcommon.HexToAddress(pool.Address))

		log.Debug().Str("address", pool.Address).Msg("Loaded pool from store")
	}

	for err := range errCh {
		log.Fatal().Err(err).Msg("Failed to load pools from store")
	}

	log.Info().Int("total", len(s.pools)).Msg("Successfully loaded pools from store")
}

func (s *Source) loadPoolsFromRPC(ctx context.Context, from uint64, to uint64) {
	log.Info().Uint64("from", from).Uint64("to", to).Msg("Loading pools from RPC")

	opts := &bind.FilterOpts{
		Context: ctx,
		Start:   from,
		End:     &to,
	}

	it, err := s.factoryContract.FilterPoolCreated(opts, nil, nil, nil)
	if err != nil {
		log.Fatal().Err(err).Uint64("from", from).Uint64("to", to).Msg("Failed to load pools from RPC")
	}

	defer it.Close()

	for it.Next() {
		if err := it.Error(); err != nil {
			log.Fatal().Err(err).Uint64("from", from).Uint64("to", to).Msg("Failed to load pool from RPC")
		}

		s.pools = append(s.pools, it.Event.Pool)

		pool := Pool{
			Address:     it.Event.Pool.String(),
			Token0:      it.Event.Token0.String(),
			Token1:      it.Event.Token1.String(),
			Fee:         it.Event.Fee.Int64(),
			TickSpacing: it.Event.TickSpacing.Int64(),
		}

		if err := s.store.AddPool(ctx, &pool); err != nil {
			log.Error().Err(err).Str("address", pool.Address).Msg("Failed to add pool to store")
		}
	}

	log.Info().Uint64("from", from).Uint64("to", to).Msg("Successfully loaded pools from RPC")

	if err := s.store.SetScannedBlock(ctx, int64(to)); err != nil {
		log.Error().Err(err).Uint64("number", to).Msg("Failed to set scanned block")
	}
}

func floatString(n *big.Rat) string {
	s := n.FloatString(18)
	s = strings.TrimRight(s, "0")
	return strings.TrimSuffix(s, ".")
}

func tokenAmount(rawAmount *big.Int, decimals uint8) *big.Rat {
	return new(big.Rat).SetFrac(rawAmount, new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil))
}

package uniswapv2

import (
	"context"
	"math/big"
	"reflect"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/trylotus/go-connector"
	"github.com/trylotus/go-connector/log"
	"github.com/trylotus/go-connector/source/evm"
	"github.com/trylotus/uniswapv2/erc20"
	"github.com/trylotus/uniswapv2/factory"
	"github.com/trylotus/uniswapv2/pair"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	pairPageSize    = 100000
	concurrenyLimit = 5
)

type Source struct {
	client *evm.Client
	store  *Store

	factoryContract  *factory.Factory
	pairs            []ethcommon.Address
	loadAllPairsOnce sync.Once
}

var _ connector.Source = (*Source)(nil)

func NewSource(client *evm.Client, store *Store, factoryContractAddr string) *Source {
	factoryContract, err := factory.NewFactory(ethcommon.HexToAddress(factoryContractAddr), client)
	if err != nil {
		log.Fatal().Err(err).Str("contract", factoryContractAddr).Msg("Failed to create factory contract")
	}

	return &Source{
		client:          client,
		store:           store,
		factoryContract: factoryContract,
	}
}

func (s *Source) BlockNumber(ctx context.Context) (int64, error) {
	n, err := s.client.BlockNumber(ctx)
	if err != nil {
		return 0, err
	}

	return int64(n), nil
}

func (s *Source) Query(ctx context.Context, fromBlock int64, toBlock int64) (<-chan proto.Message, <-chan error) {
	panic("unimplemented")
}

func (s *Source) Subscribe(ctx context.Context) (<-chan proto.Message, <-chan error) {
	msgCh := make(chan proto.Message, 2048)
	errCh := make(chan error, 1)

	go s.subscribeFactory(ctx, msgCh, errCh)

	pairs := s.AllPairs(ctx)

	for i := 0; i < len(pairs); i += pairPageSize {
		j := i + pairPageSize
		if j > len(pairs) {
			j = len(pairs)
		}
		go s.subscribePairs(ctx, pairs[i:j], msgCh, errCh)
	}

	return msgCh, errCh
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

			t, err := s.client.BlockTime(ctx, event.Raw.BlockNumber)
			if err != nil {
				log.Error().Err(err).Msg("Failed to get block time")
			}

			msg := &factory.PairCreated{
				Ts:          &timestamppb.Timestamp{Seconds: int64(t)},
				BlockNumber: event.Raw.BlockNumber,
				TxHash:      event.Raw.TxHash.Bytes(),
				Index:       uint64(event.Raw.Index),
				Arg_3:       event.Arg3.String(),
				Pair:        event.Pair.Bytes(),
				Token_0:     event.Token0.Bytes(),
				Token_1:     event.Token1.Bytes(),
			}

			msgCh <- msg

			go s.subscribePairs(ctx, []ethcommon.Address{event.Pair}, msgCh, errCh)

			go func() {
				pair := &Pair{
					Number:  event.Arg3.Int64(),
					Address: event.Pair.String(),
					Token0:  event.Token0.String(),
					Token1:  event.Token1.String(),
				}
				if err := s.store.AddPair(ctx, pair); err != nil {
					log.Error().Err(err).Int64("number", pair.Number).Str("address", pair.Address).Msg("Failed to add pair to store")
				}
			}()
		}
	}
}

func (s *Source) subscribePairs(ctx context.Context, pairs []ethcommon.Address, msgCh chan<- proto.Message, errCh chan<- error) {
	logCh := make(chan types.Log, 2048)

	sub, err := s.client.SubscribeFilterLogs(ctx, ethereum.FilterQuery{Addresses: pairs}, logCh)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to subscribe to pair contracts")
	}

	for {
		select {
		case err := <-sub.Err():
			errCh <- err
			return
		case vLog := <-logCh:
			t, err := s.client.BlockTime(ctx, vLog.BlockNumber)
			if err != nil {
				log.Error().Err(err).Msg("Failed to get block time")
				continue
			}

			ts := &timestamppb.Timestamp{Seconds: int64(t)}

			event, err := pair.UnpackLog(vLog)
			if err != nil {
				log.Error().Err(err).Msg("Failed to unpack pair log")
				continue
			}

			p, err := s.store.GetPair(ctx, vLog.Address.String())
			if err != nil {
				log.Error().Err(err).Str("address", vLog.Address.String()).Msg("Failed to get pair from store")
			}

			if p == nil {
				p, err = s.GetPair(ctx, vLog.Address)
				if err != nil {
					log.Error().Err(err).Str("address", vLog.Address.String()).Msg("Failed to get pair from RPC")
					continue
				}
			}

			switch event := event.(type) {
			case pair.PairMint:
				token0, err := s.GetToken(ctx, ethcommon.HexToAddress(p.Token0))
				if err != nil {
					log.Error().Err(err).Str("address", p.Token0).Msg("Failed to get token0")
					continue
				}

				token1, err := s.GetToken(ctx, ethcommon.HexToAddress(p.Token0))
				if err != nil {
					log.Error().Err(err).Str("address", p.Token0).Msg("Failed to get token1")
					continue
				}

				amount0 := tokenAmount(event.Amount0, token0.Decimals)
				amount1 := tokenAmount(event.Amount1, token1.Decimals)

				msgCh <- &pair.Mint{
					Ts:          ts,
					BlockNumber: vLog.BlockNumber,
					TxHash:      vLog.TxHash.Bytes(),
					Index:       uint64(vLog.Index),
					Sender:      event.Sender.Bytes(),
					Amount_0:    floatString(amount0),
					Amount_1:    floatString(amount1),
				}
			case pair.PairSwap:
				token0, err := s.GetToken(ctx, ethcommon.HexToAddress(p.Token0))
				if err != nil {
					log.Error().Err(err).Str("address", p.Token0).Msg("Failed to get token0")
					continue
				}

				token1, err := s.GetToken(ctx, ethcommon.HexToAddress(p.Token0))
				if err != nil {
					log.Error().Err(err).Str("address", p.Token0).Msg("Failed to get token1")
					continue
				}

				amount0In := tokenAmount(event.Amount0In, token0.Decimals)
				amount0Out := tokenAmount(event.Amount0Out, token0.Decimals)
				amount1In := tokenAmount(event.Amount1In, token1.Decimals)
				amount1Out := tokenAmount(event.Amount1Out, token1.Decimals)

				msgCh <- &pair.Swap{
					Ts:          ts,
					BlockNumber: vLog.BlockNumber,
					TxHash:      vLog.TxHash.Bytes(),
					Index:       uint64(vLog.Index),
					Amount_0In:  floatString(amount0In),
					Amount_0Out: floatString(amount0Out),
					Amount_1In:  floatString(amount1In),
					Amount_1Out: floatString(amount1Out),
					Sender:      event.Sender.Bytes(),
					To:          event.To.Bytes(),
				}
			case pair.PairSync:
				token0, err := s.GetToken(ctx, ethcommon.HexToAddress(p.Token0))
				if err != nil {
					log.Error().Err(err).Str("address", p.Token0).Msg("Failed to get token0")
					continue
				}

				token1, err := s.GetToken(ctx, ethcommon.HexToAddress(p.Token0))
				if err != nil {
					log.Error().Err(err).Str("address", p.Token0).Msg("Failed to get token1")
					continue
				}

				reserve0 := tokenAmount(event.Reserve0, token0.Decimals)
				reserve1 := tokenAmount(event.Reserve1, token1.Decimals)

				msgCh <- &pair.Sync{
					Ts:          ts,
					BlockNumber: vLog.BlockNumber,
					TxHash:      vLog.TxHash.Bytes(),
					Index:       uint64(vLog.Index),
					Reserve_0:   floatString(reserve0),
					Reserve_1:   floatString(reserve1),
				}
			case pair.PairTransfer:
				lpToken, err := s.GetToken(ctx, vLog.Address)
				if err != nil {
					log.Error().Err(err).Str("address", p.Token0).Msg("Failed to get lp token")
					continue
				}

				value := tokenAmount(event.Value, lpToken.Decimals)

				msgCh <- &pair.Transfer{
					Ts:          ts,
					BlockNumber: vLog.BlockNumber,
					TxHash:      vLog.TxHash.Bytes(),
					Index:       uint64(vLog.Index),
					From:        event.From.Bytes(),
					To:          event.To.Bytes(),
					Value:       floatString(value),
				}
			case pair.PairApproval:
				lpToken, err := s.GetToken(ctx, vLog.Address)
				if err != nil {
					log.Error().Err(err).Str("address", p.Token0).Msg("Failed to get lp token")
					continue
				}

				value := tokenAmount(event.Value, lpToken.Decimals)

				msgCh <- &pair.Approval{
					Ts:          ts,
					BlockNumber: vLog.BlockNumber,
					TxHash:      vLog.TxHash.Bytes(),
					Index:       uint64(vLog.Index),
					Owner:       event.Owner.Bytes(),
					Spender:     event.Spender.Bytes(),
					Value:       floatString(value),
				}
			case pair.PairBurn:
				token0, err := s.GetToken(ctx, ethcommon.HexToAddress(p.Token0))
				if err != nil {
					log.Error().Err(err).Str("address", p.Token0).Msg("Failed to get token0")
					continue
				}

				token1, err := s.GetToken(ctx, ethcommon.HexToAddress(p.Token0))
				if err != nil {
					log.Error().Err(err).Str("address", p.Token0).Msg("Failed to get token1")
					continue
				}

				amount0 := tokenAmount(event.Amount0, token0.Decimals)
				amount1 := tokenAmount(event.Amount1, token1.Decimals)

				msgCh <- &pair.Burn{
					Ts:          ts,
					BlockNumber: vLog.BlockNumber,
					TxHash:      vLog.TxHash.Bytes(),
					Index:       uint64(vLog.Index),
					Sender:      event.Sender.Bytes(),
					Amount_0:    floatString(amount0),
					Amount_1:    floatString(amount1),
					To:          event.To.Bytes(),
				}
			default:
				log.Error().Str("event", reflect.TypeOf(event).String()).Msg("Unhandled event")
			}
		}
	}
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
		return nil, err1
	}
	if err2 != nil {
		return nil, err2
	}
	if err3 != nil {
		return nil, err3
	}

	token = &Token{
		Address:  address.String(),
		Name:     name,
		Symbol:   symbol,
		Decimals: decimals,
	}

	go func() {
		if err := s.store.AddToken(ctx, token); err != nil {
			log.Error().Err(err).Str("address", token.Address).Msg("Failed to add token to store")
		}
	}()

	return token, nil
}

func (s *Source) GetPair(ctx context.Context, address ethcommon.Address) (*Pair, error) {
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

func (s *Source) GetPairsSync(ctx context.Context, from int64, to int64) ([]*Pair, error) {
	pairCount := to - from + 1
	pairs := make([]*Pair, pairCount)
	errs := make([]error, pairCount)

	var wg sync.WaitGroup

	for i := int64(0); i < pairCount; i++ {
		wg.Add(1)
		go func(i int64) {
			defer wg.Done()

			pair, err := s.GetPairByNumber(ctx, big.NewInt(from+i))
			pairs[i] = pair
			errs[i] = err
		}(i)
	}

	wg.Wait()

	for _, err := range errs {
		if err != nil {
			return nil, err
		}
	}

	return pairs, nil
}

func (s *Source) GetPairs(ctx context.Context, from int64, to int64) (<-chan *Pair, <-chan error) {
	pairCh := make(chan *Pair, 100)
	errCh := make(chan error, 1)

	go func() {
		defer close(pairCh)
		defer close(errCh)

		for i := from; i <= to; i += concurrenyLimit {
			j := i + concurrenyLimit - 1
			if j > to {
				j = to
			}

			pairs, err := s.GetPairsSync(ctx, i, j)
			if err != nil {
				errCh <- err
				return
			}

			for _, pair := range pairs {
				pairCh <- pair
			}
		}
	}()

	return pairCh, errCh
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

	pairCh, errCh := s.GetPairs(ctx, from, to)

	for pair := range pairCh {
		s.pairs = append(s.pairs, ethcommon.HexToAddress(pair.Address))

		if err := s.store.AddPair(ctx, pair); err != nil {
			log.Error().Err(err).Int64("number", pair.Number).Str("address", pair.Address).Msg("Failed to add pair to store")
		}

		log.Debug().Int64("number", pair.Number).Str("address", pair.Address).Msg("Loaded pair from RPC")
	}

	for err := range errCh {
		log.Fatal().Err(err).Int64("from", from).Int64("to", to).Msg("Failed to load pair from RPC")
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

package monad

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	protoevm "github.com/trylotus/connectors/monad/evm"
	"github.com/trylotus/go-connector"
	"github.com/trylotus/go-connector/source/evm"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Source struct {
	client *evm.Client
}

var _ connector.Source = (*Source)(nil)

func NewSource(client *evm.Client) *Source {
	return &Source{client}
}

func (s *Source) BlockNumber(ctx context.Context) (int64, error) {
	n, err := s.client.BlockNumber(ctx)
	return int64(n), err
}

func (s *Source) Query(ctx context.Context, fromBlock int64, toBlock int64) ([]proto.Message, error) {
	var msgs []proto.Message

	for blockNumber := fromBlock; blockNumber <= toBlock; blockNumber++ {
		block, err := s.client.BlockByNumber(ctx, big.NewInt(blockNumber))
		if err != nil {
			return nil, err
		}

		header := block.Header()

		msgBlock := parseBlock(header)
		msgs = append(msgs, msgBlock)

		for _, tx := range block.Transactions() {
			msgTx, err := parseTransaction(header, tx)
			if err != nil {
				return nil, err
			}

			msgs = append(msgs, msgTx)
		}
	}

	return msgs, nil
}

func (s *Source) Subscribe(ctx context.Context, msgCh chan<- proto.Message, errCh chan<- error) {
	latestBlock, err := s.BlockNumber(ctx)
	if err != nil {
		errCh <- err
		return
	}

	msgs, err := s.Query(ctx, latestBlock, latestBlock)
	if err != nil {
		errCh <- err
		return
	}

	for _, msg := range msgs {
		msgCh <- msg
	}

	t := time.NewTicker(1 * time.Second)

	for {
		<-t.C

		currentBlock, err := s.BlockNumber(ctx)
		if err != nil {
			errCh <- err
			return
		}

		if currentBlock <= latestBlock {
			continue
		}

		msgs, err := s.Query(ctx, latestBlock+1, currentBlock)
		if err != nil {
			errCh <- err
			return
		}

		for _, msg := range msgs {
			msgCh <- msg
		}

		latestBlock = currentBlock
	}
}

func parseBlock(header *types.Header) *protoevm.Block {
	return &protoevm.Block{
		Ts: &timestamppb.Timestamp{
			Seconds: int64(header.Time),
		},
		Hash:       header.Hash().Bytes(),
		Number:     header.Number.Uint64(),
		Difficulty: header.Difficulty.String(),
		GasLimit:   header.GasLimit,
		GasUsed:    header.GasUsed,
		Nonce:      header.Nonce.Uint64(),
	}
}

func parseTransaction(header *types.Header, tx *types.Transaction) (*protoevm.Transaction, error) {
	from, err := types.Sender(types.LatestSignerForChainID(tx.ChainId()), tx)
	if err != nil {
		return nil, err
	}

	var to []byte
	if tx.To() != nil {
		to = tx.To().Bytes()
	}

	v, r, s := tx.RawSignatureValues()

	msgTx := &protoevm.Transaction{
		Ts: &timestamppb.Timestamp{
			Seconds: int64(header.Time),
		},
		BlockHash:   header.Hash().Bytes(),
		BlockNumber: header.Number.Uint64(),
		Hash:        tx.Hash().Bytes(),
		From:        from.Bytes(),
		To:          to,
		Size:        tx.Size(),
		Nonce:       tx.Nonce(),
		Gas:         tx.Gas(),
		GasPrice:    tx.GasPrice().String(),
		Value:       tx.Value().String(),
		Data:        tx.Data(),
		V:           v.String(),
		R:           r.String(),
		S:           s.String(),
	}

	return msgTx, nil
}

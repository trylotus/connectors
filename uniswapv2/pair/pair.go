package pair

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
	"github.com/trylotus/go-connector/common"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var pairAbi abi.ABI

func init() {
	var err error
	pairAbi, err = abi.JSON(strings.NewReader(PairMetaData.ABI))
	if err != nil {
		log.Fatal().Err(err).Msg("error reading pair abi")
	}
}

func UnpackLog(vLog types.Log) (interface{}, error) {
	ev, err := pairAbi.EventByID(vLog.Topics[0])
	if err != nil {
		return nil, fmt.Errorf("cannot find event by ID: %v", err)
	}
	switch ev.Name {
	case "Mint":
		var event PairMint
		if err := common.UnpackLog(pairAbi, &event, ev.Name, vLog); err != nil {
			return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
		}
		return event, nil
	case "Swap":
		var event PairSwap
		if err := common.UnpackLog(pairAbi, &event, ev.Name, vLog); err != nil {
			return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
		}
		return event, nil
	case "Sync":
		var event PairSync
		if err := common.UnpackLog(pairAbi, &event, ev.Name, vLog); err != nil {
			return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
		}
		return event, nil
	case "Transfer":
		var event PairTransfer
		if err := common.UnpackLog(pairAbi, &event, ev.Name, vLog); err != nil {
			return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
		}
		return event, nil
	case "Approval":
		var event PairApproval
		if err := common.UnpackLog(pairAbi, &event, ev.Name, vLog); err != nil {
			return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
		}
		return event, nil
	case "Burn":
		var event PairBurn
		if err := common.UnpackLog(pairAbi, &event, ev.Name, vLog); err != nil {
			return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
		}
		return event, nil
	default:
		return nil, fmt.Errorf("invalid event: %s", ev.Name)
	}
}

func Message(vLog types.Log, ts *timestamppb.Timestamp) (proto.Message, error) {
	ev, err := pairAbi.EventByID(vLog.Topics[0])
	if err != nil {
		return nil, fmt.Errorf("cannot find event by ID: %v", err)
	}
	switch ev.Name {
	case "Mint":
		event := new(PairMint)
		if err := common.UnpackLog(pairAbi, event, ev.Name, vLog); err != nil {
			return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
		}
		return &Mint{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			TxHash:      vLog.TxHash.Bytes(),
			Index:       uint64(vLog.Index),
			Amount0:     event.Amount0.String(),
			Amount1:     event.Amount1.String(),
			Sender:      event.Sender.Bytes(),
		}, nil
	case "Swap":
		event := new(PairSwap)
		if err := common.UnpackLog(pairAbi, event, ev.Name, vLog); err != nil {
			return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
		}
		return &Swap{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			TxHash:      vLog.TxHash.Bytes(),
			Index:       uint64(vLog.Index),
			Amount0In:   event.Amount0In.String(),
			Amount0Out:  event.Amount0Out.String(),
			Amount1In:   event.Amount1In.String(),
			Amount1Out:  event.Amount1Out.String(),
			Sender:      event.Sender.Bytes(),
			To:          event.To.Bytes(),
		}, nil
	case "Sync":
		event := new(PairSync)
		if err := common.UnpackLog(pairAbi, event, ev.Name, vLog); err != nil {
			return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
		}
		return &Sync{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			TxHash:      vLog.TxHash.Bytes(),
			Index:       uint64(vLog.Index),
			Reserve0:    event.Reserve0.String(),
			Reserve1:    event.Reserve1.String(),
		}, nil
	case "Transfer":
		event := new(PairTransfer)
		if err := common.UnpackLog(pairAbi, event, ev.Name, vLog); err != nil {
			return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
		}
		return &Transfer{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			TxHash:      vLog.TxHash.Bytes(),
			Index:       uint64(vLog.Index),
			From:        event.From.Bytes(),
			To:          event.To.Bytes(),
			Value:       event.Value.String(),
		}, nil
	case "Approval":
		event := new(PairApproval)
		if err := common.UnpackLog(pairAbi, event, ev.Name, vLog); err != nil {
			return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
		}
		return &Approval{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			TxHash:      vLog.TxHash.Bytes(),
			Index:       uint64(vLog.Index),
			Owner:       event.Owner.Bytes(),
			Spender:     event.Spender.Bytes(),
			Value:       event.Value.String(),
		}, nil
	case "Burn":
		event := new(PairBurn)
		if err := common.UnpackLog(pairAbi, event, ev.Name, vLog); err != nil {
			return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
		}
		return &Burn{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			TxHash:      vLog.TxHash.Bytes(),
			Index:       uint64(vLog.Index),
			Amount0:     event.Amount0.String(),
			Amount1:     event.Amount1.String(),
			Sender:      event.Sender.Bytes(),
			To:          event.To.Bytes(),
		}, nil
	default:
		return nil, fmt.Errorf("invalid event: %s", ev.Name)
	}
}

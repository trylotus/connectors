package BAT

import (
	"github.com/trylotus/connector/common"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SmartContract struct{}

func (sc *SmartContract) Message(eventName string, contractAbi *abi.ABI, vLog types.Log, timestamp *timestamppb.Timestamp) protoreflect.ProtoMessage {
	switch eventName {
	case "Approval":
		e := new(BATApproval)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Approval{
			Ts:      timestamp,
			Owner:   e.Owner.Bytes(),
			Spender: e.Spender.Bytes(),
			Value:   e.Value.Bytes(),
		}
	case "CreateBAT":
		e := new(BATCreateBAT)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &CreateBAT{
			Ts:    timestamp,
			To:    e.To.Bytes(),
			Value: e.Value.Bytes(),
		}
	case "LogRefund":
		e := new(BATLogRefund)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &LogRefund{
			Ts:    timestamp,
			To:    e.To.Bytes(),
			Value: e.Value.Bytes(),
		}
	case "Transfer":
		e := new(BATTransfer)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Transfer{
			Ts:    timestamp,
			From:  e.From.Bytes(),
			To:    e.To.Bytes(),
			Value: e.Value.Bytes(),
		}
	}
	return nil
}

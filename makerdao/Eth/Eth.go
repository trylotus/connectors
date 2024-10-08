package Eth

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
		e := new(ETHApproval)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Approval{
			Ts:  timestamp,
			Src: e.Src.Bytes(),
			Guy: e.Guy.Bytes(),
			Wad: e.Wad.Bytes(),
		}
	case "Deposit":
		e := new(ETHDeposit)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Deposit{
			Ts:  timestamp,
			Dst: e.Dst.Bytes(),
			Wad: e.Wad.Bytes(),
		}
	case "Transfer":
		e := new(ETHTransfer)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Transfer{
			Ts:  timestamp,
			Src: e.Src.Bytes(),
			Dst: e.Dst.Bytes(),
			Wad: e.Wad.Bytes(),
		}
	case "Withdrawal":
		e := new(ETHWithdrawal)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Withdrawal{
			Ts:  timestamp,
			Src: e.Src.Bytes(),
			Wad: e.Wad.Bytes(),
		}
	}
	return nil
}

package ArbitrumL2GovRelay

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
	case "Closed":
		e := new(ARBITRUML2GOVRELAYClosed)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Closed{
			Ts: timestamp,
		}
	case "Deny":
		e := new(ARBITRUML2GOVRELAYDeny)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Deny{
			Ts:  timestamp,
			Usr: e.Usr.Bytes(),
		}
	case "ERC20DepositInitiated":
		e := new(ARBITRUML2GOVRELAYERC20DepositInitiated)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &ERC20DepositInitiated{
			Ts:      timestamp,
			L1Token: e.L1Token.Bytes(),
			L2Token: e.L2Token.Bytes(),
			From:    e.From.Bytes(),
			To:      e.To.Bytes(),
			Amount:  e.Amount.Bytes(),
			Data:    e.Data[:],
		}
	case "ERC20WithdrawalFinalized":
		e := new(ARBITRUML2GOVRELAYERC20WithdrawalFinalized)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &ERC20WithdrawalFinalized{
			Ts:      timestamp,
			L1Token: e.L1Token.Bytes(),
			L2Token: e.L2Token.Bytes(),
			From:    e.From.Bytes(),
			To:      e.To.Bytes(),
			Amount:  e.Amount.Bytes(),
			Data:    e.Data[:],
		}
	case "Rely":
		e := new(ARBITRUML2GOVRELAYRely)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Rely{
			Ts:  timestamp,
			Usr: e.Usr.Bytes(),
		}
	}
	return nil
}

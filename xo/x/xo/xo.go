package xo

import (
	"context"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
	"github.com/trylotus/go-connector/common"
	"github.com/trylotus/go-connector/source/evm"
	"google.golang.org/protobuf/proto"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type SmartContract struct {
	Abi  abi.ABI
	Addr ethcommon.Address
}

var _ evm.SmartContract = (*SmartContract)(nil)

func NewContract(address ethcommon.Address) *SmartContract {
	contractAbi, err := abi.JSON(strings.NewReader(XoMetaData.ABI))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to read xo ABI")
	}
	return &SmartContract{Addr: address, Abi: contractAbi}
}

func (c *SmartContract) Address() ethcommon.Address {
	return c.Addr
}

func (c *SmartContract) ParseLog(_ context.Context, vLog types.Log, ts *timestamppb.Timestamp) (proto.Message, error) {
	ev, err := c.Abi.EventByID(vLog.Topics[0])
	if err != nil {
		return nil, evm.InvalidLogError(err)
	}
	switch ev.Name {
	case "PaidDM":
		var event XoPaidDM
		if err := common.UnpackLog(c.Abi, &event, ev.Name, vLog); err != nil {
			return nil, evm.InvalidLogError(err)
		}
		return &PaidDM{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			BlockHash:   vLog.BlockHash.Bytes(),
			TxHash:      vLog.TxHash.Bytes(),
			LogIndex:    uint64(vLog.Index),
			UserId:      event.UserId.Bytes(),
			Denom:       event.Denom.Bytes(),
			Amount:      event.Amount.String(),
		}, nil
	case "RoleAdminChanged":
		var event XoRoleAdminChanged
		if err := common.UnpackLog(c.Abi, &event, ev.Name, vLog); err != nil {
			return nil, evm.InvalidLogError(err)
		}
		return &RoleAdminChanged{
			Ts:                ts,
			BlockNumber:       vLog.BlockNumber,
			BlockHash:         vLog.BlockHash.Bytes(),
			TxHash:            vLog.TxHash.Bytes(),
			LogIndex:          uint64(vLog.Index),
			Role:              event.Role[:],
			PreviousAdminRole: event.PreviousAdminRole[:],
			NewAdminRole:      event.NewAdminRole[:],
		}, nil
	case "SBTUpdated":
		var event XoSBTUpdated
		if err := common.UnpackLog(c.Abi, &event, ev.Name, vLog); err != nil {
			return nil, evm.InvalidLogError(err)
		}
		return &SBTUpdated{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			BlockHash:   vLog.BlockHash.Bytes(),
			TxHash:      vLog.TxHash.Bytes(),
			LogIndex:    uint64(vLog.Index),
			PostId:      event.PostId.Bytes(),
		}, nil
	case "NewGoodVibes":
		var event XoNewGoodVibes
		if err := common.UnpackLog(c.Abi, &event, ev.Name, vLog); err != nil {
			return nil, evm.InvalidLogError(err)
		}
		return &NewGoodVibes{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			BlockHash:   vLog.BlockHash.Bytes(),
			TxHash:      vLog.TxHash.Bytes(),
			LogIndex:    uint64(vLog.Index),
			PostId:      event.PostId.Bytes(),
		}, nil
	case "Post":
		var event XoPost
		if err := common.UnpackLog(c.Abi, &event, ev.Name, vLog); err != nil {
			return nil, evm.InvalidLogError(err)
		}
		return &Post{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			BlockHash:   vLog.BlockHash.Bytes(),
			TxHash:      vLog.TxHash.Bytes(),
			LogIndex:    uint64(vLog.Index),
			Poster:      event.Poster.Bytes(),
			PostId:      event.PostId.Bytes(),
		}, nil
	case "RoleGranted":
		var event XoRoleGranted
		if err := common.UnpackLog(c.Abi, &event, ev.Name, vLog); err != nil {
			return nil, evm.InvalidLogError(err)
		}
		return &RoleGranted{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			BlockHash:   vLog.BlockHash.Bytes(),
			TxHash:      vLog.TxHash.Bytes(),
			LogIndex:    uint64(vLog.Index),
			Role:        event.Role[:],
			Account:     event.Account.Bytes(),
			Sender:      event.Sender.Bytes(),
		}, nil
	case "Initialized":
		var event XoInitialized
		if err := common.UnpackLog(c.Abi, &event, ev.Name, vLog); err != nil {
			return nil, evm.InvalidLogError(err)
		}
		return &Initialized{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			BlockHash:   vLog.BlockHash.Bytes(),
			TxHash:      vLog.TxHash.Bytes(),
			LogIndex:    uint64(vLog.Index),
			Version:     event.Version,
		}, nil
	case "NewMutualLike":
		var event XoNewMutualLike
		if err := common.UnpackLog(c.Abi, &event, ev.Name, vLog); err != nil {
			return nil, evm.InvalidLogError(err)
		}
		return &NewMutualLike{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			BlockHash:   vLog.BlockHash.Bytes(),
			TxHash:      vLog.TxHash.Bytes(),
			LogIndex:    uint64(vLog.Index),
			MyId:        event.MyId.Bytes(),
			TargetId:    event.TargetId.Bytes(),
		}, nil
	case "RoleRevoked":
		var event XoRoleRevoked
		if err := common.UnpackLog(c.Abi, &event, ev.Name, vLog); err != nil {
			return nil, evm.InvalidLogError(err)
		}
		return &RoleRevoked{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			BlockHash:   vLog.BlockHash.Bytes(),
			TxHash:      vLog.TxHash.Bytes(),
			LogIndex:    uint64(vLog.Index),
			Role:        event.Role[:],
			Account:     event.Account.Bytes(),
			Sender:      event.Sender.Bytes(),
		}, nil
	case "Swiped":
		var event XoSwiped
		if err := common.UnpackLog(c.Abi, &event, ev.Name, vLog); err != nil {
			return nil, evm.InvalidLogError(err)
		}
		return &Swiped{
			Ts:           ts,
			BlockNumber:  vLog.BlockNumber,
			BlockHash:    vLog.BlockHash.Bytes(),
			TxHash:       vLog.TxHash.Bytes(),
			LogIndex:     uint64(vLog.Index),
			Sender:       event.Sender.Bytes(),
			SwipedUserId: event.SwipedUserId.Bytes(),
			CardId:       event.CardId.Bytes(),
		}, nil
	case "Streak":
		var event XoStreak
		if err := common.UnpackLog(c.Abi, &event, ev.Name, vLog); err != nil {
			return nil, evm.InvalidLogError(err)
		}
		return &Streak{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			BlockHash:   vLog.BlockHash.Bytes(),
			TxHash:      vLog.TxHash.Bytes(),
			LogIndex:    uint64(vLog.Index),
			UserId:      event.UserId.Bytes(),
			Streak:      event.Streak.String(),
		}, nil
	case "SaveStreak":
		var event XoSaveStreak
		if err := common.UnpackLog(c.Abi, &event, ev.Name, vLog); err != nil {
			return nil, evm.InvalidLogError(err)
		}
		return &SaveStreak{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			BlockHash:   vLog.BlockHash.Bytes(),
			TxHash:      vLog.TxHash.Bytes(),
			LogIndex:    uint64(vLog.Index),
			UserId:      event.UserId.Bytes(),
			Streak:      event.Streak.String(),
		}, nil
	default:
		return nil, evm.InvalidLogErrorf("unhandled event: %s", ev.Name)
	}
}

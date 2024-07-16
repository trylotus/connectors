package xo

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
	"github.com/trylotus/connector/common"
	"google.golang.org/protobuf/proto"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type SmartContract struct {
	Addr string
	Abi  abi.ABI
}

func NewContract(addr string) *SmartContract {
	contractAbi, err := abi.JSON(strings.NewReader(XoABI))
	if err != nil {
		log.Fatal().Err(err).Msg("error reading cyber XoABI")
	}
	return &SmartContract{Addr: addr, Abi: contractAbi}
}

func (c *SmartContract) Address() string {
	return c.Addr
}

func (c *SmartContract) Message(vLog types.Log, ts *timestamppb.Timestamp) (proto.Message, error) {
	ev, err := c.Abi.EventByID(vLog.Topics[0])
	if err != nil {
		return nil, fmt.Errorf("cannot find event by ID: %v", err)
	}
	switch ev.Name {
	case "PaidDM":
		event := new(XoPaidDM)
		if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
			return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
		}
		return &PaidDM{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			TxHash:      vLog.TxHash.Bytes(),
			Index:       uint64(vLog.Index),
			UserId:      event.UserId.Bytes(),
			Denom:       event.Denom.Bytes(),
			Amount:      event.Amount.String(),
		}, nil
	case "RoleAdminChanged":
		event := new(XoRoleAdminChanged)
		if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
			return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
		}
		return &RoleAdminChanged{
			Ts:                ts,
			BlockNumber:       vLog.BlockNumber,
			TxHash:            vLog.TxHash.Bytes(),
			Index:             uint64(vLog.Index),
			Role:              event.Role[:],
			PreviousAdminRole: event.PreviousAdminRole[:],
			NewAdminRole:      event.NewAdminRole[:],
		}, nil
	case "SBTUpdated":
		event := new(XoSBTUpdated)
		if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
			return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
		}
		return &SBTUpdated{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			TxHash:      vLog.TxHash.Bytes(),
			Index:       uint64(vLog.Index),
			PostId:      event.PostId.Bytes(),
		}, nil
	case "NewGoodVibes":
		event := new(XoNewGoodVibes)
		if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
			return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
		}
		return &NewGoodVibes{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			TxHash:      vLog.TxHash.Bytes(),
			Index:       uint64(vLog.Index),
			PostId:      event.PostId.Bytes(),
		}, nil
	case "Post":
		event := new(XoPost)
		if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
			return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
		}
		return &Post{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			TxHash:      vLog.TxHash.Bytes(),
			Index:       uint64(vLog.Index),
			Poster:      event.Poster.Bytes(),
			PostId:      event.PostId.Bytes(),
		}, nil
	case "RoleGranted":
		event := new(XoRoleGranted)
		if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
			return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
		}
		return &RoleGranted{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			TxHash:      vLog.TxHash.Bytes(),
			Index:       uint64(vLog.Index),
			Role:        event.Role[:],
			Account:     event.Account.Bytes(),
			Sender:      event.Sender.Bytes(),
		}, nil
	case "Initialized":
		event := new(XoInitialized)
		if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
			return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
		}
		return &Initialized{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			TxHash:      vLog.TxHash.Bytes(),
			Index:       uint64(vLog.Index),
			Version:     event.Version,
		}, nil
	case "NewMutualLike":
		event := new(XoNewMutualLike)
		if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
			return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
		}
		return &NewMutualLike{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			TxHash:      vLog.TxHash.Bytes(),
			Index:       uint64(vLog.Index),
			MyId:        event.MyId.Bytes(),
			TargetId:    event.TargetId.Bytes(),
		}, nil
	case "RoleRevoked":
		event := new(XoRoleRevoked)
		if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
			return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
		}
		return &RoleRevoked{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			TxHash:      vLog.TxHash.Bytes(),
			Index:       uint64(vLog.Index),
			Role:        event.Role[:],
			Account:     event.Account.Bytes(),
			Sender:      event.Sender.Bytes(),
		}, nil
	case "Swiped":
		event := new(XoSwiped)
		if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
			return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
		}
		return &Swiped{
			Ts:           ts,
			BlockNumber:  vLog.BlockNumber,
			TxHash:       vLog.TxHash.Bytes(),
			Index:        uint64(vLog.Index),
			Sender:       event.Sender.Bytes(),
			SwipedUserId: event.SwipedUserId.Bytes(),
			CardId:       event.CardId.Bytes(),
		}, nil
	case "Streak":
		event := new(XoStreak)
		if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
			return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
		}
		return &Streak{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			TxHash:      vLog.TxHash.Bytes(),
			Index:       uint64(vLog.Index),
			UserId:      event.UserId.Bytes(),
			Streak:      event.Streak.String(),
		}, nil
	case "SaveStreak":
		event := new(XoSaveStreak)
		if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
			return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
		}
		return &SaveStreak{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			TxHash:      vLog.TxHash.Bytes(),
			Index:       uint64(vLog.Index),
			UserId:      event.UserId.Bytes(),
			Streak:      event.Streak.String(),
		}, nil
	default:
		return nil, fmt.Errorf("invalid event: %s", ev.Name)
	}
}

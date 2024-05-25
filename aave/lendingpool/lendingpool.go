package lendingpool

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
	"github.com/trylotus/connector/common"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SmartContract struct {
	Addr string
	Abi  abi.ABI
}

func NewContract(addr string) *SmartContract {
	contractAbi, err := abi.JSON(strings.NewReader(LendingpoolABI))
	if err != nil {
		log.Fatal().Err(err).Msg("error reading aave lending pool abi")
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
	case "Borrow":
		event := new(LendingpoolBorrow)
		if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
			return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
		}
		return &Borrow{
			Ts:             ts,
			BlockNumber:    vLog.BlockNumber,
			LogIndex:       uint64(vLog.Index),
			TxHash:         vLog.TxHash.Bytes(),
			Reserve:        event.Reserve.Bytes(),
			User:           event.User.Bytes(),
			OnBehalfOf:     event.OnBehalfOf.Bytes(),
			Amount:         event.Amount.Bytes(),
			BorrowRate:     event.BorrowRate.Bytes(),
			BorrowRateMode: event.BorrowRateMode.Bytes(),
			Referral:       uint32(event.Referral),
		}, nil
	case "Deposit":
		event := new(LendingpoolDeposit)
		if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
			return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
		}
		return &Deposit{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			LogIndex:    uint64(vLog.Index),
			TxHash:      vLog.TxHash.Bytes(),
			Reserve:     event.Reserve.Bytes(),
			User:        event.User.Bytes(),
			OnBehalfOf:  event.OnBehalfOf.Bytes(),
			Amount:      event.Amount.Bytes(),
			Referral:    uint32(event.Referral),
		}, nil
	case "FlashLoan":
		event := new(LendingpoolFlashLoan)
		if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
			return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
		}
		return &FlashLoan{
			Ts:           ts,
			BlockNumber:  vLog.BlockNumber,
			LogIndex:     uint64(vLog.Index),
			TxHash:       vLog.TxHash.Bytes(),
			Target:       event.Target.Bytes(),
			Initiator:    event.Initiator.Bytes(),
			Asset:        event.Asset.Bytes(),
			Amount:       event.Amount.Bytes(),
			Premium:      event.Premium.Bytes(),
			ReferralCode: uint32(event.ReferralCode),
		}, nil
	case "LiquidationCall":
		event := new(LendingpoolLiquidationCall)
		if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
			return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
		}
		return &LiquidationCall{
			Ts:                         ts,
			BlockNumber:                vLog.BlockNumber,
			LogIndex:                   uint64(vLog.Index),
			TxHash:                     vLog.TxHash.Bytes(),
			CollateralAsset:            event.CollateralAsset.Bytes(),
			DebtAsset:                  event.DebtAsset.Bytes(),
			User:                       event.User.Bytes(),
			DebtToCover:                event.DebtToCover.Bytes(),
			LiquidatedCollateralAmount: event.LiquidatedCollateralAmount.Bytes(),
			Liquidator:                 event.Liquidator.Bytes(),
			ReceiveAToken:              event.ReceiveAToken,
		}, nil
	case "RebalanceStableBorrowRate":
		event := new(LendingpoolRebalanceStableBorrowRate)
		if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
			return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
		}
		return &RebalanceStableBorrowRate{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			LogIndex:    uint64(vLog.Index),
			TxHash:      vLog.TxHash.Bytes(),
			Reserve:     event.Reserve.Bytes(),
			User:        event.User.Bytes(),
		}, nil
	case "Repay":
		event := new(LendingpoolRepay)
		if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
			return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
		}
		return &Repay{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			LogIndex:    uint64(vLog.Index),
			TxHash:      vLog.TxHash.Bytes(),
			Reserve:     event.Reserve.Bytes(),
			User:        event.User.Bytes(),
			Repayer:     event.Repayer.Bytes(),
			Amount:      event.Amount.Bytes(),
		}, nil
	case "ReserveDataUpdated":
		event := new(LendingpoolReserveDataUpdated)
		if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
			return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
		}
		return &ReserveDataUpdated{
			Ts:                  ts,
			BlockNumber:         vLog.BlockNumber,
			LogIndex:            uint64(vLog.Index),
			TxHash:              vLog.TxHash.Bytes(),
			LiquidityRate:       event.LiquidityRate.Bytes(),
			LiquidityIndex:      event.LiquidityIndex.Bytes(),
			StableBorrowRate:    event.StableBorrowRate.Bytes(),
			VariableBorrowIndex: event.VariableBorrowIndex.Bytes(),
			VariableBorrowRate:  event.VariableBorrowRate.Bytes(),
		}, nil
	case "ReserveUsedAsCollateralDisabled":
		event := new(LendingpoolReserveUsedAsCollateralDisabled)
		if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
			return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
		}
		return &ReserveUsedAsCollateralDisabled{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			LogIndex:    uint64(vLog.Index),
			TxHash:      vLog.TxHash.Bytes(),
			Reserve:     event.Reserve.Bytes(),
			User:        event.User.Bytes(),
		}, nil
	case "ReserveUsedAsCollateralEnabled":
		event := new(LendingpoolReserveUsedAsCollateralEnabled)
		if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
			return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
		}
		return &ReserveUsedAsCollateralEnabled{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			LogIndex:    uint64(vLog.Index),
			TxHash:      vLog.TxHash.Bytes(),
			Reserve:     event.Reserve.Bytes(),
			User:        event.User.Bytes(),
		}, nil
	case "Swap":
		event := new(LendingpoolSwap)
		if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
			return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
		}
		return &Swap{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			LogIndex:    uint64(vLog.Index),
			TxHash:      vLog.TxHash.Bytes(),
			Reserve:     event.Reserve.Bytes(),
			User:        event.User.Bytes(),
			RateMode:    event.RateMode.Bytes(),
		}, nil
	case "Withdraw":
		event := new(LendingpoolWithdraw)
		if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
			return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
		}
		return &Withdraw{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			LogIndex:    uint64(vLog.Index),
			TxHash:      vLog.TxHash.Bytes(),
			Reserve:     event.Reserve.Bytes(),
			User:        event.User.Bytes(),
			To:          event.To.Bytes(),
			Amount:      event.Amount.Bytes(),
		}, nil
	case "Paused":
		event := new(LendingpoolPaused)
		if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
			return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
		}
		return &Paused{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			LogIndex:    uint64(vLog.Index),
			TxHash:      vLog.TxHash.Bytes(),
		}, nil
	case "Unpaused":
		event := new(LendingpoolUnpaused)
		if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
			return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
		}
		return &Unpaused{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			LogIndex:    uint64(vLog.Index),
			TxHash:      vLog.TxHash.Bytes(),
		}, nil
	default:
		return nil, fmt.Errorf("invalid event: %s", ev.Name)
	}
}

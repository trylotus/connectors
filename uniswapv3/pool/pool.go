package pool

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
	"github.com/trylotus/go-connector/common"
)

var poolAbi abi.ABI

func init() {
	var err error
	poolAbi, err = abi.JSON(strings.NewReader(PoolMetaData.ABI))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to read pool ABI")
	}
}

func UnpackLog(vLog types.Log) (interface{}, error) {
	ev, err := poolAbi.EventByID(vLog.Topics[0])
	if err != nil {
		return nil, fmt.Errorf("no event with id: %s", vLog.Topics[0].Hex())
	}
	switch ev.Name {
	case "Burn":
		var event PoolBurn
		if err := common.UnpackLog(poolAbi, &event, ev.Name, vLog); err != nil {
			return nil, fmt.Errorf("error unpacking event %s: %w", ev.Name, err)
		}
		return event, nil
	case "Collect":
		var event PoolCollect
		if err := common.UnpackLog(poolAbi, &event, ev.Name, vLog); err != nil {
			return nil, fmt.Errorf("error unpacking event %s: %w", ev.Name, err)
		}
		return event, nil
	case "CollectProtocol":
		var event PoolCollectProtocol
		if err := common.UnpackLog(poolAbi, &event, ev.Name, vLog); err != nil {
			return nil, fmt.Errorf("error unpacking event %s: %w", ev.Name, err)
		}
		return event, nil
	case "Flash":
		var event PoolFlash
		if err := common.UnpackLog(poolAbi, &event, ev.Name, vLog); err != nil {
			return nil, fmt.Errorf("error unpacking event %s: %w", ev.Name, err)
		}
		return event, nil
	case "IncreaseObservationCardinalityNext":
		var event PoolIncreaseObservationCardinalityNext
		if err := common.UnpackLog(poolAbi, &event, ev.Name, vLog); err != nil {
			return nil, fmt.Errorf("error unpacking event %s: %w", ev.Name, err)
		}
		return event, nil
	case "Initialize":
		var event PoolInitialize
		if err := common.UnpackLog(poolAbi, &event, ev.Name, vLog); err != nil {
			return nil, fmt.Errorf("error unpacking event %s: %w", ev.Name, err)
		}
		return event, nil
	case "Mint":
		var event PoolMint
		if err := common.UnpackLog(poolAbi, &event, ev.Name, vLog); err != nil {
			return nil, fmt.Errorf("error unpacking event %s: %w", ev.Name, err)
		}
		return event, nil
	case "SetFeeProtocol":
		var event PoolSetFeeProtocol
		if err := common.UnpackLog(poolAbi, &event, ev.Name, vLog); err != nil {
			return nil, fmt.Errorf("error unpacking event %s: %w", ev.Name, err)
		}
		return event, nil
	case "Swap":
		var event PoolSwap
		if err := common.UnpackLog(poolAbi, &event, ev.Name, vLog); err != nil {
			return nil, fmt.Errorf("error unpacking event %s: %w", ev.Name, err)
		}
		return event, nil
	default:
		return nil, fmt.Errorf("invalid event: %s", ev.Name)
	}
}

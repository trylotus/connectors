package factory

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
	"github.com/trylotus/go-connector/common"
)

var factoryAbi abi.ABI

func init() {
	var err error
	factoryAbi, err = abi.JSON(strings.NewReader(FactoryMetaData.ABI))
	if err != nil {
		log.Fatal().Err(err).Msg("error reading factory abi")
	}
}

func UnpackLog(vLog types.Log) (interface{}, error) {
	ev, err := factoryAbi.EventByID(vLog.Topics[0])
	if err != nil {
		return nil, fmt.Errorf("cannot find event by ID: %v", err)
	}
	switch ev.Name {
	case "FeeAmountEnabled":
		var event FactoryFeeAmountEnabled
		if err := common.UnpackLog(factoryAbi, &event, ev.Name, vLog); err != nil {
			return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
		}
		return event, nil
	case "OwnerChanged":
		var event FactoryOwnerChanged
		if err := common.UnpackLog(factoryAbi, &event, ev.Name, vLog); err != nil {
			return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
		}
		return event, nil
	case "PoolCreated":
		var event FactoryPoolCreated
		if err := common.UnpackLog(factoryAbi, &event, ev.Name, vLog); err != nil {
			return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
		}
		return event, nil
	default:
		return nil, fmt.Errorf("invalid event: %s", ev.Name)
	}
}

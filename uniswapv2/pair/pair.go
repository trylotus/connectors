package pair

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
	"github.com/trylotus/go-connector/common"
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

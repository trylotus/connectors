package TransparentUpgradeableProxy

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
	contractAbi, err := abi.JSON(strings.NewReader(TransparentUpgradeableProxyABI))
	if err != nil {
		log.Fatal().Err(err).Msg("error reading cyber TransparentUpgradeableProxyABI")
	}
	for _, event := range contractAbi.Events {
		fmt.Printf("%s -> %s\n", event.Name, event.ID)
	}
	return &SmartContract{Addr: addr, Abi: contractAbi}
}

func (c *SmartContract) Address() string {
	return c.Addr
}

func (c *SmartContract) Topics() []string {
	return []string{
		"0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f",
		"0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b",
	}
}

func (c *SmartContract) Message(vLog types.Log, ts *timestamppb.Timestamp) (proto.Message, error) {
	fmt.Println(vLog.Topics)
	fmt.Println(c.Abi.Events)
	ev, err := c.Abi.EventByID(vLog.Topics[0])
	if err != nil {
		return nil, fmt.Errorf("cannot find event by ID: %v", err)
	}
	switch ev.Name {
	case "AdminChanged":
		event := new(AdminChanged)
		if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
			return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
		}
		return &AdminChanged{
			Ts:            ts,
			PreviousAdmin: event.PreviousAdmin,
			NewAdmin:      event.NewAdmin,
		}, nil
	case "Upgraded":
		event := new(Upgraded)
		if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
			return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
		}
		return &Upgraded{
			Ts:             ts,
			Implementation: event.Implementation,
		}, nil
	default:
		return nil, fmt.Errorf("invalid event: %s", ev.Name)
	}
}

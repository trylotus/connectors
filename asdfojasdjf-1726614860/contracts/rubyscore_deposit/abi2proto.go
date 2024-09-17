package rubyscore_deposit

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"github.com/trylotus/go-connector/common"
)

type SmartContract struct {
	Addr string
	Abi  abi.ABI
}

func NewContract(addr string) *SmartContract {
	contractAbi, err := abi.JSON(strings.NewReader(ABIMetaData.ABI))
	if err != nil {
		log.Fatal().Err(err).Msg("error reading abi")
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

/* Code generated by https://github.com/trylotus/connector-bot. DO NOT EDIT */



case "Claimed":
	event := new(ABIClaimed)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &Claimed{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		Amount: event.Amount.String(),
		Recipient: event.Recipient.Bytes(),
	}


	return protoMsg, nil


case "Deposit":
	event := new(ABIDeposit)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &Deposit{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		Amount: event.Amount.String(),
		User: event.User.Bytes(),
	}


	return protoMsg, nil


case "EIP712DomainChanged":
	event := new(ABIEIP712DomainChanged)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &EIP712DomainChanged{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

	}


	return protoMsg, nil


case "RoleAdminChanged":
	event := new(ABIRoleAdminChanged)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &RoleAdminChanged{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		NewAdminRole: event.NewAdminRole[:],
		PreviousAdminRole: event.PreviousAdminRole[:],
		Role: event.Role[:],
	}


	return protoMsg, nil


case "RoleGranted":
	event := new(ABIRoleGranted)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &RoleGranted{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		Account: event.Account.Bytes(),
		Role: event.Role[:],
		Sender: event.Sender.Bytes(),
	}


	return protoMsg, nil


case "RoleRevoked":
	event := new(ABIRoleRevoked)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &RoleRevoked{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		Account: event.Account.Bytes(),
		Role: event.Role[:],
		Sender: event.Sender.Bytes(),
	}


	return protoMsg, nil


case "Withdrawal":
	event := new(ABIWithdrawal)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &Withdrawal{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		Amount: event.Amount.String(),
		Tax: event.Tax.String(),
		User: event.User.Bytes(),
	}


	return protoMsg, nil



/* Code generated by https://github.com/trylotus/connector-bot. DO NOT EDIT */

	default:
		return nil, fmt.Errorf("invalid event: %s", ev.Name)
	}
}

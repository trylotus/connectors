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

/* Code generated by https://github.com/trylotus/connector-bot. DO NOT EDIT */


case "Paused":
	event := new(XoPaused)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &Paused{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		Account: event.Account.String(),
	}


	return protoMsg, nil


case "RoleAdminChanged":
	event := new(XoRoleAdminChanged)
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


case "TokenInfoUpdated":
	event := new(XoTokenInfoUpdated)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &TokenInfoUpdated{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		Artist: event.Artist.String(),
		MintStartTimestamp: event.MintStartTimestamp.Bytes(),
		Project: event.Project.String(),
		TokenId: event.TokenId.Bytes(),
		TokenUri: event.TokenUri,
	}


	return protoMsg, nil


case "TokenMintStartTimestampUpdated":
	event := new(XoTokenMintStartTimestampUpdated)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &TokenMintStartTimestampUpdated{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		MintStartTimestamp: event.MintStartTimestamp.Bytes(),
		TokenId: event.TokenId.Bytes(),
	}


	return protoMsg, nil


case "Unpaused":
	event := new(XoUnpaused)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &Unpaused{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		Account: event.Account.String(),
	}


	return protoMsg, nil


case "ApprovalForAll":
	event := new(XoApprovalForAll)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &ApprovalForAll{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		Account: event.Account.String(),
		Approved: event.Approved,
		Operator: event.Operator.String(),
	}


	return protoMsg, nil


case "Mint":
	event := new(XoMint)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &Mint{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		Amount: event.Amount.Bytes(),
		To: event.To.String(),
		TokenId: event.TokenId.Bytes(),
	}


	return protoMsg, nil


case "RoleGranted":
	event := new(XoRoleGranted)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &RoleGranted{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		Account: event.Account.String(),
		Role: event.Role[:],
		Sender: event.Sender.String(),
	}


	return protoMsg, nil


case "RoleRevoked":
	event := new(XoRoleRevoked)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &RoleRevoked{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		Account: event.Account.String(),
		Role: event.Role[:],
		Sender: event.Sender.String(),
	}


	return protoMsg, nil


case "TransferBatch":
	event := new(XoTransferBatch)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &TransferBatch{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		From: event.From.String(),
		Operator: event.Operator.String(),
		To: event.To.String(),
	}


	for _, v := range event.Ids {
		protoMsg.Ids = append(protoMsg.Ids, v.Bytes())
	}

	for _, v := range event.Values {
		protoMsg.Values = append(protoMsg.Values, v.Bytes())
	}

	return protoMsg, nil


case "TransferSingle":
	event := new(XoTransferSingle)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &TransferSingle{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		From: event.From.String(),
		Id: event.Id.Bytes(),
		Operator: event.Operator.String(),
		To: event.To.String(),
		Value: event.Value.Bytes(),
	}


	return protoMsg, nil


case "URI":
	event := new(XoURI)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &URI{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		Id: event.Id.Bytes(),
		Value: event.Value,
	}


	return protoMsg, nil



/* Code generated by https://github.com/trylotus/connector-bot. DO NOT EDIT */

	default:
		return nil, fmt.Errorf("invalid event: %s", ev.Name)
	}
}




package quest_object

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



case "TransferBatch":
	event := new(ABITransferBatch)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &TransferBatch{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		From: event.From.Bytes(),
		Operator: event.Operator.Bytes(),
		To: event.To.Bytes(),
	}


	for _, v := range event.Ids {
		protoMsg.Ids = append(protoMsg.Ids, v.String())
	}

	for _, v := range event.Values {
		protoMsg.Values = append(protoMsg.Values, v.String())
	}

	return protoMsg, nil


case "CreateQuestObject":
	event := new(ABICreateQuestObject)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &CreateQuestObject{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		Creator: event.Creator.Bytes(),
		MaxClaimed: event.MaxClaimed.String(),
		Price: event.Price.String(),
		TokenId: event.TokenId.String(),
		Uri: event.Uri,
	}


	return protoMsg, nil


case "SetOpenForSale":
	event := new(ABISetOpenForSale)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &SetOpenForSale{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		Check: event.Check,
		TokenId: event.TokenId.String(),
	}


	return protoMsg, nil


case "SetShopAddress":
	event := new(ABISetShopAddress)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &SetShopAddress{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		ShopAddress: event.ShopAddress.Bytes(),
	}


	return protoMsg, nil


case "SetTokenURI":
	event := new(ABISetTokenURI)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &SetTokenURI{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		TokenId: event.TokenId.String(),
		Uri: event.Uri,
	}


	return protoMsg, nil


case "URI":
	event := new(ABIURI)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &URI{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		Id: event.Id.String(),
		Value: event.Value,
	}


	return protoMsg, nil


case "ApprovalForAll":
	event := new(ABIApprovalForAll)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &ApprovalForAll{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		Account: event.Account.Bytes(),
		Approved: event.Approved,
		Operator: event.Operator.Bytes(),
	}


	return protoMsg, nil


case "SetMaxClaimed":
	event := new(ABISetMaxClaimed)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &SetMaxClaimed{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		NewMaxClaimed: event.NewMaxClaimed.String(),
		TokenId: event.TokenId.String(),
	}


	return protoMsg, nil


case "SetExp":
	event := new(ABISetExp)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &SetExp{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		Exp: event.Exp.String(),
		TokenId: event.TokenId.String(),
	}


	return protoMsg, nil


case "SetSecondaryRoyalityFee":
	event := new(ABISetSecondaryRoyalityFee)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &SetSecondaryRoyalityFee{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		SecondaryRoyalty: event.SecondaryRoyalty.String(),
	}


	return protoMsg, nil


case "SetTreasuryAddress":
	event := new(ABISetTreasuryAddress)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &SetTreasuryAddress{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		TreasuryAddress: event.TreasuryAddress.Bytes(),
	}


	return protoMsg, nil


case "TransferSingle":
	event := new(ABITransferSingle)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &TransferSingle{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		From: event.From.Bytes(),
		Id: event.Id.String(),
		Operator: event.Operator.Bytes(),
		To: event.To.Bytes(),
		Value: event.Value.String(),
	}


	return protoMsg, nil


case "ChangeTokenPrice":
	event := new(ABIChangeTokenPrice)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &ChangeTokenPrice{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		NewPrice: event.NewPrice.String(),
		TokenId: event.TokenId.String(),
	}


	return protoMsg, nil


case "OwnershipGranted":
	event := new(ABIOwnershipGranted)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &OwnershipGranted{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		Operator: event.Operator.Bytes(),
		Target: event.Target.Bytes(),
	}


	return protoMsg, nil


case "PaymentReceivedOwner":
	event := new(ABIPaymentReceivedOwner)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &PaymentReceivedOwner{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		Amount: event.Amount.String(),
	}


	return protoMsg, nil


case "PaymentWithdrawnOwner":
	event := new(ABIPaymentWithdrawnOwner)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &PaymentWithdrawnOwner{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		Amount: event.Amount.String(),
	}


	return protoMsg, nil


case "SetCreator":
	event := new(ABISetCreator)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &SetCreator{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		Creator: event.Creator.Bytes(),
		TokenId: event.TokenId.String(),
	}


	return protoMsg, nil


case "SetRoyalityFee":
	event := new(ABISetRoyalityFee)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &SetRoyalityFee{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		RoyalityFee: event.RoyalityFee.String(),
	}


	return protoMsg, nil


case "SetSize":
	event := new(ABISetSize)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &SetSize{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		TokenId: event.TokenId.String(),
		X: uint32(event.X),
		Y: uint32(event.Y),
		Z: uint32(event.Z),
	}


	return protoMsg, nil


case "SetbaseMetadataURI":
	event := new(ABISetbaseMetadataURI)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &SetbaseMetadataURI{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		Baseuri: event.Baseuri,
	}


	return protoMsg, nil


case "LogGetQuestObject":
	event := new(ABILogGetQuestObject)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &LogGetQuestObject{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		Sender: event.Sender.Bytes(),
		TokenId: event.TokenId.String(),
	}


	return protoMsg, nil


case "OwnershipRemoved":
	event := new(ABIOwnershipRemoved)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &OwnershipRemoved{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		Operator: event.Operator.Bytes(),
		Target: event.Target.Bytes(),
	}


	return protoMsg, nil



/* Code generated by https://github.com/trylotus/connector-bot. DO NOT EDIT */

	default:
		return nil, fmt.Errorf("invalid event: %s", ev.Name)
	}
}

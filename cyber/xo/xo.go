package xo

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
	"github.com/trylotus/go-connector/common"
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


case "SetDefaultAdapterParamsForChainId":
	event := new(XoSetDefaultAdapterParamsForChainId)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &SetDefaultAdapterParamsForChainId{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		AdapterParams: event.AdapterParams,
		ChainId: uint32(event.ChainId),
		ProofType: uint32(event.ProofType),
	}


	return protoMsg, nil


case "WithdrawNative":
	event := new(XoWithdrawNative)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &WithdrawNative{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		Amount: event.Amount.Bytes(),
		MsgSender: event.MsgSender.String(),
		To: event.To.String(),
	}


	return protoMsg, nil


case "WithdrawZRO":
	event := new(XoWithdrawZRO)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &WithdrawZRO{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		Amount: event.Amount.Bytes(),
		MsgSender: event.MsgSender.String(),
		To: event.To.String(),
	}


	return protoMsg, nil


case "AppConfigUpdated":
	event := new(XoAppConfigUpdated)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &AppConfigUpdated{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		ConfigType: event.ConfigType.Bytes(),
		NewConfig: event.NewConfig,
		UserApplication: event.UserApplication.String(),
	}


	return protoMsg, nil


case "OwnershipTransferred":
	event := new(XoOwnershipTransferred)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &OwnershipTransferred{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		NewOwner: event.NewOwner.String(),
		PreviousOwner: event.PreviousOwner.String(),
	}


	return protoMsg, nil


case "PacketReceived":
	event := new(XoPacketReceived)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &PacketReceived{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		DstAddress: event.DstAddress.String(),
		Nonce: uint64(event.Nonce),
		PayloadHash: event.PayloadHash[:],
		SrcAddress: event.SrcAddress,
		SrcChainId: uint32(event.SrcChainId),
	}


	return protoMsg, nil


case "InvalidDst":
	event := new(XoInvalidDst)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &InvalidDst{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		DstAddress: event.DstAddress.String(),
		Nonce: uint64(event.Nonce),
		PayloadHash: event.PayloadHash[:],
		SrcAddress: event.SrcAddress,
		SrcChainId: uint32(event.SrcChainId),
	}


	return protoMsg, nil


case "HashReceived":
	event := new(XoHashReceived)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &HashReceived{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		BlockData: event.BlockData[:],
		Confirmations: event.Confirmations.Bytes(),
		LookupHash: event.LookupHash[:],
		Oracle: event.Oracle.String(),
		SrcChainId: uint32(event.SrcChainId),
	}


	return protoMsg, nil


case "SetChainAddressSize":
	event := new(XoSetChainAddressSize)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &SetChainAddressSize{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		ChainId: uint32(event.ChainId),
		Size: event.Size.Bytes(),
	}


	return protoMsg, nil


case "SetLayerZeroToken":
	event := new(XoSetLayerZeroToken)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &SetLayerZeroToken{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		TokenAddress: event.TokenAddress.String(),
	}


	return protoMsg, nil


case "SetRemoteUln":
	event := new(XoSetRemoteUln)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &SetRemoteUln{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		ChainId: uint32(event.ChainId),
		Uln: event.Uln[:],
	}


	return protoMsg, nil


case "SetTreasury":
	event := new(XoSetTreasury)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &SetTreasury{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		TreasuryAddress: event.TreasuryAddress.String(),
	}


	return protoMsg, nil


case "AddInboundProofLibraryForChain":
	event := new(XoAddInboundProofLibraryForChain)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &AddInboundProofLibraryForChain{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		ChainId: uint32(event.ChainId),
		Lib: event.Lib.String(),
	}


	return protoMsg, nil


case "Packet":
	event := new(XoPacket)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &Packet{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		Payload: event.Payload,
	}


	return protoMsg, nil


case "RelayerParams":
	event := new(XoRelayerParams)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &RelayerParams{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		AdapterParams: event.AdapterParams,
		OutboundProofType: uint32(event.OutboundProofType),
	}


	return protoMsg, nil


case "SetDefaultConfigForChainId":
	event := new(XoSetDefaultConfigForChainId)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &SetDefaultConfigForChainId{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		ChainId: uint32(event.ChainId),
		InboundBlockConfirm: uint64(event.InboundBlockConfirm),
		InboundProofLib: uint32(event.InboundProofLib),
		Oracle: event.Oracle.String(),
		OutboundBlockConfirm: uint64(event.OutboundBlockConfirm),
		OutboundProofType: uint32(event.OutboundProofType),
		Relayer: event.Relayer.String(),
	}


	return protoMsg, nil


case "EnableSupportedOutboundProof":
	event := new(XoEnableSupportedOutboundProof)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &EnableSupportedOutboundProof{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		Index:       uint64(vLog.Index),

		ChainId: uint32(event.ChainId),
		ProofType: uint32(event.ProofType),
	}


	return protoMsg, nil



/* Code generated by https://github.com/trylotus/connector-bot. DO NOT EDIT */

	default:
		return nil, fmt.Errorf("invalid event: %s", ev.Name)
	}
}




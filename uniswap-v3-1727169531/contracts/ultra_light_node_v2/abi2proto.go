package ultra_light_node_v2

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
	contractAbi, err := abi.JSON(strings.NewReader(UltraLightNodeV2MetaData.ABI))
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



case "InvalidDst":
	event := new(UltraLightNodeV2InvalidDst)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &InvalidDst{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		LogIndex:       uint64(vLog.Index),

		DstAddress: event.DstAddress.Bytes(),
		Nonce: uint64(event.Nonce),
		PayloadHash: event.PayloadHash[:],
		SrcAddress: event.SrcAddress,
		SrcChainId: uint32(event.SrcChainId),
	}


	return protoMsg, nil


case "SetLayerZeroToken":
	event := new(UltraLightNodeV2SetLayerZeroToken)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &SetLayerZeroToken{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		LogIndex:       uint64(vLog.Index),

		TokenAddress: event.TokenAddress.Bytes(),
	}


	return protoMsg, nil


case "WithdrawZRO":
	event := new(UltraLightNodeV2WithdrawZRO)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &WithdrawZRO{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		LogIndex:       uint64(vLog.Index),

		Amount: event.Amount.String(),
		MsgSender: event.MsgSender.Bytes(),
		To: event.To.Bytes(),
	}


	return protoMsg, nil


case "AddInboundProofLibraryForChain":
	event := new(UltraLightNodeV2AddInboundProofLibraryForChain)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &AddInboundProofLibraryForChain{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		LogIndex:       uint64(vLog.Index),

		ChainId: uint32(event.ChainId),
		Lib: event.Lib.Bytes(),
	}


	return protoMsg, nil


case "HashReceived":
	event := new(UltraLightNodeV2HashReceived)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &HashReceived{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		LogIndex:       uint64(vLog.Index),

		BlockData: event.BlockData[:],
		Confirmations: event.Confirmations.String(),
		LookupHash: event.LookupHash[:],
		Oracle: event.Oracle.Bytes(),
		SrcChainId: uint32(event.SrcChainId),
	}


	return protoMsg, nil


case "SetDefaultAdapterParamsForChainId":
	event := new(UltraLightNodeV2SetDefaultAdapterParamsForChainId)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &SetDefaultAdapterParamsForChainId{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		LogIndex:       uint64(vLog.Index),

		AdapterParams: event.AdapterParams,
		ChainId: uint32(event.ChainId),
		ProofType: uint32(event.ProofType),
	}


	return protoMsg, nil


case "SetDefaultConfigForChainId":
	event := new(UltraLightNodeV2SetDefaultConfigForChainId)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &SetDefaultConfigForChainId{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		LogIndex:       uint64(vLog.Index),

		ChainId: uint32(event.ChainId),
		InboundBlockConfirm: uint64(event.InboundBlockConfirm),
		InboundProofLib: uint32(event.InboundProofLib),
		Oracle: event.Oracle.Bytes(),
		OutboundBlockConfirm: uint64(event.OutboundBlockConfirm),
		OutboundProofType: uint32(event.OutboundProofType),
		Relayer: event.Relayer.Bytes(),
	}


	return protoMsg, nil


case "WithdrawNative":
	event := new(UltraLightNodeV2WithdrawNative)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &WithdrawNative{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		LogIndex:       uint64(vLog.Index),

		Amount: event.Amount.String(),
		MsgSender: event.MsgSender.Bytes(),
		To: event.To.Bytes(),
	}


	return protoMsg, nil


case "EnableSupportedOutboundProof":
	event := new(UltraLightNodeV2EnableSupportedOutboundProof)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &EnableSupportedOutboundProof{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		LogIndex:       uint64(vLog.Index),

		ChainId: uint32(event.ChainId),
		ProofType: uint32(event.ProofType),
	}


	return protoMsg, nil


case "OwnershipTransferred":
	event := new(UltraLightNodeV2OwnershipTransferred)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &OwnershipTransferred{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		LogIndex:       uint64(vLog.Index),

		NewOwner: event.NewOwner.Bytes(),
		PreviousOwner: event.PreviousOwner.Bytes(),
	}


	return protoMsg, nil


case "SetRemoteUln":
	event := new(UltraLightNodeV2SetRemoteUln)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &SetRemoteUln{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		LogIndex:       uint64(vLog.Index),

		ChainId: uint32(event.ChainId),
		Uln: event.Uln[:],
	}


	return protoMsg, nil


case "RelayerParams":
	event := new(UltraLightNodeV2RelayerParams)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &RelayerParams{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		LogIndex:       uint64(vLog.Index),

		AdapterParams: event.AdapterParams,
		OutboundProofType: uint32(event.OutboundProofType),
	}


	return protoMsg, nil


case "SetChainAddressSize":
	event := new(UltraLightNodeV2SetChainAddressSize)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &SetChainAddressSize{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		LogIndex:       uint64(vLog.Index),

		ChainId: uint32(event.ChainId),
		Size: event.Size.String(),
	}


	return protoMsg, nil


case "PacketReceived":
	event := new(UltraLightNodeV2PacketReceived)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &PacketReceived{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		LogIndex:       uint64(vLog.Index),

		DstAddress: event.DstAddress.Bytes(),
		Nonce: uint64(event.Nonce),
		PayloadHash: event.PayloadHash[:],
		SrcAddress: event.SrcAddress,
		SrcChainId: uint32(event.SrcChainId),
	}


	return protoMsg, nil


case "SetTreasury":
	event := new(UltraLightNodeV2SetTreasury)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &SetTreasury{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		LogIndex:       uint64(vLog.Index),

		TreasuryAddress: event.TreasuryAddress.Bytes(),
	}


	return protoMsg, nil


case "AppConfigUpdated":
	event := new(UltraLightNodeV2AppConfigUpdated)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &AppConfigUpdated{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		LogIndex:       uint64(vLog.Index),

		ConfigType: event.ConfigType.String(),
		NewConfig: event.NewConfig,
		UserApplication: event.UserApplication.Bytes(),
	}


	return protoMsg, nil


case "Packet":
	event := new(UltraLightNodeV2Packet)
	if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
		return nil, fmt.Errorf("error unpacking event: %s", ev.Name)
	}
	protoMsg := &Packet{
		Ts:          ts,
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Bytes(),
		LogIndex:       uint64(vLog.Index),

		Payload: event.Payload,
	}


	return protoMsg, nil



/* Code generated by https://github.com/trylotus/connector-bot. DO NOT EDIT */

	default:
		return nil, fmt.Errorf("invalid event: %s", ev.Name)
	}
}

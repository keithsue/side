package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgSubmitBlockHeaderRequest{}, "btclightclient/MsgSubmitBlockHeaderRequest", nil)
	cdc.RegisterConcrete(&MsgUpdateSendersRequest{}, "btclightclient/MsgUpdateSendersRequest", nil)
	cdc.RegisterConcrete(&MsgSubmitTransactionRequest{}, "btclightclient/MsgSubmitTransactionRequest", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgSubmitBlockHeaderRequest{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgUpdateSendersRequest{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgSubmitTransactionRequest{})
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)

package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sideprotocol/side/x/btclightclient/types"
)

type msgServer struct {
	Keeper
}

// SubmitBlockHeaders implements types.MsgServer.
func (m msgServer) SubmitBlockHeaders(goCtx context.Context, msg *types.MsgSubmitBlockHeaderRequest) (*types.MsgSubmitBlockHeadersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// check if the sender is one of the authorized senders
	param := m.GetParams(ctx)
	if !param.IsAuthorizedSender(msg.Sender) {
		return nil, types.ErrSenderAddressNotAuthorized
	}

	// Set block headers
	err := m.SetBlockHeaders(ctx, msg.BlockHeaders)
	if err != nil {
		return nil, err
	}

	// Emit events
	// m.EmitEvent(
	// 	ctx,
	// 	msg.Sender,
	// 	sdk.Attribute{
	// 		Key:   types.AttributeKeyPoolCreator,
	// 		Value: msg.Sender,
	// 	},
	// )
	return &types.MsgSubmitBlockHeadersResponse{}, nil
}

// UpdateSenders implements types.MsgServer.
func (m msgServer) UpdateSenders(goCtx context.Context, msg *types.MsgUpdateSendersRequest) (*types.MsgUpdateSendersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// check if the sender is one of the authorized senders
	param := m.GetParams(ctx)
	if !param.IsAuthorizedSender(msg.Sender) {
		return nil, types.ErrSenderAddressNotAuthorized
	}

	// Set block headers
	m.SetParams(ctx, types.NewParams(msg.Senders))

	// Emit events

	return &types.MsgUpdateSendersResponse{}, nil
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}
// x/ems/keeper/msg_server.go
package keeper

import (
	"context"
	"fmt"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/rollchains/hackmoschain/x/ems/types"
)

type msgServer struct {
	k Keeper
}

var _ types.MsgServer = msgServer{}

// NewMsgServerImpl returns an implementation of the module MsgServer interface.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{k: keeper}
}

func (ms msgServer) UpdateParams(ctx context.Context, msg *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	if ms.k.authority != msg.Authority {
		return nil, errors.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", ms.k.authority, msg.Authority)
	}

	return nil, ms.k.Params.Set(ctx, msg.Params)
}

// CreateEvent implements types.MsgServer.
func (ms msgServer) CreateEvent(ctx context.Context, msg *types.MsgCreateEvent) (*types.MsgCreateEventResponse, error) {
	err := ms.k.CreateEvent(ctx, msg) // Use the updated CreateEvent function
	if err != nil {
		return nil, err
	}
	return &types.MsgCreateEventResponse{}, nil
}

// IssueEventNFT implements types.MsgServer.
func (ms msgServer) IssueEventNFT(ctx context.Context, msg *types.MsgIssueEventNFT) (*types.MsgIssueEventNFTResponse, error) {
	event, err := ms.k.GetEvent(ctx, msg.Id)
	if err != nil {
		return nil, err
	}

	isOrganizer := false
	for _, v := range event.Organizers {
		if v == msg.Organizer {
			isOrganizer = true
			break
		}
	}

	if !isOrganizer {
		return nil, fmt.Errorf("permission denied")
	}

	organizer, err := ms.k.addressCodec.StringToBytes(msg.Organizer)
	if err != nil {
		return nil, sdkerrors.ErrInvalidAddress.Wrapf("invalid organizer address: %s", err)
	}

	receiver, err := ms.k.addressCodec.StringToBytes(msg.Receiver)
	if err != nil {
		return nil, sdkerrors.ErrInvalidAddress.Wrapf("invalid receiver address: %s", err)
	}

	err = ms.k.MintEventNFT(ctx, organizer, receiver, msg.Id)
	if err != nil {
		return nil, err
	}

	return &types.MsgIssueEventNFTResponse{}, nil
}

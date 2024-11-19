package keeper

import (
	"context"
	"fmt"

	"gluon/x/customauth/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreatePairing(goCtx context.Context, msg *types.MsgCreatePairing) (*types.MsgCreatePairingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var pairing = types.Pairing{
		User:      msg.User,
		Address:   msg.Address,
		PublicKey: msg.PublicKey,
	}

	id := k.AppendPairing(
		ctx,
		pairing,
	)

	return &types.MsgCreatePairingResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdatePairing(goCtx context.Context, msg *types.MsgUpdatePairing) (*types.MsgUpdatePairingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var pairing = types.Pairing{
		User:      msg.User,
		Id:        msg.Id,
		Address:   msg.Address,
		PublicKey: msg.PublicKey,
	}

	// Checks that the element exists
	val, found := k.GetPairing(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg user is the same as the current owner
	if msg.User != val.User {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetPairing(ctx, pairing)

	return &types.MsgUpdatePairingResponse{}, nil
}

func (k msgServer) DeletePairing(goCtx context.Context, msg *types.MsgDeletePairing) (*types.MsgDeletePairingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetPairing(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg user is the same as the current owner
	if msg.User != val.User {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemovePairing(ctx, msg.Id)

	return &types.MsgDeletePairingResponse{}, nil
}

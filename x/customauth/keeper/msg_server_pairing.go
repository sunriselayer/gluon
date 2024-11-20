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
		Address:   msg.User,
		PublicKey: msg.PublicKey,
	}

	id, err := k.AppendPairing(
		ctx,
		pairing,
	)
	if err != nil {
		return nil, err
	}

	return &types.MsgCreatePairingResponse{
		Id: id,
	}, nil
}

func (k msgServer) DeletePairing(goCtx context.Context, msg *types.MsgDeletePairing) (*types.MsgDeletePairingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetPairing(ctx, msg.User, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg user is the same as the current owner
	if msg.User != val.Address {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemovePairing(ctx, msg.User, msg.Id)

	return &types.MsgDeletePairingResponse{}, nil
}

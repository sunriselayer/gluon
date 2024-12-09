package keeper

import (
	"context"
	"fmt"

	"gluon/x/customauth/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
)

func (k msgServer) CreatePairing(goCtx context.Context, msg *types.MsgCreatePairing) (*types.MsgCreatePairingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var pubKey cryptotypes.PubKey
	err := k.cdc.UnpackAny(&msg.PublicKey, &pubKey)
	if err != nil {
		return nil, err
	}

	index := ""

	var pairing = types.Pairing{
		Owner:     msg.User,
		Index:     index,
		PublicKey: msg.PublicKey,
	}

	k.SetPairing(
		ctx,
		pairing,
	)

	return &types.MsgCreatePairingResponse{
		PairingIndex: index,
	}, nil
}

func (k msgServer) DeletePairing(goCtx context.Context, msg *types.MsgDeletePairing) (*types.MsgDeletePairingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetPairing(ctx, msg.User, msg.PairingIndex)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("pairing of address %s, index %s doesn't exist", msg.User, msg.PairingIndex))
	}

	// Checks if the msg user is the same as the current owner
	if msg.User != val.Owner {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemovePairing(ctx, msg.User, msg.PairingIndex)

	return &types.MsgDeletePairingResponse{}, nil
}

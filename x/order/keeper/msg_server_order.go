package keeper

import (
	"context"

	"gluon/x/order/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdkmath "cosmossdk.io/math"

	gogo "github.com/gogo/protobuf/types"
)

func (k msgServer) LazyRegisterOrder(goCtx context.Context, msg *types.MsgLazyRegisterOrder) (*types.MsgLazyRegisterOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	hash, err := types.GetOrderHash(msg.Order)
	if err != nil {
		return nil, err
	}

	// Check if the value already exists
	_, isFound := k.GetOrder(
		ctx,
		msg.User,
		hash,
	)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	pairing, found := k.customAuthKeeper.GetPairing(ctx, msg.User, msg.PairingId)
	if !found {
		return nil, types.ErrPairingNotFound
	}
	pubKey, err := k.customAuthKeeper.GetPairingPubKey(ctx, pairing)
	if err != nil {
		return nil, err
	}
	orderBody, err := types.UnpackOrderAny(k.cdc, msg.Order)
	if err != nil {
		return nil, err
	}
	err = orderBody.ValidateBasic()
	if err != nil {
		return nil, err
	}
	err = types.VerifyOrderSignature(orderBody, pubKey, msg.Signature)
	if err != nil {
		return nil, err
	}

	k.SetOrder(
		ctx,
		types.Order{
			Hash:             hash,
			Owner:            msg.User,
			Body:             msg.Order,
			Cancelled:        false,
			ContractedAmount: sdkmath.ZeroInt(),
		},
	)

	expiry := orderBody.GetExpiry()
	timestamp, err := gogo.TimestampProto(expiry)
	if err != nil {
		return nil, err
	}

	k.SetSortedOrder(ctx, types.SortedOrder{
		Seconds:   uint64(timestamp.Seconds),
		Nanos:     uint32(timestamp.Nanos),
		OrderHash: hash,
	})
	return &types.MsgLazyRegisterOrderResponse{}, nil
}

func (k msgServer) CancelOrder(goCtx context.Context, msg *types.MsgCancelOrder) (*types.MsgCancelOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetOrder(
		ctx,
		msg.User,
		msg.OrderHash,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the msg user is the same as the current owner
	if msg.User != valFound.Owner {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	if valFound.Cancelled {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "order already cancelled")
	}

	valFound.Cancelled = true
	k.SetOrder(ctx, valFound)

	return &types.MsgCancelOrderResponse{}, nil
}

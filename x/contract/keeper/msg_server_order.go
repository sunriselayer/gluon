package keeper

import (
	"context"

	"gluon/x/contract/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdkmath "cosmossdk.io/math"
)

func (k msgServer) LazyRegisterOrder(goCtx context.Context, msg *types.MsgLazyRegisterOrder) (*types.MsgLazyRegisterOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetOrder(
		ctx,
		msg.Order.Id,
	)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	err := k.ValidateOrder(ctx, msg.Order, msg.PairingId, msg.Signature)
	if err != nil {
		return nil, err
	}

	k.SetOrder(
		ctx,
		msg.Order,
	)
	k.SetSortedOrder(ctx, types.SortedOrder{
		Expiry:           uint64(msg.Order.Expiry.UnixMilli()),
		Id:               msg.Order.Id,
		Cancelled:        false,
		ContractedAmount: sdkmath.ZeroInt(),
	})
	return &types.MsgLazyRegisterOrderResponse{}, nil
}

func (k msgServer) CancelOrder(goCtx context.Context, msg *types.MsgCancelOrder) (*types.MsgCancelOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetOrder(
		ctx,
		msg.Id,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the msg user is the same as the current owner
	if msg.User != valFound.Address {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	sortedOrder, isFound := k.GetSortedOrder(ctx, uint64(valFound.Expiry.UnixMilli()), valFound.Id)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "sorted order not found")
	}

	sortedOrder.Cancelled = true
	k.SetSortedOrder(ctx, sortedOrder)

	return &types.MsgCancelOrderResponse{}, nil
}

package keeper

import (
	"gluon/x/order/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetOrderAndBody(ctx sdk.Context, orderOwner string, orderHash string) (types.Order, types.OrderBody, error) {

	order, found := k.GetOrder(ctx, orderOwner, orderHash)
	if !found {
		return types.Order{}, nil, types.ErrOrderNotFound
	}

	orderBody, err := types.UnpackOrderAny(k.cdc, order.Body)
	if err != nil {
		return types.Order{}, nil, err
	}

	return order, orderBody, nil
}

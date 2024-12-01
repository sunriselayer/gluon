package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	ordertypes "gluon/x/order/types"
	"gluon/x/perp/types"
)

func (k Keeper) GetOrderBodyDenomsDirection(ctx context.Context, orderBody ordertypes.OrderBody) (perpOrder types.PerpOrder, denomBase, denomQuote string, direction ordertypes.OrderDirection, err error) {
	switch orderBody := orderBody.(type) {
	case *types.PerpPositionCreateOrder:
		perpOrder = orderBody
		denomBase = orderBody.DenomBase
		denomQuote = orderBody.DenomQuote
		direction = orderBody.Direction
		return
	case *types.PerpPositionCancelOrder:
		perpOrder = orderBody
		position, found := k.GetPosition(ctx, orderBody.AddressString, orderBody.PositionId)
		if !found {
			err = errorsmod.Wrapf(sdkerrors.ErrNotFound, "position id: %d", orderBody.PositionId)
			return
		}
		denomBase = position.DenomBase
		denomQuote = position.DenomQuote

		switch position.Direction {
		case types.PositionDirection_POSITION_DIRECTION_LONG:
			direction = ordertypes.OrderDirection_ORDER_DIRECTION_SELL
		case types.PositionDirection_POSITION_DIRECTION_SHORT:
			direction = ordertypes.OrderDirection_ORDER_DIRECTION_BUY
		default:
			err = errorsmod.Wrapf(ordertypes.ErrInvalidOrderDirection, "position direction is invalid")
			return
		}
		return
	default:
		err = types.ErrInvalidOrderType
		return
	}
}

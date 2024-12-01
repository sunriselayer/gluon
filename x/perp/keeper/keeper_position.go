package keeper

import (
	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	ordertypes "gluon/x/order/types"
	types "gluon/x/perp/types"
)

func (k Keeper) CreateUpdateCancelPosition(
	ctx sdk.Context,
	order types.PerpOrder,
	amount sdkmath.Int,
	price sdkmath.LegacyDec,
) error {
	switch order := order.(type) {
	case *types.PerpPositionCreateOrder:
		return k.CreateUpdatePosition(ctx, *order, amount, price)
	case *types.PerpPositionCancelOrder:
		return k.CancelPosition(ctx, *order, amount, price)
	default:
		return types.ErrInvalidOrderType
	}
}

func (k Keeper) CreateUpdatePosition(
	ctx sdk.Context,
	order types.PerpPositionCreateOrder,
	amount sdkmath.Int,
	price sdkmath.LegacyDec,
) error {
	var direction types.PositionDirection
	switch order.Direction {
	case ordertypes.OrderDirection_ORDER_DIRECTION_BUY:
		direction = types.PositionDirection_POSITION_DIRECTION_LONG
	case ordertypes.OrderDirection_ORDER_DIRECTION_SELL:
		direction = types.PositionDirection_POSITION_DIRECTION_SHORT
	default:
		return ordertypes.ErrInvalidOrderDirection
	}

	positionId := k.AppendPosition(ctx, types.Position{
		Owner:          order.AddressString,
		DenomBase:      order.DenomBase,
		DenomQuote:     order.DenomQuote,
		Direction:      direction,
		IsolatedMargin: order.IsolatedMargin,
	})

	owner, err := sdk.AccAddressFromBech32(order.AddressString)
	if err != nil {
		return err
	}

	var marginAddressModule string
	if order.IsolatedMargin {
		marginAddressModule = types.GetIsolatedMarginAddressModule(order.AddressString, positionId)
	} else {
		marginAddressModule = types.GetCrossMarginAddressModule(order.AddressString)
	}

	coins := sdk.NewCoins(sdk.NewCoin(order.DenomQuote, order.Margin.Amount))

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, owner, marginAddressModule, coins)
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) CancelPosition(
	ctx sdk.Context,
	order types.PerpPositionCancelOrder,
	amount sdkmath.Int,
	price sdkmath.LegacyDec,
) error {
	position, found := k.GetPosition(ctx, order.AddressString, order.PositionId)
	if !found {
		return errorsmod.Wrapf(sdkerrors.ErrNotFound, "position id: %d", order.PositionId)
	}

	owner, err := sdk.AccAddressFromBech32(order.AddressString)
	if err != nil {
		return err
	}

	// TODO
	remainingAmount := sdkmath.ZeroInt()

	if amount.GT(remainingAmount) {
		return types.ErrPositionCancelAmountExceed
	} else if amount.LT(remainingAmount) {

	} else {
		k.RemovePosition(ctx, order.AddressString, order.PositionId)

		if position.IsolatedMargin {
			marginAddressModule := types.GetIsolatedMarginAddressModule(order.AddressString, order.PositionId)
			coins := k.bankKeeper.SpendableCoins(ctx, k.accountKeeper.GetModuleAddress(marginAddressModule))

			err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, owner, marginAddressModule, coins)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

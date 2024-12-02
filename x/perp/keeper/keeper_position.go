package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	ordertypes "gluon/x/order/types"
	types "gluon/x/perp/types"
)

func (k Keeper) CreateUpdateCancelPosition(
	ctx sdk.Context,
	orderHash string,
	order types.PerpOrder,
	price sdkmath.LegacyDec,
	quantity sdkmath.Int,
) error {
	switch order := order.(type) {
	case *types.PerpPositionCreateOrder:
		return k.CreateUpdatePosition(ctx, orderHash, *order, price, quantity)

	case *types.PerpPositionCancelOrder:
		return k.CancelPosition(ctx, *order, price, quantity)
	default:
		return types.ErrInvalidOrderType
	}
}

func (k Keeper) CreateUpdatePosition(
	ctx sdk.Context,
	orderHash string,
	order types.PerpPositionCreateOrder,
	price sdkmath.LegacyDec,
	quantity sdkmath.Int,
) error {
	position, found := k.GetPosition(ctx, order.AddressString, orderHash)
	if found {
		// Update
		position = k.DeductFundingFee(ctx, position)

		// Add Position Amount
		position.Amount = position.Amount.Add(quantity)
		k.SetPosition(ctx, position)

		// Add Price Quantity
		priceQuantity, found := k.GetPositionPriceQuantity(ctx, order.AddressString, orderHash, price.String())
		if found {
			priceQuantity.Quantity = priceQuantity.Quantity.Add(quantity)
			k.SetPositionPriceQuantity(ctx, priceQuantity)
		} else {
			k.SetPositionPriceQuantity(ctx, types.PositionPriceQuantity{
				Owner:             order.AddressString,
				PositionOrderHash: orderHash,
				Price:             price.String(),
				Quantity:          quantity,
			})
		}

		// No need for margin transfer
		return nil
	}

	return k.CreatePosition(ctx, orderHash, order, price, quantity)
}

func (k Keeper) CreatePosition(
	ctx sdk.Context,
	orderHash string,
	order types.PerpPositionCreateOrder,
	price sdkmath.LegacyDec,
	quantity sdkmath.Int,
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

	// Transfer margin
	owner, err := sdk.AccAddressFromBech32(order.AddressString)
	if err != nil {
		return err
	}

	coins := sdk.NewCoins(sdk.NewCoin(order.DenomQuote, order.MarginAmount))

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, owner, types.GetMarginAddressModule(), coins)
	if err != nil {
		return err
	}

	isolatedMarginAmount := sdkmath.ZeroInt()
	if order.IsolatedMargin {
		isolatedMarginAmount = order.MarginAmount
	}

	// Store
	k.SetPosition(ctx, types.Position{
		Owner:                order.AddressString,
		OrderHash:            orderHash,
		DenomBase:            order.DenomBase,
		DenomQuote:           order.DenomQuote,
		Direction:            direction,
		Amount:               quantity,
		IsolatedMargin:       order.IsolatedMargin,
		IsolatedMarginAmount: isolatedMarginAmount,
		FundingFeeResetAt:    ctx.BlockTime(),
	})
	k.SetPositionPriceQuantity(ctx, types.PositionPriceQuantity{
		Owner:             order.AddressString,
		PositionOrderHash: orderHash,
		Price:             price.String(),
		Quantity:          quantity,
	})

	return nil
}

func (k Keeper) CancelPosition(
	ctx sdk.Context,
	order types.PerpPositionCancelOrder,
	price sdkmath.LegacyDec,
	quantity sdkmath.Int,
) error {
	position, found := k.GetPosition(ctx, order.AddressString, order.PositionOrderHash)
	if !found {
		return errorsmod.Wrapf(sdkerrors.ErrNotFound, "position owner: %s, order hash: %s", position.Owner, position.OrderHash)
	}
	position = k.DeductFundingFee(ctx, position)

	owner, err := sdk.AccAddressFromBech32(order.AddressString)
	if err != nil {
		return err
	}

	if quantity.GT(position.Amount) {
		return types.ErrPositionCancelAmountExceed
	}

	deleteLater := position.Amount.Equal(quantity)

	pnl := price.MulInt(quantity).TruncateInt()

	// Process PositionPriceQuantity and get PnL
	err = k.iteratePositionPriceQuantityByOwnerAndPositionOrderHash(ctx, order.AddressString, order.PositionOrderHash, func(val types.PositionPriceQuantity) (finish bool, err error) {
		var sub sdkmath.Int
		if quantity.GT(val.Quantity) {
			sub = val.Quantity
		} else {
			sub = quantity
		}

		val.Quantity = val.Quantity.Sub(sub)
		if val.Quantity.IsZero() {
			k.RemovePositionPriceQuantity(ctx, order.AddressString, order.PositionOrderHash, val.Price)
		} else {
			k.SetPositionPriceQuantity(ctx, val)
		}

		valPrice, err := sdkmath.LegacyNewDecFromStr(val.Price)
		if err != nil {
			return true, err
		}

		pnl = pnl.Sub(valPrice.MulInt(sub).TruncateInt())

		quantity = quantity.Sub(sub)
		return quantity.IsZero(), nil
	})
	if err != nil {
		return err
	}
	if position.Direction == types.PositionDirection_POSITION_DIRECTION_SHORT {
		pnl = pnl.Neg()
	}

	// Transfer PnL
	if position.IsolatedMargin {
		position.IsolatedMarginAmount = position.IsolatedMarginAmount.Add(pnl)
	} else {
		coin := sdk.NewCoin(position.DenomQuote, pnl)
		margins := k.AddCrossMargin(ctx, order.AddressString, coin)
		if margins.IsAnyNegative() {
			// TODO: emit event
		}
	}

	// Process Position
	if deleteLater {
		if position.IsolatedMargin {
			if position.IsolatedMarginAmount.IsPositive() {
				coins := sdk.NewCoins(sdk.NewCoin(position.DenomQuote, position.IsolatedMarginAmount))
				err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.GetMarginAddressModule(), owner, coins)
				if err != nil {
					return err
				}
			} else {
				// TODO: emit event
			}
		}
		k.RemovePosition(ctx, order.AddressString, order.PositionOrderHash)
	} else {
		position.Amount = position.Amount.Sub(quantity)
		k.SetPosition(ctx, position)
	}

	return nil
}

func (k Keeper) iteratePositionPriceQuantityByOwnerAndPositionOrderHash(ctx context.Context, owner string, positionOrderHash string, callback func(val types.PositionPriceQuantity) (finish bool, err error)) error {
	iterator := k.GetPositionPriceQuantitiesByOwnerAndPositionOrderHashIterator(ctx, owner, positionOrderHash)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var priceQuantity types.PositionPriceQuantity
		err := k.cdc.Unmarshal(iterator.Value(), &priceQuantity)
		if err != nil {
			return err
		}

		finish, err := callback(priceQuantity)
		if err != nil {
			return err
		}

		if finish {
			break
		}
	}

	return nil
}

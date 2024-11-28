package keeper

import (
	"context"
	"gluon/x/contract/types"
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) Contract(
	ctx sdk.Context,
	buy types.Order,
	sell types.Order,
	amount sdkmath.Int,
	price sdkmath.LegacyDec,
) error {
	denomBase := buy.DenomBase
	denomQuote := buy.DenomQuote

	addressBuy, err := sdk.AccAddressFromBech32(buy.Address)
	if err != nil {
		return err
	}
	addressSell, err := sdk.AccAddressFromBech32(sell.Address)
	if err != nil {
		return err
	}

	amountBase := amount
	amountQuote := price.MulInt(amount).RoundInt()

	if amountQuote.IsPositive() {
		err = k.bankKeeper.SendCoins(
			ctx,
			addressBuy,
			addressSell,
			sdk.NewCoins(
				sdk.NewCoin(denomQuote, amountQuote),
			),
		)
		if err != nil {
			return err
		}
	}

	if amountBase.IsPositive() {
		err = k.bankKeeper.SendCoins(
			ctx,
			addressSell,
			addressBuy,
			sdk.NewCoins(
				sdk.NewCoin(denomBase, amountBase),
			),
		)
		if err != nil {
			return err
		}
	}

	err = k.AddContractedAmount(ctx, buy.Expiry, buy.Id, amount)
	if err != nil {
		return err
	}
	err = k.AddContractedAmount(ctx, sell.Expiry, sell.Id, amount)
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) AddContractedAmount(ctx context.Context, expiry time.Time, id string, amount sdkmath.Int) error {
	sortedOrder, found := k.GetSortedOrder(ctx, expiry, id)
	if !found {
		return types.ErrOrderNotFound
	}

	sortedOrder.ContractedAmount = sortedOrder.ContractedAmount.Add(amount)
	k.SetSortedOrder(ctx, sortedOrder)

	return nil
}

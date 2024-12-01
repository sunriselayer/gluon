package keeper

import (
	"gluon/x/spot/types"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) Swap(
	ctx sdk.Context,
	buy types.SpotOrder,
	sell types.SpotOrder,
	amount sdkmath.Int,
	price sdkmath.LegacyDec,
) error {
	denomBase := buy.DenomBase
	denomQuote := buy.DenomQuote

	addressBuy, err := sdk.AccAddressFromBech32(buy.AddressString)
	if err != nil {
		return err
	}
	addressSell, err := sdk.AccAddressFromBech32(sell.AddressString)
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

	return nil
}

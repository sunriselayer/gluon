package keeper

import (
	"gluon/x/contract/types"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) Settle(
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

	err = k.bankKeeper.SendCoins(
		ctx,
		addressBuy,
		addressSell,
		sdk.NewCoins(
			sdk.NewCoin(denomQuote, price.MulInt(amount).RoundInt()),
		),
	)
	if err != nil {
		return err
	}

	err = k.bankKeeper.SendCoins(
		ctx,
		addressSell,
		addressBuy,
		sdk.NewCoins(
			sdk.NewCoin(denomBase, amount),
		),
	)
	if err != nil {
		return err
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

// TODO
func (k Keeper) SettleLazy(
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

	_ = denomBase
	_ = denomQuote
	_ = addressBuy
	_ = addressSell
	_ = price

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

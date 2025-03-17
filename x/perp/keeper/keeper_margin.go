package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"gluon/x/perp/types"
)

// AddCrossMargin adds cross margin asset to the store
// BankKeeper.SendCoins is not included
func (k Keeper) AddCrossMargin(ctx context.Context, owner string, asset sdk.Coin) sdk.Coins {
	crossMargin, found := k.GetCrossMargin(ctx, owner)

	if found {
		crossMargin.Assets = crossMargin.Assets.Add(asset)
		k.SetCrossMargin(ctx, crossMargin)

		return crossMargin.Assets
	}

	crossMargin = types.CrossMargin{
		Owner:  owner,
		Assets: sdk.NewCoins(asset),
	}

	k.SetCrossMargin(ctx, crossMargin)

	return crossMargin.Assets
}

func (k Keeper) depositMargin(ctx context.Context, user string, assets sdk.Coins) error {
	address, err := sdk.AccAddressFromBech32(user)
	if err != nil {
		return err
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, address, types.GetMarginAddressModule(), assets)
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) withdrawMargin(ctx context.Context, user string, assets sdk.Coins) error {
	address, err := sdk.AccAddressFromBech32(user)
	if err != nil {
		return err
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.GetMarginAddressModule(), address, assets)
	if err != nil {
		return err
	}

	return nil
}

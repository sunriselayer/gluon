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

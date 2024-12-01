package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	sdkmath "cosmossdk.io/math"
	types "gluon/x/perp/types"
)

func (k Keeper) CreatePosition(
	ctx sdk.Context,
	buy types.PerpOrder,
	sell types.PerpOrder,
	amount sdkmath.Int,
	price sdkmath.LegacyDec,
) error {
	return nil
}

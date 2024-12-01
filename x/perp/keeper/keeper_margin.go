package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	types "gluon/x/perp/types"
)

func (k Keeper) GetCrossMarginAddress(ctx sdk.Context, positionOwner string) sdk.AccAddress {
	return k.accountKeeper.GetModuleAddress(fmt.Sprintf("%s/margin/%s", types.ModuleName, positionOwner))
}

func (k Keeper) GetIsolatedMarginAddress(ctx sdk.Context, positionOwner string, positionId uint64) sdk.AccAddress {
	return k.accountKeeper.GetModuleAddress(fmt.Sprintf("%s/margin/%s/%d", types.ModuleName, positionOwner, positionId))
}

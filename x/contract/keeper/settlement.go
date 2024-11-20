package keeper

import (
	"gluon/x/contract/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) Settle(ctx sdk.Context, earlier types.Order, later types.Order) error {
	earlierAddress, err := sdk.AccAddressFromBech32(earlier.Body.Address)
	if err != nil {
		return err
	}
	laterAddress, err := sdk.AccAddressFromBech32(later.Body.Address)
	if err != nil {
		return err
	}

	err = k.bankKeeper.SendCoins(ctx, earlierAddress, laterAddress, sdk.NewCoins())
	if err != nil {
		return err
	}

	err = k.bankKeeper.SendCoins(ctx, laterAddress, earlierAddress, sdk.NewCoins())
	if err != nil {
		return err
	}

	return nil
}

package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"gluon/x/perp/types"
)

func (k Keeper) DeductFundingFee(ctx context.Context, position types.Position) types.Position {
	fundingRates := k.GetFundingRatesByPositionReverse(ctx, position)
	fee := types.GetFundingFee(position, fundingRates)

	if position.IsolatedMargin != nil {
		subtracted := position.IsolatedMargin.Sub(fee.Amount)
		position.IsolatedMargin = &subtracted
	} else {
		fee.Amount = fee.Amount.Neg()
		k.AddCrossMargin(ctx, position.Owner, fee)
	}

	position.FundingFeeResetAt = sdk.UnwrapSDKContext(ctx).BlockTime()
	k.SetPosition(ctx, position)

	// TODO: emit event

	return position
}

package keeper

import (
	"context"
	"gluon/x/contract/types"

	sdkmath "cosmossdk.io/math"
)

func (k Keeper) AddContractedAmount(ctx context.Context, orderHash string, additionalContractedAmount sdkmath.Int) error {
	order, found := k.GetOrder(ctx, orderHash)
	if !found {
		return types.ErrOrderNotFound
	}

	order.ContractedAmount = order.ContractedAmount.Add(additionalContractedAmount)
	k.SetOrder(ctx, order)

	return nil
}

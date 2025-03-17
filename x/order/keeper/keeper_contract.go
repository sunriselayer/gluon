package keeper

import (
	"context"
	"gluon/x/order/types"

	sdkmath "cosmossdk.io/math"
)

func (k Keeper) AddContractedAmount(ctx context.Context, orderOwner string, orderHash string, additionalContractedAmount sdkmath.Int) error {
	order, found := k.GetOrder(ctx, orderOwner, orderHash)
	if !found {
		return types.ErrOrderNotFound
	}

	order.ContractedAmount = order.ContractedAmount.Add(additionalContractedAmount)
	k.SetOrder(ctx, order)

	return nil
}

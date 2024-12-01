package keeper

import (
	"gluon/x/order/types"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) ValidateContractAmount(ctx sdk.Context, orderHash string, additionalAmount sdkmath.Int) error {

	order, found := k.GetOrder(ctx, orderHash)
	if !found {
		return types.ErrOrderNotFound
	}

	orderBody, err := types.UnpackOrderAny(k.cdc, order.Body)
	if err != nil {
		return err
	}

	amount := orderBody.GetAmount()

	if order.ContractedAmount.Add(additionalAmount).GT(amount) {
		return types.ErrContractAmountExceed
	}

	return nil
}

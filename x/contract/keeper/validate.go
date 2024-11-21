package keeper

import (
	"gluon/x/contract/types"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) ValidateOrder(ctx sdk.Context, order types.Order, pairingId uint64, signature []byte) error {
	pairing, found := k.customAuthKeeper.GetPairing(ctx, order.Address, pairingId)
	if !found {
		return types.ErrPairingNotFound
	}
	pubKey, err := k.customAuthKeeper.GetPairingPubKey(ctx, pairing)
	if err != nil {
		return err
	}
	err = order.VerifySignature(pubKey, signature)
	if err != nil {
		return err
	}
	return nil
}

func (k Keeper) ValidateContractAmount(ctx sdk.Context, order types.Order, amount sdkmath.Int) error {
	sortedOrder, found := k.GetSortedOrder(ctx, uint64(order.Expiry.UnixMilli()), order.Id)
	if !found {
		return types.ErrOrderNotFound
	}

	if sortedOrder.ContractedAmount.Add(amount).GT(order.Amount) {
		return types.ErrContractAmountExceed
	}

	return nil
}

package keeper

import (
	"gluon/x/contract/types"

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
	err = order.Validate(pubKey, signature)
	if err != nil {
		return err
	}
	return nil
}

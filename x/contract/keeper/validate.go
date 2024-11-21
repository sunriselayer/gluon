package keeper

import (
	"gluon/x/contract/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) ValidateOrder(ctx sdk.Context, order types.Order) error {
	pairing, found := k.customAuthKeeper.GetPairing(ctx, order.Body.Address, order.ParingId)
	if !found {
		return types.ErrPairingNotFound
	}
	pubKey, err := k.customAuthKeeper.GetPairingPubKey(ctx, pairing)
	if err != nil {
		return err
	}
	err = order.Validate(pubKey)
	if err != nil {
		return err
	}
	return nil
}

func (k Keeper) ValidateOrderPair(ctx sdk.Context, earlier types.Order, later types.Order) error {
	pairingEarlier, found := k.customAuthKeeper.GetPairing(ctx, earlier.Body.Address, earlier.ParingId)
	if !found {
		return types.ErrPairingNotFound
	}
	pairingLater, found := k.customAuthKeeper.GetPairing(ctx, later.Body.Address, later.ParingId)
	if !found {
		return types.ErrPairingNotFound
	}

	pubKeyEarlier, err := k.customAuthKeeper.GetPairingPubKey(ctx, pairingEarlier)
	if err != nil {
		return err
	}
	pubKeyLater, err := k.customAuthKeeper.GetPairingPubKey(ctx, pairingLater)
	if err != nil {
		return err
	}

	err = earlier.Validate(pubKeyEarlier)
	if err != nil {
		return err
	}
	err = later.Validate(pubKeyLater)
	if err != nil {
		return err
	}

	err = earlier.CrossValidate(later)
	if err != nil {
		return err
	}

	return nil
}

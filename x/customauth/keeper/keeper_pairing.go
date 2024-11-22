package keeper

import (
	"context"

	"gluon/x/customauth/types"

	"gluon/x/customauth/types/pairing"

	errorsmod "cosmossdk.io/errors"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) getPairingPubKey(pairing types.Pairing) (cryptotypes.PubKey, error) {
	var pubKey cryptotypes.PubKey
	err := k.cdc.UnpackAny(&pairing.PublicKey, &pubKey)
	if err != nil {
		return nil, err
	}

	return pubKey, nil
}

func (k Keeper) GetPairingPubKey(goCtx context.Context, user sdk.AccAddress, pairingId uint64) (*pairing.PubKeyInternal, error) {
	pairingVal, found := k.GetPairing(goCtx, user.String(), pairingId)
	if !found {
		return nil, errorsmod.Wrapf(types.ErrPairingNotFound, "address: %s, pairing_id: %d", user.String(), pairingId)
	}

	params := k.GetParams(goCtx)

	ctx := sdk.UnwrapSDKContext(goCtx)

	if ctx.BlockTime().Compare(pairingVal.CreatedAt.Add(params.ParingDelay)) < 0 {
		return nil, errorsmod.Wrapf(types.ErrPairingDelayPeriod, "address: %s, pairing_id: %d", user.String(), pairingId)
	}

	pubKey := pairing.PubKeyInternal{
		User:              user,
		PairingPublicKey:  pairingVal.PublicKey,
		OperatorPublicKey: params.OperatorPublicKey,
	}

	// IMPORTANT: To create cache
	// because pairing.PubKey depends on cache
	var buffer cryptotypes.PubKey
	if err := k.cdc.UnpackAny(&pubKey.PairingPublicKey, &buffer); err != nil {
		return nil, err
	}
	if err := k.cdc.UnpackAny(&pubKey.OperatorPublicKey, &buffer); err != nil {
		return nil, err
	}

	return &pubKey, nil
}

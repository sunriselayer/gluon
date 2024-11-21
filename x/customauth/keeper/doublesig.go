package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"gluon/x/customauth/types"
	"gluon/x/customauth/types/doublesig"
)

func (k Keeper) GetDoubleSigPubKey(goCtx context.Context, userAddress string, pairingId uint64) (*doublesig.PubKey, error) {
	pairing, found := k.GetPairing(goCtx, userAddress, pairingId)
	if !found {
		return nil, errorsmod.Wrapf(types.ErrPairingNotFound, "address: %s, pairing_id: %d", userAddress, pairingId)
	}
	params := k.GetParams(goCtx)

	ctx := sdk.UnwrapSDKContext(goCtx)

	if ctx.BlockTime().Compare(pairing.CreatedAt.Add(params.ParingDelay)) < 0 {
		return nil, errorsmod.Wrapf(types.ErrPairingDelayPeriod, "address: %s, pairing_id: %d", userAddress, pairingId)
	}

	pubKey := doublesig.PubKey{
		User:     pairing.PublicKey,
		Operator: params.OperatorPublicKey,
	}

	return &pubKey, nil
}

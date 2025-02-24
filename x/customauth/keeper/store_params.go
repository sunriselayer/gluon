package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/runtime"

	"gluon/x/customauth/types"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx context.Context) (params types.Params) {
	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	bz := store.Get(types.ParamsKey)
	if bz == nil {
		return params
	}

	k.cdc.MustUnmarshal(bz, &params)
	return params
}

// SetParams set the params
func (k Keeper) SetParams(ctx context.Context, params types.Params) error {
	// <gluon>
	var pubKey cryptotypes.PubKey
	err := k.cdc.UnpackAny(&params.OperatorPublicKey, &pubKey)
	if err != nil {
		return err
	}
	// </gluon>

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	bz, err := k.cdc.Marshal(&params)
	if err != nil {
		return err
	}
	store.Set(types.ParamsKey, bz)

	return nil
}

func (k Keeper) GetOperatorPubKey(ctx context.Context) (cryptotypes.PubKey, error) {
	params := k.GetParams(ctx)

	var pubKey cryptotypes.PubKey
	err := k.cdc.UnpackAny(&params.OperatorPublicKey, &pubKey)
	if err != nil {
		return nil, err
	}

	return pubKey, nil
}

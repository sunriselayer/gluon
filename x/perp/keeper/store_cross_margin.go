package keeper

import (
	"context"

	"gluon/x/perp/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetCrossMargin set a specific crossMargin in the store from its index
func (k Keeper) SetCrossMargin(ctx context.Context, crossMargin types.CrossMargin) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.CrossMarginKeyPrefix))
	b := k.cdc.MustMarshal(&crossMargin)
	store.Set(types.CrossMarginKey(
		crossMargin.Owner,
	), b)
}

// GetCrossMargin returns a crossMargin from its index
func (k Keeper) GetCrossMargin(
	ctx context.Context,
	owner string,

) (val types.CrossMargin, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.CrossMarginKeyPrefix))

	b := store.Get(types.CrossMarginKey(
		owner,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveCrossMargin removes a crossMargin from the store
func (k Keeper) RemoveCrossMargin(
	ctx context.Context,
	owner string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.CrossMarginKeyPrefix))
	store.Delete(types.CrossMarginKey(
		owner,
	))
}

// GetAllCrossMargin returns all crossMargin
func (k Keeper) GetAllCrossMargin(ctx context.Context) (list []types.CrossMargin) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.CrossMarginKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.CrossMargin
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

package keeper

import (
	"context"

	"gluon/x/perp/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetPositionPriceQuantity set a specific positionPriceQuantity in the store from its index
func (k Keeper) SetPositionPriceQuantity(ctx context.Context, positionPriceQuantity types.PositionPriceQuantity) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PositionPriceQuantityKeyPrefix))
	b := k.cdc.MustMarshal(&positionPriceQuantity)
	store.Set(types.PositionPriceQuantityKey(
		positionPriceQuantity.Owner,
		positionPriceQuantity.PositionOrderHash,
		positionPriceQuantity.Price,
	), b)
}

// GetPositionPriceQuantity returns a positionPriceQuantity from its index
func (k Keeper) GetPositionPriceQuantity(
	ctx context.Context,
	owner string,
	positionOrderHash string,
	price string,

) (val types.PositionPriceQuantity, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PositionPriceQuantityKeyPrefix))

	b := store.Get(types.PositionPriceQuantityKey(
		owner,
		positionOrderHash,
		price,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePositionPriceQuantity removes a positionPriceQuantity from the store
func (k Keeper) RemovePositionPriceQuantity(
	ctx context.Context,
	owner string,
	positionOrderHash string,
	price string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PositionPriceQuantityKeyPrefix))
	store.Delete(types.PositionPriceQuantityKey(
		owner,
		positionOrderHash,
		price,
	))
}

// GetAllPositionPriceQuantity returns all positionPriceQuantity
func (k Keeper) GetAllPositionPriceQuantity(ctx context.Context) (list []types.PositionPriceQuantity) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PositionPriceQuantityKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PositionPriceQuantity
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetPositionPriceQuantity returns all positionPriceQuantity
func (k Keeper) GetPositionPriceQuantitiesByOwnerAndPositionOrderHash(ctx context.Context, owner string, positionOrderHash string) (list []types.PositionPriceQuantity) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.PositionPriceQuantityKeyPrefixByOwnerAndPositionOrderHash(owner, positionOrderHash))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PositionPriceQuantity
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

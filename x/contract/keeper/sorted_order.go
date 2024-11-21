package keeper

import (
	"context"

	"gluon/x/contract/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetSortedOrder set a specific sortedOrder in the store from its index
func (k Keeper) SetSortedOrder(ctx context.Context, sortedOrder types.SortedOrder) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SortedOrderKeyPrefix))
	b := k.cdc.MustMarshal(&sortedOrder)
	store.Set(types.SortedOrderKey(
		sortedOrder.Expiry,
		sortedOrder.Index,
	), b)
}

// GetSortedOrder returns a sortedOrder from its index
func (k Keeper) GetSortedOrder(
	ctx context.Context,
	expiry string,
	index string,

) (val types.SortedOrder, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SortedOrderKeyPrefix))

	b := store.Get(types.SortedOrderKey(
		expiry,
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSortedOrder removes a sortedOrder from the store
func (k Keeper) RemoveSortedOrder(
	ctx context.Context,
	expiry string,
	index string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SortedOrderKeyPrefix))
	store.Delete(types.SortedOrderKey(
		expiry,
		index,
	))
}

// GetAllSortedOrder returns all sortedOrder
func (k Keeper) GetAllSortedOrder(ctx context.Context) (list []types.SortedOrder) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SortedOrderKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.SortedOrder
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

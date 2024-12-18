package keeper

import (
	"context"

	"gluon/x/order/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetOrder set a specific order in the store from its index
func (k Keeper) SetOrder(ctx context.Context, order types.Order) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.OrderKeyPrefix))
	b := k.cdc.MustMarshal(&order)
	store.Set(types.OrderKey(
		order.Owner,
		order.Hash,
	), b)
}

// GetOrder returns a order from its index
func (k Keeper) GetOrder(
	ctx context.Context,
	owner string,
	hash string,

) (val types.Order, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.OrderKeyPrefix))

	b := store.Get(types.OrderKey(
		owner,
		hash,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveOrder removes a order from the store
func (k Keeper) RemoveOrder(
	ctx context.Context,
	owner string,
	hash string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.OrderKeyPrefix))
	store.Delete(types.OrderKey(
		owner,
		hash,
	))
}

// GetAllOrder returns all order
func (k Keeper) GetAllOrder(ctx context.Context) (list []types.Order) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.OrderKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Order
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

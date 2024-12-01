package keeper

import (
	"context"

	"gluon/x/perp/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetPosition set a specific position in the store from its index
func (k Keeper) SetPosition(ctx context.Context, position types.Position) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PositionKeyPrefix))
	b := k.cdc.MustMarshal(&position)
	store.Set(types.PositionKey(
		position.Owner,
		position.OrderHash,
	), b)
}

// GetPosition returns a position from its index
func (k Keeper) GetPosition(
	ctx context.Context,
	owner string,
	orderHash string,
) (val types.Position, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PositionKeyPrefix))

	b := store.Get(types.PositionKey(
		owner,
		orderHash,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePosition removes a position from the store
func (k Keeper) RemovePosition(
	ctx context.Context,
	owner string,
	orderHash string,
) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PositionKeyPrefix))
	store.Delete(types.PositionKey(
		owner,
		orderHash,
	))
}

// GetAllPosition returns all position
func (k Keeper) GetAllPosition(ctx context.Context) (list []types.Position) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PositionKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Position
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

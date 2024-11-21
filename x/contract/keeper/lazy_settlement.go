package keeper

import (
	"context"
	"encoding/binary"

	"gluon/x/contract/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// GetLazySettlementCount get the total number of lazySettlement
func (k Keeper) GetLazySettlementCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.LazySettlementCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetLazySettlementCount set the total number of lazySettlement
func (k Keeper) SetLazySettlementCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.LazySettlementCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendLazySettlement appends a lazySettlement in the store with a new id and update the count
func (k Keeper) AppendLazySettlement(
	ctx context.Context,
	lazySettlement types.LazySettlement,
) uint64 {
	// Create the lazySettlement
	count := k.GetLazySettlementCount(ctx)

	// Set the ID of the appended value
	lazySettlement.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.LazySettlementKey))
	appendedValue := k.cdc.MustMarshal(&lazySettlement)
	store.Set(GetLazySettlementIDBytes(lazySettlement.Id), appendedValue)

	// Update lazySettlement count
	k.SetLazySettlementCount(ctx, count+1)

	return count
}

// SetLazySettlement set a specific lazySettlement in the store
func (k Keeper) SetLazySettlement(ctx context.Context, lazySettlement types.LazySettlement) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.LazySettlementKey))
	b := k.cdc.MustMarshal(&lazySettlement)
	store.Set(GetLazySettlementIDBytes(lazySettlement.Id), b)
}

// GetLazySettlement returns a lazySettlement from its id
func (k Keeper) GetLazySettlement(ctx context.Context, id uint64) (val types.LazySettlement, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.LazySettlementKey))
	b := store.Get(GetLazySettlementIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveLazySettlement removes a lazySettlement from the store
func (k Keeper) RemoveLazySettlement(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.LazySettlementKey))
	store.Delete(GetLazySettlementIDBytes(id))
}

// GetAllLazySettlement returns all lazySettlement
func (k Keeper) GetAllLazySettlement(ctx context.Context) (list []types.LazySettlement) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.LazySettlementKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.LazySettlement
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetLazySettlementIDBytes returns the byte representation of the ID
func GetLazySettlementIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.LazySettlementKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}

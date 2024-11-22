package keeper

import (
	"context"
	"encoding/binary"

	"gluon/x/contract/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// GetLazyContractCount get the total number of lazyContract
func (k Keeper) GetLazyContractCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.LazyContractCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetLazyContractCount set the total number of lazyContract
func (k Keeper) SetLazyContractCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.LazyContractCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendLazyContract appends a lazyContract in the store with a new id and update the count
func (k Keeper) AppendLazyContract(
	ctx context.Context,
	lazyContract types.LazyContract,
) uint64 {
	// Create the lazyContract
	count := k.GetLazyContractCount(ctx)

	// Set the ID of the appended value
	lazyContract.Id = count

	k.SetLazyContract(ctx, lazyContract)

	// Update lazyContract count
	k.SetLazyContractCount(ctx, count+1)

	return count
}

// SetLazyContract set a specific lazyContract in the store
func (k Keeper) SetLazyContract(ctx context.Context, lazyContract types.LazyContract) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.LazyContractKey))
	b := k.cdc.MustMarshal(&lazyContract)
	store.Set(GetLazyContractIDBytes(lazyContract.Id), b)
}

// GetLazyContract returns a lazyContract from its id
func (k Keeper) GetLazyContract(ctx context.Context, id uint64) (val types.LazyContract, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.LazyContractKey))
	b := store.Get(GetLazyContractIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveLazyContract removes a lazyContract from the store
func (k Keeper) RemoveLazyContract(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.LazyContractKey))
	store.Delete(GetLazyContractIDBytes(id))
}

// GetAllLazyContract returns all lazyContract
func (k Keeper) GetAllLazyContract(ctx context.Context) (list []types.LazyContract) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.LazyContractKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.LazyContract
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetLazyContractIDBytes returns the byte representation of the ID
func GetLazyContractIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.LazyContractKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}

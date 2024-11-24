package keeper

import (
	"context"

	"gluon/x/contract/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"

	"time"
)

// SetSortedLazyContract set a specific sortedLazyContract in the store from its index
func (k Keeper) SetSortedLazyContract(ctx context.Context, sortedLazyContract types.SortedLazyContract) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SortedLazyContractKeyPrefix))
	b := k.cdc.MustMarshal(&sortedLazyContract)
	store.Set(types.SortedLazyContractKey(
		sortedLazyContract.Expiry,
		sortedLazyContract.Id,
	), b)
}

// GetSortedLazyContract returns a sortedLazyContract from its index
func (k Keeper) GetSortedLazyContract(
	ctx context.Context,
	expiry time.Time,
	index uint64,

) (val types.SortedLazyContract, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SortedLazyContractKeyPrefix))

	b := store.Get(types.SortedLazyContractKey(
		uint64(expiry.UnixMilli()),
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSortedLazyContract removes a sortedLazyContract from the store
func (k Keeper) RemoveSortedLazyContract(
	ctx context.Context,
	expiry uint64,
	index uint64,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SortedLazyContractKeyPrefix))
	store.Delete(types.SortedLazyContractKey(
		expiry,
		index,
	))
}

// GetAllSortedLazyContract returns all sortedLazyContract
func (k Keeper) GetAllSortedLazyContract(ctx context.Context) (list []types.SortedLazyContract) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SortedLazyContractKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.SortedLazyContract
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

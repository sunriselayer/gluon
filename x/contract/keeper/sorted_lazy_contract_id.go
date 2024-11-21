package keeper

import (
	"context"

	"gluon/x/contract/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

func (k Keeper) SetSortedLazyContractId(ctx context.Context, expiry time.Time, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.SortedLazyContractIdPrefix(expiry))
	idBytes := sdk.Uint64ToBigEndian(id)
	store.Set(idBytes, idBytes)
}

func (k Keeper) RemoveSortedLazyContractId(ctx context.Context, expiry time.Time, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.SortedLazyContractIdPrefix(expiry))
	idBytes := sdk.Uint64ToBigEndian(id)
	store.Delete(idBytes)
}

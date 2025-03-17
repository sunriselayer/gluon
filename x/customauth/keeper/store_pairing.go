package keeper

import (
	"context"

	"gluon/x/customauth/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetPairing set a specific pairing in the store from its index
func (k Keeper) SetPairing(ctx context.Context, pairing types.Pairing) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PairingKeyPrefix))
	b := k.cdc.MustMarshal(&pairing)
	store.Set(types.PairingKey(
		pairing.Owner,
		pairing.Index,
	), b)
}

// GetPairing returns a pairing from its index
func (k Keeper) GetPairing(
	ctx context.Context,
	owner string,
	index string,

) (val types.Pairing, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PairingKeyPrefix))

	b := store.Get(types.PairingKey(
		owner,
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePairing removes a pairing from the store
func (k Keeper) RemovePairing(
	ctx context.Context,
	owner string,
	index string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PairingKeyPrefix))
	store.Delete(types.PairingKey(
		owner,
		index,
	))
}

// GetAllPairing returns all pairing
func (k Keeper) GetAllPairing(ctx context.Context) (list []types.Pairing) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PairingKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Pairing
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

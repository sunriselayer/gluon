package keeper

import (
	"context"
	"encoding/binary"

	"gluon/x/customauth/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
)

// GetPairingCount get the total number of pairing
func (k Keeper) GetPairingCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.PairingCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetPairingCount set the total number of pairing
func (k Keeper) SetPairingCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.PairingCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendPairing appends a pairing in the store with a new id and update the count
func (k Keeper) AppendPairing(
	ctx context.Context,
	pairing types.Pairing,
) (uint64, error) {
	//
	var pubKey cryptotypes.PubKey
	err := k.cdc.UnpackAny(&pairing.PublicKey, &pubKey)
	if err != nil {
		return 0, err
	}
	//

	// Create the pairing
	count := k.GetPairingCount(ctx)

	// Set the ID of the appended value
	pairing.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PairingKey))
	appendedValue := k.cdc.MustMarshal(&pairing)
	store.Set(GetPairingIDBytes(pairing.Address, pairing.Id), appendedValue)

	// Update pairing count
	k.SetPairingCount(ctx, count+1)

	return count, nil
}

// SetPairing set a specific pairing in the store
func (k Keeper) SetPairing(ctx context.Context, pairing types.Pairing) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PairingKey))
	b := k.cdc.MustMarshal(&pairing)
	store.Set(GetPairingIDBytes(pairing.Address, pairing.Id), b)
}

// GetPairing returns a pairing from its id
func (k Keeper) GetPairing(ctx context.Context, address string, id uint64) (val types.Pairing, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PairingKey))
	b := store.Get(GetPairingIDBytes(address, id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePairing removes a pairing from the store
func (k Keeper) RemovePairing(ctx context.Context, address string, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PairingKey))
	store.Delete(GetPairingIDBytes(address, id))
}

// GetAllPairing returns all pairing
func (k Keeper) GetAllPairing(ctx context.Context) (list []types.Pairing) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PairingKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Pairing
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetPairingIDBytes returns the byte representation of the ID
func GetPairingIDBytes(address string, id uint64) []byte {
	bz := types.PairingKeyPrefix(address)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}

func (k Keeper) GetPairingPubKey(ctx context.Context, pairing types.Pairing) (cryptotypes.PubKey, error) {
	var pubKey cryptotypes.PubKey
	err := k.cdc.UnpackAny(&pairing.PublicKey, &pubKey)
	if err != nil {
		return nil, err
	}

	return pubKey, nil
}

package keeper

import (
	"context"

	"gluon/x/perp/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetFundingRate set a specific fundingRate in the store from its index
func (k Keeper) SetFundingRate(ctx context.Context, fundingRate types.FundingRate) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FundingRateKeyPrefix))
	b := k.cdc.MustMarshal(&fundingRate)
	store.Set(types.FundingRateKey(
		fundingRate.DenomBase,
		fundingRate.DenomQuote,
		fundingRate.Seconds,
		fundingRate.Nanos,
	), b)
}

// GetFundingRate returns a fundingRate from its index
func (k Keeper) GetFundingRate(
	ctx context.Context,
	denomBase string,
	denomQuote string,
	seconds uint64,
	nanos uint32,

) (val types.FundingRate, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FundingRateKeyPrefix))

	b := store.Get(types.FundingRateKey(
		denomBase,
		denomQuote,
		seconds,
		nanos,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveFundingRate removes a fundingRate from the store
func (k Keeper) RemoveFundingRate(
	ctx context.Context,
	denomBase string,
	denomQuote string,
	seconds uint64,
	nanos uint32,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FundingRateKeyPrefix))
	store.Delete(types.FundingRateKey(
		denomBase,
		denomQuote,
		seconds,
		nanos,
	))
}

// GetAllFundingRate returns all fundingRate
func (k Keeper) GetAllFundingRate(ctx context.Context) (list []types.FundingRate) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FundingRateKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.FundingRate
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

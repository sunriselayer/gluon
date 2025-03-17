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

// Original

func (k Keeper) GetFundingRatesByPositionReverseIterator(ctx context.Context, denomBase string, denomQuote string) storetypes.Iterator {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.FundingRateKeyPrefixByDenoms(denomBase, denomQuote))
	iterator := storetypes.KVStoreReversePrefixIterator(store, []byte{})

	return iterator
}

func (k Keeper) GetFundingRatesByPositionReverse(ctx context.Context, position types.Position) []types.FundingRate {
	var fundingRates []types.FundingRate
	iterator := k.GetFundingRatesByPositionReverseIterator(ctx, position.DenomBase, position.DenomQuote)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var fundingRate types.FundingRate
		k.cdc.MustUnmarshal(iterator.Value(), &fundingRate)

		// At least one element will be appended because if-statement is after this
		fundingRates = append(fundingRates, fundingRate)

		fundingRateTime := fundingRate.GetTime()
		if fundingRateTime.Before(position.FundingFeeResetAt) || fundingRateTime.Equal(position.FundingFeeResetAt) {
			break
		}
	}

	return fundingRates
}

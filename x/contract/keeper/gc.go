package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"

	"gluon/x/contract/types"
)

func (k Keeper) GarbageCollect(goCtx context.Context) error {
	ctx := sdk.UnwrapSDKContext(goCtx)

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SortedOrderKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	// Because expiry is contained in the key, expiry will be iterated in ascending order
	for ; iterator.Valid(); iterator.Next() {
		var val types.SortedOrder
		k.cdc.MustUnmarshal(iterator.Value(), &val)

		if uint64(ctx.BlockTime().UnixMilli()) < val.Expiry {
			break
		}

		store.Delete(iterator.Key())
		k.RemoveOrder(ctx, val.Id)
	}

	return nil
}

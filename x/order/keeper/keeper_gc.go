package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"

	"gluon/x/order/types"

	gogo "github.com/gogo/protobuf/types"
)

func (k Keeper) GarbageCollect(goCtx context.Context) error {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.GarbageCollectOrder(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) GarbageCollectOrder(ctx sdk.Context) error {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SortedOrderKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	// Because expiry is contained in the key, expiry will be iterated in ascending order
	for ; iterator.Valid(); iterator.Next() {
		var val types.SortedOrder
		k.cdc.MustUnmarshal(iterator.Value(), &val)

		timestamp, err := gogo.TimestampProto(ctx.BlockTime())
		if err != nil {
			return err
		}

		if uint64(timestamp.Seconds) <= val.Seconds && uint32(timestamp.Nanos) < val.Nanos {
			break
		}

		store.Delete(iterator.Key())
		k.RemoveOrder(ctx, val.OrderOwner, val.OrderHash)
	}

	return nil
}

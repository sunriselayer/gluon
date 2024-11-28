package keeper

import (
	"context"

	"gluon/x/contract/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"time"
)

func (k Keeper) SortedOrders(ctx context.Context, req *types.QuerySortedOrdersRequest) (*types.QuerySortedOrdersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var sortedOrders []types.SortedOrder

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	sortedOrderStore := prefix.NewStore(store, types.KeyPrefix(types.SortedOrderKeyPrefix))

	pageRes, err := query.Paginate(sortedOrderStore, req.Pagination, func(key []byte, value []byte) error {
		var sortedOrder types.SortedOrder
		if err := k.cdc.Unmarshal(value, &sortedOrder); err != nil {
			return err
		}

		sortedOrders = append(sortedOrders, sortedOrder)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QuerySortedOrdersResponse{SortedOrders: sortedOrders, Pagination: pageRes}, nil
}

func (k Keeper) SortedOrder(ctx context.Context, req *types.QueryGetSortedOrderRequest) (*types.QueryGetSortedOrderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetSortedOrder(
		ctx,
		time.UnixMilli(int64(req.Expiry)),
		req.Id,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetSortedOrderResponse{SortedOrders: val}, nil
}

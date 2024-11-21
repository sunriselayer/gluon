package keeper

import (
	"context"

	"gluon/x/contract/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) OrderAll(ctx context.Context, req *types.QueryAllOrderRequest) (*types.QueryAllOrderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var orders []types.Order

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	orderStore := prefix.NewStore(store, types.KeyPrefix(types.OrderKeyPrefix))

	pageRes, err := query.Paginate(orderStore, req.Pagination, func(key []byte, value []byte) error {
		var order types.Order
		if err := k.cdc.Unmarshal(value, &order); err != nil {
			return err
		}

		orders = append(orders, order)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllOrderResponse{Order: orders, Pagination: pageRes}, nil
}

func (k Keeper) Order(ctx context.Context, req *types.QueryGetOrderRequest) (*types.QueryGetOrderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetOrder(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetOrderResponse{Order: val}, nil
}

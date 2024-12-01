package keeper

import (
	"context"

	"gluon/x/perp/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PositionPriceQuantities(ctx context.Context, req *types.QueryPositionPriceQuantitiesRequest) (*types.QueryPositionPriceQuantitiesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var positionPriceQuantities []types.PositionPriceQuantity

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	positionPriceQuantityStore := prefix.NewStore(store, types.KeyPrefix(types.PositionPriceQuantityKeyPrefix))

	pageRes, err := query.Paginate(positionPriceQuantityStore, req.Pagination, func(key []byte, value []byte) error {
		var positionPriceQuantity types.PositionPriceQuantity
		if err := k.cdc.Unmarshal(value, &positionPriceQuantity); err != nil {
			return err
		}

		positionPriceQuantities = append(positionPriceQuantities, positionPriceQuantity)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryPositionPriceQuantitiesResponse{PositionPriceQuantity: positionPriceQuantities, Pagination: pageRes}, nil
}

func (k Keeper) PositionPriceQuantity(ctx context.Context, req *types.QueryPositionPriceQuantityRequest) (*types.QueryPositionPriceQuantityResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetPositionPriceQuantity(
		ctx,
		req.Owner,
		req.PositionOrderHash,
		req.Price,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryPositionPriceQuantityResponse{PositionPriceQuantity: val}, nil
}

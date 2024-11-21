package keeper

import (
	"context"

	"gluon/x/contract/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) LazySettlementAll(ctx context.Context, req *types.QueryAllLazySettlementRequest) (*types.QueryAllLazySettlementResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var lazySettlements []types.LazySettlement

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	lazySettlementStore := prefix.NewStore(store, types.KeyPrefix(types.LazySettlementKey))

	pageRes, err := query.Paginate(lazySettlementStore, req.Pagination, func(key []byte, value []byte) error {
		var lazySettlement types.LazySettlement
		if err := k.cdc.Unmarshal(value, &lazySettlement); err != nil {
			return err
		}

		lazySettlements = append(lazySettlements, lazySettlement)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllLazySettlementResponse{LazySettlement: lazySettlements, Pagination: pageRes}, nil
}

func (k Keeper) LazySettlement(ctx context.Context, req *types.QueryGetLazySettlementRequest) (*types.QueryGetLazySettlementResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	lazySettlement, found := k.GetLazySettlement(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetLazySettlementResponse{LazySettlement: lazySettlement}, nil
}

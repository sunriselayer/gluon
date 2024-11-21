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

func (k Keeper) LazyContractAll(ctx context.Context, req *types.QueryAllLazyContractRequest) (*types.QueryAllLazyContractResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var lazyContracts []types.LazyContract

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	lazyContractStore := prefix.NewStore(store, types.KeyPrefix(types.LazyContractKey))

	pageRes, err := query.Paginate(lazyContractStore, req.Pagination, func(key []byte, value []byte) error {
		var lazyContract types.LazyContract
		if err := k.cdc.Unmarshal(value, &lazyContract); err != nil {
			return err
		}

		lazyContracts = append(lazyContracts, lazyContract)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllLazyContractResponse{LazyContract: lazyContracts, Pagination: pageRes}, nil
}

func (k Keeper) LazyContract(ctx context.Context, req *types.QueryGetLazyContractRequest) (*types.QueryGetLazyContractResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	lazyContract, found := k.GetLazyContract(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetLazyContractResponse{LazyContract: lazyContract}, nil
}

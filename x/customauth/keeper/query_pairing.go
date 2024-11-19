package keeper

import (
	"context"

	"gluon/x/customauth/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PairingAll(ctx context.Context, req *types.QueryAllPairingRequest) (*types.QueryAllPairingResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var pairings []types.Pairing

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	pairingStore := prefix.NewStore(store, types.PairingKeyPrefix(req.Address))

	pageRes, err := query.Paginate(pairingStore, req.Pagination, func(key []byte, value []byte) error {
		var pairing types.Pairing
		if err := k.cdc.Unmarshal(value, &pairing); err != nil {
			return err
		}

		pairings = append(pairings, pairing)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPairingResponse{Pairing: pairings, Pagination: pageRes}, nil
}

func (k Keeper) Pairing(ctx context.Context, req *types.QueryGetPairingRequest) (*types.QueryGetPairingResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	pairing, found := k.GetPairing(ctx, req.Address, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetPairingResponse{Pairing: pairing}, nil
}

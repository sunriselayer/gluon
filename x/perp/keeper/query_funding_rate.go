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

func (k Keeper) FundingRates(ctx context.Context, req *types.QueryFundingRatesRequest) (*types.QueryFundingRatesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var fundingRates []types.FundingRate

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	fundingRateStore := prefix.NewStore(store, types.FundingRateKeyPrefixByDenoms(req.DenomBase, req.DenomQuote))

	pageRes, err := query.Paginate(fundingRateStore, req.Pagination, func(key []byte, value []byte) error {
		var fundingRate types.FundingRate
		if err := k.cdc.Unmarshal(value, &fundingRate); err != nil {
			return err
		}

		fundingRates = append(fundingRates, fundingRate)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryFundingRatesResponse{FundingRates: fundingRates, Pagination: pageRes}, nil
}

func (k Keeper) FundingRate(ctx context.Context, req *types.QueryFundingRateRequest) (*types.QueryFundingRateResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetFundingRate(
		ctx,
		req.DenomBase,
		req.DenomQuote,
		req.Seconds,
		req.Nanos,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryFundingRateResponse{FundingRate: val}, nil
}

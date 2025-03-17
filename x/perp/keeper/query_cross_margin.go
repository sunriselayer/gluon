package keeper

import (
	"context"

	"gluon/x/perp/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CrossMargin(ctx context.Context, req *types.QueryCrossMarginRequest) (*types.QueryCrossMarginResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetCrossMargin(
		ctx,
		req.Owner,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryCrossMarginResponse{CrossMargin: val}, nil
}

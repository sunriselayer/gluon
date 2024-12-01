package keeper

import (
	"context"

	"gluon/x/spot/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) MatchOrder(goCtx context.Context, msg *types.MsgMatchOrder) (*types.MsgMatchOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgMatchOrderResponse{}, nil
}

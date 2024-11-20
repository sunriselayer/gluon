package keeper

import (
	"context"

	"gluon/x/contract/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) MatchOrder(goCtx context.Context, msg *types.MsgMatchOrder) (*types.MsgMatchOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.ValidateOrderPair(ctx, msg.Earlier, msg.Later)
	if err != nil {
		return nil, err
	}

	return &types.MsgMatchOrderResponse{}, nil
}

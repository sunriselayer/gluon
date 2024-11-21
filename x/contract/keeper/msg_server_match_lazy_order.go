package keeper

import (
	"context"

	"gluon/x/contract/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) MatchLazyOrder(goCtx context.Context, msg *types.MsgMatchLazyOrder) (*types.MsgMatchLazyOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_ = ctx

	return &types.MsgMatchLazyOrderResponse{}, nil
}

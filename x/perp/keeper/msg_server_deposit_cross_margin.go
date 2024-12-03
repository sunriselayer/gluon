package keeper

import (
	"context"

	"gluon/x/perp/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) DepositCrossMargin(goCtx context.Context, msg *types.MsgDepositCrossMargin) (*types.MsgDepositCrossMarginResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgDepositCrossMarginResponse{}, nil
}

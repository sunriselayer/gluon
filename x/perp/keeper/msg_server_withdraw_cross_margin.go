package keeper

import (
	"context"

	"gluon/x/perp/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) WithdrawCrossMargin(goCtx context.Context, msg *types.MsgWithdrawCrossMargin) (*types.MsgWithdrawCrossMarginResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgWithdrawCrossMarginResponse{}, nil
}

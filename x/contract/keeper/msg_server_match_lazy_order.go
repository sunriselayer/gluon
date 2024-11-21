package keeper

import (
	"context"

	"gluon/x/contract/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) MatchLazyOrder(goCtx context.Context, msg *types.MsgMatchLazyOrder) (*types.MsgMatchLazyOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	earlier, found := k.GetOrder(ctx, msg.EarlierOrderId)
	if !found {
		return nil, types.ErrOrderNotFound
	}

	later, found := k.GetOrder(ctx, msg.LaterOrderId)
	if !found {
		return nil, types.ErrOrderNotFound
	}

	err := earlier.CrossValidate(later, msg.Price)
	if err != nil {
		return nil, err
	}

	err = k.ValidateContractAmount(ctx, earlier, msg.Amount)
	if err != nil {
		return nil, err
	}

	err = k.ValidateContractAmount(ctx, later, msg.Amount)
	if err != nil {
		return nil, err
	}

	err = k.SettleLazy(ctx, earlier, later, msg.Amount)
	if err != nil {
		return nil, err
	}

	err = k.AddContractedAmount(ctx, uint64(earlier.Expiry.UnixMilli()), earlier.Id, msg.Amount)
	if err != nil {
		return nil, err
	}
	err = k.AddContractedAmount(ctx, uint64(later.Expiry.UnixMilli()), later.Id, msg.Amount)
	if err != nil {
		return nil, err
	}

	return &types.MsgMatchLazyOrderResponse{}, nil
}

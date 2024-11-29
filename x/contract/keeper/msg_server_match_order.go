package keeper

import (
	"context"

	"gluon/x/contract/types"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) MatchOrder(goCtx context.Context, msg *types.MsgMatchOrder) (*types.MsgMatchOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	buy, found := k.GetOrder(ctx, msg.OrderIdBuy)
	if !found {
		return nil, types.ErrOrderNotFound
	}

	sell, found := k.GetOrder(ctx, msg.OrderIdSell)
	if !found {
		return nil, types.ErrOrderNotFound
	}

	price, err := sdkmath.LegacyNewDecFromStr(msg.Price)
	if err != nil {
		return nil, err
	}

	err = buy.CrossValidate(sell, price, ctx.BlockTime())
	if err != nil {
		return nil, err
	}

	err = k.ValidateContractAmount(ctx, buy, msg.Amount)
	if err != nil {
		return nil, err
	}

	err = k.ValidateContractAmount(ctx, sell, msg.Amount)
	if err != nil {
		return nil, err
	}

	err = k.Contract(ctx, buy, sell, msg.Amount, price)
	if err != nil {
		return nil, err
	}

	return &types.MsgMatchOrderResponse{}, nil
}

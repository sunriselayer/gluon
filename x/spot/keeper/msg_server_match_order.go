package keeper

import (
	"context"

	ordertypes "gluon/x/order/types"
	"gluon/x/spot/types"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) MatchOrder(goCtx context.Context, msg *types.MsgMatchOrder) (*types.MsgMatchOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	buy, buyBody, err := k.orderKeeper.GetOrderAndBody(ctx, msg.OrderHashBuy)
	if err != nil {
		return nil, err
	}

	sell, sellBody, err := k.orderKeeper.GetOrderAndBody(ctx, msg.OrderHashSell)
	if err != nil {
		return nil, err
	}

	buySpot, ok := buyBody.(*types.SpotOrder)
	if !ok {
		return nil, errorsmod.Wrapf(types.ErrInvalidOrderType, "OrderHashBuy: %s", buy.Hash)
	}

	sellSpot, ok := sellBody.(*types.SpotOrder)
	if !ok {
		return nil, errorsmod.Wrapf(types.ErrInvalidOrderType, "OrderHashSell: %s", sell.Hash)
	}

	price, err := sdkmath.LegacyNewDecFromStr(msg.Price)
	if err != nil {
		return nil, err
	}

	err = ordertypes.CrossValidateBasic(buySpot.BaseOrder, sellSpot.BaseOrder, price, ctx.BlockTime())
	if err != nil {
		return nil, err
	}

	err = ordertypes.ValidateOrderContractAmount(buySpot.Amount, buy.ContractedAmount, msg.Amount)
	if err != nil {
		return nil, err
	}

	err = ordertypes.ValidateOrderContractAmount(sellSpot.Amount, sell.ContractedAmount, msg.Amount)
	if err != nil {
		return nil, err
	}

	err = k.Swap(ctx, *buySpot, *sellSpot, msg.Amount, price)
	if err != nil {
		return nil, err
	}

	err = k.orderKeeper.AddContractedAmount(ctx, buy.Hash, msg.Amount)
	if err != nil {
		return nil, err
	}
	err = k.orderKeeper.AddContractedAmount(ctx, sell.Hash, msg.Amount)
	if err != nil {
		return nil, err
	}

	return &types.MsgMatchOrderResponse{}, nil
}

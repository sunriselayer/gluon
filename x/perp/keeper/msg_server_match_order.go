package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	ordertypes "gluon/x/order/types"
	"gluon/x/perp/types"
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

	var buyPerp types.PerpOrder
	var sellPerp types.PerpOrder
	var ok bool

	buyPerp, ok = buyBody.(*types.PerpPositionCreateOrder)
	if !ok {
		return nil, errorsmod.Wrapf(types.ErrInvalidOrderType, "OrderHashBuy: %s", buy.Hash)
	}

	sellPerp, ok = sellBody.(*types.PerpPositionCreateOrder)
	if !ok {
		return nil, errorsmod.Wrapf(types.ErrInvalidOrderType, "OrderHashSell: %s", sell.Hash)
	}

	price, err := sdkmath.LegacyNewDecFromStr(msg.Price)
	if err != nil {
		return nil, err
	}

	err = types.PerpOrderCrossValidateBasic(buyPerp, sellPerp, price, ctx.BlockTime())
	if err != nil {
		return nil, err
	}

	err = ordertypes.ValidateOrderContractAmount(buyPerp.GetAmount(), buy.ContractedAmount, msg.Amount)
	if err != nil {
		return nil, err
	}

	err = ordertypes.ValidateOrderContractAmount(sellPerp.GetAmount(), sell.ContractedAmount, msg.Amount)
	if err != nil {
		return nil, err
	}

	err = k.CreatePosition(ctx, buyPerp, sellPerp, msg.Amount, price)
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

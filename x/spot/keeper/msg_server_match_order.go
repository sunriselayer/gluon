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

	buy, buyBody, err := k.orderKeeper.GetOrderAndBody(ctx, msg.Buyer, msg.BuyerOrderHash)
	if err != nil {
		return nil, err
	}

	sell, sellBody, err := k.orderKeeper.GetOrderAndBody(ctx, msg.Seller, msg.SellerOrderHash)
	if err != nil {
		return nil, err
	}

	buySpot, ok := buyBody.(*types.SpotOrder)
	if !ok {
		return nil, errorsmod.Wrapf(types.ErrInvalidOrderType, "BuyerOrderHash: %s", msg.BuyerOrderHash)
	}

	sellSpot, ok := sellBody.(*types.SpotOrder)
	if !ok {
		return nil, errorsmod.Wrapf(types.ErrInvalidOrderType, "SellerOrderHash: %s", msg.SellerOrderHash)
	}

	price, err := sdkmath.LegacyNewDecFromStr(msg.Price)
	if err != nil {
		return nil, err
	}

	err = types.SpotOrderCrossValidateBasic(*buySpot, *sellSpot, price, ctx.BlockTime())
	if err != nil {
		return nil, err
	}

	err = ordertypes.ValidateOrderContractAmount(buySpot.Amount, buy.ContractedAmount, msg.Quantity)
	if err != nil {
		return nil, err
	}

	err = ordertypes.ValidateOrderContractAmount(sellSpot.Amount, sell.ContractedAmount, msg.Quantity)
	if err != nil {
		return nil, err
	}

	err = k.Swap(ctx, *buySpot, *sellSpot, msg.Quantity, price)
	if err != nil {
		return nil, err
	}

	err = k.orderKeeper.AddContractedAmount(ctx, buy.Owner, buy.Hash, msg.Quantity)
	if err != nil {
		return nil, err
	}
	err = k.orderKeeper.AddContractedAmount(ctx, sell.Owner, sell.Hash, msg.Quantity)
	if err != nil {
		return nil, err
	}

	return &types.MsgMatchOrderResponse{}, nil
}

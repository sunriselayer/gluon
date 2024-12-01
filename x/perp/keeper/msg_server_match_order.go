package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	// sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	ordertypes "gluon/x/order/types"
	"gluon/x/perp/types"
)

func (k msgServer) MatchOrder(goCtx context.Context, msg *types.MsgMatchOrder) (*types.MsgMatchOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	price, err := sdkmath.LegacyNewDecFromStr(msg.Price)
	if err != nil {
		return nil, err
	}

	long, longBody, err := k.orderKeeper.GetOrderAndBody(ctx, msg.OrderHashBuy)
	if err != nil {
		return nil, err
	}
	short, shortBody, err := k.orderKeeper.GetOrderAndBody(ctx, msg.OrderHashSell)
	if err != nil {
		return nil, err
	}

	longPerp, longDenomBase, longDenomQuote, longDirection, err := k.GetOrderBodyDenomsDirection(ctx, longBody)
	if err != nil {
		return nil, errorsmod.Wrapf(err, "OrderHashBuy: %s", long.Hash)
	}
	shortPerp, shortDenomBase, shortDenomQuote, shortDirection, err := k.GetOrderBodyDenomsDirection(ctx, shortBody)
	if err != nil {
		return nil, errorsmod.Wrapf(err, "OrderHashSell: %s", short.Hash)
	}

	// Denom and Direction
	if longDenomBase != shortDenomBase || longDenomQuote != shortDenomQuote {
		return nil, ordertypes.ErrDenomMismatch
	}
	if longDirection != ordertypes.OrderDirection_ORDER_DIRECTION_BUY {
		return nil, ordertypes.ErrInvalidOrderDirection
	}
	if shortDirection != ordertypes.OrderDirection_ORDER_DIRECTION_SELL {
		return nil, ordertypes.ErrInvalidOrderDirection
	}

	// Cross validate
	err = types.PerpOrderCrossValidateBasic(longPerp, shortPerp, price, ctx.BlockTime())
	if err != nil {
		return nil, err
	}

	err = ordertypes.ValidateOrderContractAmount(longPerp.GetAmount(), long.ContractedAmount, msg.Amount)
	if err != nil {
		return nil, err
	}
	err = ordertypes.ValidateOrderContractAmount(shortPerp.GetAmount(), short.ContractedAmount, msg.Amount)
	if err != nil {
		return nil, err
	}

	// Create or Cancel
	err = k.CreateUpdateCancelPosition(ctx, msg.OrderHashBuy, longPerp, msg.Amount, price)
	if err != nil {
		return nil, err
	}
	err = k.CreateUpdateCancelPosition(ctx, msg.OrderHashSell, shortPerp, msg.Amount, price)
	if err != nil {
		return nil, err
	}

	err = k.orderKeeper.AddContractedAmount(ctx, long.Hash, msg.Amount)
	if err != nil {
		return nil, err
	}
	err = k.orderKeeper.AddContractedAmount(ctx, short.Hash, msg.Amount)
	if err != nil {
		return nil, err
	}

	return &types.MsgMatchOrderResponse{}, nil
}

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

	long, longBody, err := k.orderKeeper.GetOrderAndBody(ctx, msg.Buyer, msg.OrderHashBuyer)
	if err != nil {
		return nil, err
	}
	short, shortBody, err := k.orderKeeper.GetOrderAndBody(ctx, msg.Seller, msg.OrderHashSeller)
	if err != nil {
		return nil, err
	}

	longPerp, longDenomBase, longDenomQuote, longDirection, err := k.GetOrderBodyDenomsDirection(ctx, longBody)
	if err != nil {
		return nil, errorsmod.Wrapf(err, "OrderHashBuyer: %s", msg.OrderHashBuyer)
	}
	shortPerp, shortDenomBase, shortDenomQuote, shortDirection, err := k.GetOrderBodyDenomsDirection(ctx, shortBody)
	if err != nil {
		return nil, errorsmod.Wrapf(err, "OrderHashSeller: %s", msg.OrderHashSeller)
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

	err = ordertypes.ValidateOrderContractAmount(longPerp.GetBaseOrder().Amount, long.ContractedAmount, msg.Quantity)
	if err != nil {
		return nil, err
	}
	err = ordertypes.ValidateOrderContractAmount(shortPerp.GetBaseOrder().Amount, short.ContractedAmount, msg.Quantity)
	if err != nil {
		return nil, err
	}

	// Create or Cancel
	err = k.CreateUpdateCancelPosition(ctx, msg.OrderHashBuyer, longPerp, price, msg.Quantity)
	if err != nil {
		return nil, err
	}
	err = k.CreateUpdateCancelPosition(ctx, msg.OrderHashSeller, shortPerp, price, msg.Quantity)
	if err != nil {
		return nil, err
	}

	err = k.orderKeeper.AddContractedAmount(ctx, long.Owner, long.Hash, msg.Quantity)
	if err != nil {
		return nil, err
	}
	err = k.orderKeeper.AddContractedAmount(ctx, short.Owner, short.Hash, msg.Quantity)
	if err != nil {
		return nil, err
	}

	return &types.MsgMatchOrderResponse{}, nil
}

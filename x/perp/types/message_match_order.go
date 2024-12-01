package types

import (
	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	// sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	ordertypes "gluon/x/order/types"
)

var _ sdk.Msg = &MsgMatchOrder{}

func NewMsgMatchOrder(
	buyer string,
	seller string,
	orderHashBuy string,
	orderHashSell string,
	amount sdkmath.Int,
	price string,
) *MsgMatchOrder {
	return &MsgMatchOrder{
		Buyer:         buyer,
		Seller:        seller,
		OrderHashBuy:  orderHashBuy,
		OrderHashSell: orderHashSell,
		Amount:        amount,
		Price:         price,
	}
}

func (msg *MsgMatchOrder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Buyer)
	if err != nil {
		return err
	}

	_, err = sdk.AccAddressFromBech32(msg.Seller)
	if err != nil {
		return err
	}

	if len(msg.OrderHashBuy) == 0 {
		return ordertypes.ErrEmptyOrderHash
	}

	if len(msg.OrderHashSell) == 0 {
		return ordertypes.ErrEmptyOrderHash
	}

	if !msg.Amount.IsPositive() {
		return errorsmod.Wrapf(ordertypes.ErrNotPositive, "Amount: %s", msg.Amount.String())
	}

	price, err := sdkmath.LegacyNewDecFromStr(msg.Price)
	if err != nil {
		return err
	}

	if !price.IsPositive() {
		return errorsmod.Wrapf(ordertypes.ErrNotPositive, "Price: %s", price.String())
	}

	return nil
}

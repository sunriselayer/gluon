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
	orderHashBuyer string,
	orderHashSeller string,
	quantity sdkmath.Int,
	price string,
) *MsgMatchOrder {
	return &MsgMatchOrder{
		Buyer:           buyer,
		Seller:          seller,
		OrderHashBuyer:  orderHashBuyer,
		OrderHashSeller: orderHashSeller,
		Quantity:        quantity,
		Price:           price,
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

	if len(msg.OrderHashBuyer) == 0 {
		return ordertypes.ErrEmptyOrderHash
	}

	if len(msg.OrderHashSeller) == 0 {
		return ordertypes.ErrEmptyOrderHash
	}

	price, err := sdkmath.LegacyNewDecFromStr(msg.Price)
	if err != nil {
		return err
	}

	if !price.IsPositive() {
		return errorsmod.Wrapf(ordertypes.ErrNotPositive, "Price: %s", price.String())
	}

	if !msg.Quantity.IsPositive() {
		return errorsmod.Wrapf(ordertypes.ErrNotPositive, "Quantity: %s", msg.Quantity.String())
	}

	return nil
}

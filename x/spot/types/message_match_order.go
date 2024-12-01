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
	buyerOrderHash string,
	seller string,
	sellerOrderHash string,
	quantity sdkmath.Int,
	price string,
) *MsgMatchOrder {
	return &MsgMatchOrder{
		Buyer:           buyer,
		BuyerOrderHash:  buyerOrderHash,
		Seller:          seller,
		SellerOrderHash: sellerOrderHash,
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

	if len(msg.BuyerOrderHash) == 0 {
		return ordertypes.ErrEmptyOrderHash
	}

	if len(msg.SellerOrderHash) == 0 {
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

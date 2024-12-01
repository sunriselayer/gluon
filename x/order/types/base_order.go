package types

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	errorsmod "cosmossdk.io/errors"
)

var _ OrderI = &BaseOrder{}

func (order BaseOrder) ValidateBasic() error {
	var err error

	_, err = sdk.AccAddressFromBech32(order.AddressString)
	if err != nil {
		return err
	}
	err = sdk.ValidateDenom(order.DenomBase)
	if err != nil {
		return err
	}
	err = sdk.ValidateDenom(order.DenomQuote)
	if err != nil {
		return err
	}

	if !order.Amount.IsPositive() {
		return errorsmod.Wrapf(ErrNotPositive, "amount: %s", order.Amount.String())
	}

	if len(order.LimitPriceString) > 0 {
		limitPrice, err := sdkmath.LegacyNewDecFromStr(order.LimitPriceString)
		if err != nil {
			return err
		}

		if !limitPrice.IsPositive() {
			return errorsmod.Wrapf(ErrNotPositive, "price: %s", limitPrice.String())
		}
	}

	return nil
}

func (order BaseOrder) GetAddress() sdk.AccAddress {
	val, err := sdk.AccAddressFromBech32(order.AddressString)
	if err != nil {
		return nil
	}
	return val
}

func (order BaseOrder) GetAmount() sdkmath.Int {
	return order.Amount
}

func (order BaseOrder) GetLimitPrice() *sdkmath.LegacyDec {
	if len(order.LimitPriceString) == 0 {
		return nil
	}
	val, err := sdkmath.LegacyNewDecFromStr(order.LimitPriceString)
	if err != nil {
		return nil
	}
	return &val
}

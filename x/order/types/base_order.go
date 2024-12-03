package types

import (
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	errorsmod "cosmossdk.io/errors"
)

func BaseOrderCrossValidateBasic(buy BaseOrder, sell BaseOrder, price sdkmath.LegacyDec, blockTime time.Time) error {
	if buy.GetAddress().Equals(sell.GetAddress()) {
		return ErrSameAddress
	}

	if buy.DenomBase != sell.DenomBase || buy.DenomQuote != sell.DenomQuote {
		return ErrDenomMismatch
	}

	if buy.Direction != OrderDirection_ORDER_DIRECTION_BUY {
		return ErrInvalidOrderDirection
	}
	if sell.Direction != OrderDirection_ORDER_DIRECTION_SELL {
		return ErrInvalidOrderDirection
	}

	buyLimitPrice := buy.GetLimitPrice()
	sellLimitPrice := sell.GetLimitPrice()

	if buyLimitPrice == nil && sellLimitPrice == nil {
		return ErrBothMarketPriceOrder
	}

	if buyLimitPrice != nil && price.GT(*buyLimitPrice) {
		return errorsmod.Wrapf(ErrPriceMismatch, "price: %s, buy limit price: %s", price.String(), buyLimitPrice.String())
	}

	if sellLimitPrice != nil && price.LT(*sellLimitPrice) {
		return errorsmod.Wrapf(ErrPriceMismatch, "price: %s, sell limit price: %s", price.String(), sellLimitPrice.String())
	}

	if blockTime.After(buy.GetExpiry()) {
		return ErrOrderExpired
	}

	if blockTime.After(sell.GetExpiry()) {
		return ErrOrderExpired
	}

	return nil
}

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

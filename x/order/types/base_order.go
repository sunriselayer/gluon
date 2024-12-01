package types

import (
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	errorsmod "cosmossdk.io/errors"
)

func NewBaseOrder(
	address string,
	nonce uint64,
	denomBase string,
	denomQuote string,
	direction OrderDirection,
	amount sdkmath.Int,
	limitPrice *sdkmath.LegacyDec,
) BaseOrder {
	return BaseOrder{
		Address:    address,
		Nonce:      nonce,
		DenomBase:  denomBase,
		DenomQuote: denomQuote,
		Direction:  direction,
		Amount:     amount,
		LimitPrice: limitPrice.String(),
	}
}

func (order BaseOrder) Validate() error {
	var err error

	_, err = sdk.AccAddressFromBech32(order.Address)
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

	if len(order.LimitPrice) > 0 {
		limitPrice, err := sdkmath.LegacyNewDecFromStr(order.LimitPrice)
		if err != nil {
			return err
		}

		if !limitPrice.IsPositive() {
			return errorsmod.Wrapf(ErrNotPositive, "price: %s", limitPrice.String())
		}
	}

	return nil
}

func (buy BaseOrder) CrossValidate(sell BaseOrder, price sdkmath.LegacyDec, blockTime time.Time) error {
	if buy.Address == sell.Address {
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

	var buyLimitPrice *sdkmath.LegacyDec
	var sellLimitPrice *sdkmath.LegacyDec

	if len(buy.LimitPrice) > 0 {
		limitPrice, err := sdkmath.LegacyNewDecFromStr(buy.LimitPrice)
		if err != nil {
			return err
		}
		buyLimitPrice = &limitPrice
	}

	if len(sell.LimitPrice) > 0 {
		limitPrice, err := sdkmath.LegacyNewDecFromStr(sell.LimitPrice)
		if err != nil {
			return err
		}
		sellLimitPrice = &limitPrice
	}

	if buyLimitPrice == nil && sellLimitPrice == nil {
		return ErrBothMarketPriceOrder
	}

	if buyLimitPrice != nil && price.GT(*buyLimitPrice) {
		return errorsmod.Wrapf(ErrPriceMismatch, "price: %s, buy limit price: %s", price.String(), buyLimitPrice.String())
	}

	if sellLimitPrice != nil && price.LT(*sellLimitPrice) {
		return errorsmod.Wrapf(ErrPriceMismatch, "price: %s, sell limit price: %s", price.String(), sellLimitPrice.String())
	}

	if blockTime.After(buy.Expiry) {
		return ErrOrderExpired
	}

	if blockTime.After(sell.Expiry) {
		return ErrOrderExpired
	}

	return nil
}

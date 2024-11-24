package types

import (
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	errorsmod "cosmossdk.io/errors"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
)

func NewOrder(
	address string,
	denomBase string,
	denomQuote string,
	direction OrderDirection,
	amount sdkmath.Int,
	limitPrice *sdkmath.LegacyDec,
	stopPrice *sdkmath.LegacyDec,
	lazyContract bool,
	allowLazyCounterparty bool,
) Order {
	return Order{
		Address:               address,
		DenomBase:             denomBase,
		DenomQuote:            denomQuote,
		Direction:             direction,
		Amount:                amount,
		LimitPrice:            limitPrice,
		StopPrice:             stopPrice,
		LazyContract:          lazyContract,
		AllowLazyCounterparty: allowLazyCounterparty,
	}
}

func (order Order) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(order.Address)
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
		return errorsmod.Wrap(ErrNotPositiveAmount, "amount must be positive")
	}

	return nil
}

func (order Order) VerifySignature(pubKey cryptotypes.PubKey, signature []byte) error {
	bodyBytes, err := order.Marshal()
	if err != nil {
		return err
	}

	if !pubKey.VerifySignature(bodyBytes, signature) {
		return errorsmod.Wrapf(ErrInvalidSignature, "address: %s", order.Address)
	}
	return nil
}

func (buy Order) CrossValidate(sell Order, price sdkmath.LegacyDec, blockTime time.Time) error {
	if buy.Address == sell.Address {
		return ErrSameAddress
	}

	if buy.DenomBase != sell.DenomBase || buy.DenomQuote != sell.DenomQuote {
		return ErrDenomMismatch
	}

	if buy.Direction != OrderDirection_BUY {
		return ErrInvalidOrderDirection
	}
	if sell.Direction != OrderDirection_SELL {
		return ErrInvalidOrderDirection
	}

	if buy.LimitPrice == nil && sell.LimitPrice == nil {
		return ErrBothMarketPriceOrder
	}

	if buy.LimitPrice != nil && price.GT(*buy.LimitPrice) {
		return errorsmod.Wrapf(ErrPriceMismatch, "price: %s, buy limit price: %s", price.String(), buy.LimitPrice.String())
	}

	if sell.LimitPrice != nil && price.LT(*sell.LimitPrice) {
		return errorsmod.Wrapf(ErrPriceMismatch, "price: %s, sell limit price: %s", price.String(), sell.LimitPrice.String())
	}

	if buy.StopPrice != nil && price.LT(*buy.StopPrice) {
		return errorsmod.Wrapf(ErrPriceMismatch, "price: %s, buy stop price: %s", price.String(), buy.StopPrice.String())
	}

	if sell.StopPrice != nil && price.GT(*sell.StopPrice) {
		return errorsmod.Wrapf(ErrPriceMismatch, "price: %s, sell stop price: %s", price.String(), sell.StopPrice.String())
	}

	if blockTime.After(buy.Expiry) {
		return ErrOrderExpired
	}

	if blockTime.After(sell.Expiry) {
		return ErrOrderExpired
	}

	if buy.LazyContract && !sell.AllowLazyCounterparty {
		return ErrLazyContractNotAllowed
	}

	if sell.LazyContract && !buy.AllowLazyCounterparty {
		return ErrLazyContractNotAllowed
	}

	return nil
}

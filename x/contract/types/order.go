package types

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	errorsmod "cosmossdk.io/errors"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
)

func NewOrder(address string, baseDenom string, quoteDenom string, direction OrderDirection, type_ OrderType, amount sdkmath.Int, limitPrice *sdkmath.LegacyDec, stopPrice *sdkmath.LegacyDec) Order {
	return Order{
		Address:    address,
		BaseDenom:  baseDenom,
		QuoteDenom: quoteDenom,
		Direction:  direction,
		Type:       type_,
		Amount:     amount,
		LimitPrice: limitPrice,
		StopPrice:  stopPrice,
	}
}

func (order Order) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(order.Address)
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

func (order Order) CrossValidate(other Order, price sdkmath.LegacyDec) error {
	if order.BaseDenom != other.BaseDenom || order.QuoteDenom != other.QuoteDenom {
		return ErrDenomMismatch
	}

	if order.Direction == OrderDirection_UNKNOWN || other.Direction == OrderDirection_UNKNOWN {
		return ErrUnknownOrderDirection
	}
	if order.Direction == other.Direction {
		return ErrSameOrderDirection
	}

	if order.LimitPrice == nil && other.LimitPrice == nil {
		return ErrBothMarketPriceOrder
	}

	// TODO
	_ = price

	return nil
}

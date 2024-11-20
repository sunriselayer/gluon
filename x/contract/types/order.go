package types

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	errorsmod "cosmossdk.io/errors"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
)

func NewOrderBody(address string, baseDenom string, quoteDenom string, direction OrderDirection, type_ OrderType, amount sdkmath.Int, limitPrice *sdkmath.LegacyDec, stopPrice *sdkmath.LegacyDec) OrderBody {
	return OrderBody{
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

func (orderBody OrderBody) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(orderBody.Address)
	if err != nil {
		return err
	}
	if !orderBody.Amount.IsPositive() {
		return errorsmod.Wrap(ErrNotPositiveAmount, "amount must be positive")
	}

	return nil
}

func (order Order) ValidateBasic() error {
	err := order.Body.ValidateBasic()
	if err != nil {
		return err
	}

	if len(order.Signature) == 0 {
		return errorsmod.Wrap(ErrEmptySignature, "signature must not be empty")
	}

	return nil
}

func (order Order) Validate(pubKey cryptotypes.PubKey) error {
	bodyBytes, err := order.Body.Marshal()
	if err != nil {
		return err
	}

	if !pubKey.VerifySignature(bodyBytes, order.Signature) {
		return errorsmod.Wrapf(ErrInvalidSignature, "address: %s", order.Body.Address)
	}
	return nil
}

func (order Order) CrossValidate(other Order) error {
	if order.Body.BaseDenom != other.Body.BaseDenom || order.Body.QuoteDenom != other.Body.QuoteDenom {
		return ErrDenomMismatch
	}

	if order.Body.Direction == OrderDirection_UNKNOWN || other.Body.Direction == OrderDirection_UNKNOWN {
		return ErrUnknownOrderDirection
	}
	if order.Body.Direction == other.Body.Direction {
		return ErrSameOrderDirection
	}

	if order.Body.LimitPrice == nil && other.Body.LimitPrice == nil {
		return ErrBothMarketPriceOrder
	}

	return nil
}

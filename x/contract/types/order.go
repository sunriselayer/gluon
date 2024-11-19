package types

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	errorsmod "cosmossdk.io/errors"
)

func NewOrder(address string, direction OrderDirection, type_ OrderType, amount sdkmath.Int, limitPrice *sdkmath.LegacyDec, stopPrice *sdkmath.LegacyDec) Order {
	return Order{
		Address:    address,
		Direction:  direction,
		Type:       type_,
		Amount:     amount,
		LimitPrice: limitPrice,
		StopPrice:  stopPrice,
	}
}

func (order Order) Validate() error {
	_, err := sdk.AccAddressFromBech32(order.Address)
	if err != nil {
		return err
	}
	if !order.Amount.IsPositive() {
		return errorsmod.Wrap(ErrNotPositiveAmount, "amount must be positive")
	}

	return nil
}

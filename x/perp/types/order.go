package types

import (
	"time"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	ordertypes "gluon/x/order/types"
)

type PerpOrder interface {
	ordertypes.OrderBody
}

func PerpOrderCrossValidateBasic(buy PerpOrder, sell PerpOrder, price sdkmath.LegacyDec, blockTime time.Time) error {
	err := ordertypes.OrderInterfaceCrossValidateBasic(buy, sell, price, blockTime)
	if err != nil {
		return err
	}

	return nil
}

// PerpPositionCreateOrder

var _ PerpOrder = &PerpPositionCreateOrder{}

func (order PerpPositionCreateOrder) ValidateBasic() error {
	err := order.BaseOrder.ValidateBasic()
	if err != nil {
		return err
	}

	if order.MarginAmount.IsNegative() {
		return errorsmod.Wrap(ErrInvalidMargin, "margin amount must not be negative")
	}
	if order.IsolatedMargin && order.MarginAmount.IsZero() {
		return errorsmod.Wrap(ErrInvalidMargin, "isolated margin amount must not be zero")
	}
	return nil
}

func (order PerpPositionCreateOrder) PackAny() (codectypes.Any, error) {
	val, err := codectypes.NewAnyWithValue(&order)
	if err != nil {
		return codectypes.Any{}, err
	}
	return *val, nil
}

// PerpPositionCancelOrder

var _ PerpOrder = &PerpPositionCancelOrder{}

func (order PerpPositionCancelOrder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(order.AddressString)
	if err != nil {
		return err
	}

	if !order.Amount.IsPositive() {
		return errorsmod.Wrapf(ordertypes.ErrNotPositive, "amount: %s", order.Amount.String())
	}

	if len(order.LimitPriceString) > 0 {
		limitPrice, err := sdkmath.LegacyNewDecFromStr(order.LimitPriceString)
		if err != nil {
			return err
		}

		if !limitPrice.IsPositive() {
			return errorsmod.Wrapf(ordertypes.ErrNotPositive, "price: %s", limitPrice.String())
		}
	}

	return nil
}

func (order PerpPositionCancelOrder) GetAddress() sdk.AccAddress {
	val, err := sdk.AccAddressFromBech32(order.AddressString)
	if err != nil {
		return nil
	}
	return val
}

func (order PerpPositionCancelOrder) GetAmount() sdkmath.Int {
	return order.Amount
}

func (order PerpPositionCancelOrder) GetLimitPrice() *sdkmath.LegacyDec {
	if len(order.LimitPriceString) == 0 {
		return nil
	}
	val, err := sdkmath.LegacyNewDecFromStr(order.LimitPriceString)
	if err != nil {
		return nil
	}
	return &val
}

func (order PerpPositionCancelOrder) PackAny() (codectypes.Any, error) {
	val, err := codectypes.NewAnyWithValue(&order)
	if err != nil {
		return codectypes.Any{}, err
	}
	return *val, nil
}

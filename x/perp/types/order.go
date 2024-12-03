package types

import (
	"time"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"

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
	err := order.BaseOrder.ValidateBasic()
	if err != nil {
		return err
	}
	if len(order.PositionOrderHash) == 0 {
		return errorsmod.Wrap(ordertypes.ErrEmptyOrderHash, "position order hash must not be empty")
	}

	return nil
}

func (order PerpPositionCancelOrder) PackAny() (codectypes.Any, error) {
	val, err := codectypes.NewAnyWithValue(&order)
	if err != nil {
		return codectypes.Any{}, err
	}
	return *val, nil
}

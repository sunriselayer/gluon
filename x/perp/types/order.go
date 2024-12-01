package types

import (
	"time"

	sdkmath "cosmossdk.io/math"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	ordertypes "gluon/x/order/types"
)

// PerpPositionCreateOrder

var _ ordertypes.OrderBody = &PerpPositionCreateOrder{}

func (order PerpPositionCreateOrder) Validate() error {
	err := order.BaseOrder.Validate()
	if err != nil {
		return err
	}
	return nil
}

func (order PerpPositionCreateOrder) GetAddress() sdk.AccAddress {
	val, err := sdk.AccAddressFromBech32(order.Address)
	if err != nil {
		return nil
	}
	return val
}

func (order PerpPositionCreateOrder) GetAmount() sdkmath.Int {
	return order.Amount
}

func (order PerpPositionCreateOrder) GetExpiry() time.Time {
	return order.Expiry
}

func (order PerpPositionCreateOrder) PackAny() (codectypes.Any, error) {
	val, err := codectypes.NewAnyWithValue(&order)
	if err != nil {
		return codectypes.Any{}, err
	}
	return *val, nil
}

// PerpPositionCancelOrder

var _ ordertypes.OrderBody = &PerpPositionCancelOrder{}

func (order PerpPositionCancelOrder) Validate() error {
	err := order.BaseOrder.Validate()
	if err != nil {
		return err
	}
	return nil
}

func (order PerpPositionCancelOrder) GetAddress() sdk.AccAddress {
	val, err := sdk.AccAddressFromBech32(order.Address)
	if err != nil {
		return nil
	}
	return val
}

func (order PerpPositionCancelOrder) GetAmount() sdkmath.Int {
	return order.Amount
}

func (order PerpPositionCancelOrder) GetExpiry() time.Time {
	return order.Expiry
}

func (order PerpPositionCancelOrder) PackAny() (codectypes.Any, error) {
	val, err := codectypes.NewAnyWithValue(&order)
	if err != nil {
		return codectypes.Any{}, err
	}
	return *val, nil
}

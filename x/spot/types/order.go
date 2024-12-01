package types

import (
	"time"

	sdkmath "cosmossdk.io/math"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	ordertypes "gluon/x/order/types"
)

var _ ordertypes.OrderBody = &SpotOrder{}

func (order SpotOrder) ValidateBasic() error {
	err := order.BaseOrder.ValidateBasic()
	if err != nil {
		return err
	}
	return nil
}

func (order SpotOrder) GetAddress() sdk.AccAddress {
	val, err := sdk.AccAddressFromBech32(order.Address)
	if err != nil {
		return nil
	}
	return val
}

func (order SpotOrder) GetAmount() sdkmath.Int {
	return order.Amount
}

func (order SpotOrder) GetExpiry() time.Time {
	return order.Expiry
}

func (order SpotOrder) PackAny() (codectypes.Any, error) {
	val, err := codectypes.NewAnyWithValue(&order)
	if err != nil {
		return codectypes.Any{}, err
	}
	return *val, nil
}

package types

import (
	"time"

	sdkmath "cosmossdk.io/math"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"

	ordertypes "gluon/x/order/types"
)

func SpotOrderCrossValidateBasic(buy SpotOrder, sell SpotOrder, price sdkmath.LegacyDec, blockTime time.Time) error {
	err := ordertypes.BaseOrderCrossValidateBasic(buy.BaseOrder, sell.BaseOrder, price, blockTime)
	if err != nil {
		return err
	}

	return nil
}

var _ ordertypes.OrderBody = &SpotOrder{}

func (order SpotOrder) PackAny() (codectypes.Any, error) {
	val, err := codectypes.NewAnyWithValue(&order)
	if err != nil {
		return codectypes.Any{}, err
	}
	return *val, nil
}

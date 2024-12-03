package types

import (
	"time"

	sdkmath "cosmossdk.io/math"

	"github.com/cosmos/gogoproto/proto"
)

type OrderI interface {
	proto.Message
	GetBaseOrder() BaseOrder
	ValidateBasic() error
}

func OrderInterfaceCrossValidateBasic(buy OrderI, sell OrderI, price sdkmath.LegacyDec, blockTime time.Time) error {
	return BaseOrderCrossValidateBasic(buy.GetBaseOrder(), sell.GetBaseOrder(), price, blockTime)
}

func ValidateOrderContractAmount(orderAmount sdkmath.Int, orderContractedAmount sdkmath.Int, additionalContractAmount sdkmath.Int) error {
	if orderContractedAmount.Add(additionalContractAmount).GT(orderAmount) {
		return ErrContractAmountExceed
	}

	return nil
}

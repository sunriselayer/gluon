package types

import (
	sdkmath "cosmossdk.io/math"
)

func ValidateOrderContractAmount(orderAmount sdkmath.Int, orderContractedAmount sdkmath.Int, additionalContractAmount sdkmath.Int) error {
	if orderContractedAmount.Add(additionalContractAmount).GT(orderAmount) {
		return ErrContractAmountExceed
	}

	return nil
}

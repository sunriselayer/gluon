package types

import (
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/gogoproto/proto"
)

type OrderI interface {
	proto.Message
	ValidateBasic() error
	GetAddress() sdk.AccAddress
	// Amount of DenomBase
	GetAmount() sdkmath.Int
	GetLimitPrice() *sdkmath.LegacyDec
	GetExpiry() time.Time
}

func OrderInterfaceCrossValidateBasic(buy BaseOrder, sell BaseOrder, price sdkmath.LegacyDec, blockTime time.Time) error {
	if buy.DenomBase != sell.DenomBase || buy.DenomQuote != sell.DenomQuote {
		return ErrDenomMismatch
	}

	if buy.Direction != OrderDirection_ORDER_DIRECTION_BUY {
		return ErrInvalidOrderDirection
	}
	if sell.Direction != OrderDirection_ORDER_DIRECTION_SELL {
		return ErrInvalidOrderDirection
	}

	return nil
}

func ValidateOrderContractAmount(orderAmount sdkmath.Int, orderContractedAmount sdkmath.Int, additionalContractAmount sdkmath.Int) error {
	if orderContractedAmount.Add(additionalContractAmount).GT(orderAmount) {
		return ErrContractAmountExceed
	}

	return nil
}

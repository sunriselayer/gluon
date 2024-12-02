package types

import (
	"time"

	errorsmod "cosmossdk.io/errors"
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

func OrderInterfaceCrossValidateBasic(buy OrderI, sell OrderI, price sdkmath.LegacyDec, blockTime time.Time) error {
	if buy.GetAddress().Equals(sell.GetAddress()) {
		return ErrSameAddress
	}

	buyLimitPrice := buy.GetLimitPrice()
	sellLimitPrice := sell.GetLimitPrice()

	if buyLimitPrice == nil && sellLimitPrice == nil {
		return ErrBothMarketPriceOrder
	}

	if buyLimitPrice != nil && price.GT(*buyLimitPrice) {
		return errorsmod.Wrapf(ErrPriceMismatch, "price: %s, buy limit price: %s", price.String(), buyLimitPrice.String())
	}

	if sellLimitPrice != nil && price.LT(*sellLimitPrice) {
		return errorsmod.Wrapf(ErrPriceMismatch, "price: %s, sell limit price: %s", price.String(), sellLimitPrice.String())
	}

	if blockTime.After(buy.GetExpiry()) {
		return ErrOrderExpired
	}

	if blockTime.After(sell.GetExpiry()) {
		return ErrOrderExpired
	}

	return nil
}

func ValidateOrderContractAmount(orderAmount sdkmath.Int, orderContractedAmount sdkmath.Int, additionalContractAmount sdkmath.Int) error {
	if orderContractedAmount.Add(additionalContractAmount).GT(orderAmount) {
		return ErrContractAmountExceed
	}

	return nil
}

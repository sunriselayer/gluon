package types

import (
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	gogo "github.com/gogo/protobuf/types"
)

func (fundingRate FundingRate) GetTime() time.Time {
	timestamp := gogo.Timestamp{Seconds: int64(fundingRate.Seconds), Nanos: int32(fundingRate.Nanos)}
	t, err := gogo.TimestampFromProto(&timestamp)
	if err != nil {
		panic(err)
	}

	return t
}

func GetFundingFee(position Position, fundingRatesReverse []FundingRate) sdk.Coin {
	if len(fundingRatesReverse) == 0 {
		return sdk.NewCoin(position.DenomQuote, sdkmath.ZeroInt())
	}
	last := fundingRatesReverse[len(fundingRatesReverse)-1]
	fundingRatesReverse = fundingRatesReverse[:len(fundingRatesReverse)-1]
	feeSum := sdkmath.ZeroInt()

	for i := len(fundingRatesReverse) - 1; i >= 0; i-- {
		var duration time.Duration
		if i == 0 {
			duration = last.GetTime().Sub(position.FundingFeeResetAt)
		} else {
			duration = last.GetTime().Sub(fundingRatesReverse[i].GetTime())
		}
		numerator := sdkmath.LegacyMustNewDecFromStr(duration.String())
		denominator := sdkmath.LegacyMustNewDecFromStr(fundingRatesReverse[i].Interval.String())
		intervalRate := sdkmath.LegacyMustNewDecFromStr(fundingRatesReverse[i].Rate)
		price := sdkmath.LegacyMustNewDecFromStr(fundingRatesReverse[i].Price)

		rate := numerator.Quo(denominator).Mul(intervalRate)
		positionValueQuote := price.MulInt(position.Amount)
		fee := positionValueQuote.Mul(rate).TruncateInt()
		feeSum = feeSum.Add(fee)
	}
	if position.Direction == PositionDirection_POSITION_DIRECTION_SHORT {
		feeSum = feeSum.Neg()
	}

	return sdk.NewCoin(position.DenomQuote, feeSum)
}

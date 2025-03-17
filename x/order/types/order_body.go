package types

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/gogoproto/proto"
)

type OrderBody interface {
	proto.Message
	GetBaseOrder() BaseOrder
	ValidateBasic() error
	PackAny() (codectypes.Any, error)
}

func OrderInterfaceCrossValidateBasic(buy OrderBody, sell OrderBody, price sdkmath.LegacyDec, blockTime time.Time) error {
	return BaseOrderCrossValidateBasic(buy.GetBaseOrder(), sell.GetBaseOrder(), price, blockTime)
}

func ValidateOrderContractAmount(orderAmount sdkmath.Int, orderContractedAmount sdkmath.Int, additionalContractAmount sdkmath.Int) error {
	if orderContractedAmount.Add(additionalContractAmount).GT(orderAmount) {
		return ErrContractAmountExceed
	}

	return nil
}

func UnpackOrderAny(cdc codec.BinaryCodec, orderAny codectypes.Any) (OrderBody, error) {
	var orderBody OrderBody
	err := cdc.UnpackAny(&orderAny, &orderBody)
	if err != nil {
		return nil, err
	}
	return orderBody, nil
}

func GetOrderSignMessage(order OrderBody) ([]byte, error) {
	orderAny, err := order.PackAny()
	if err != nil {
		return nil, err
	}
	bz, err := orderAny.Marshal()
	if err != nil {
		return nil, err
	}
	return bz, nil
}

func VerifyOrderSignature(orderBody OrderBody, pubKey cryptotypes.PubKey, signature []byte) error {
	signMsg, err := GetOrderSignMessage(orderBody)
	if err != nil {
		return err
	}

	if !pubKey.VerifySignature(signMsg, signature) {
		return errorsmod.Wrapf(ErrInvalidSignature, "address: %s", orderBody.GetBaseOrder().AddressString)
	}
	return nil
}

func GetOrderHash(orderAny codectypes.Any) (string, error) {
	bz, err := orderAny.Marshal()
	if err != nil {
		return "", err
	}
	hasher := sha256.New()
	hasher.Write(bz)
	orderHash := hasher.Sum(nil)

	return hex.EncodeToString(orderHash), nil
}

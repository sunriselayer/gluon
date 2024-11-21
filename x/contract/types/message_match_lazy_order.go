package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &MsgMatchLazyOrder{}

func NewMsgMatchLazyOrder(
	addressBuy string,
	orderIdBuy string,
	addressSell string,
	orderIdSell string,
) *MsgMatchLazyOrder {
	return &MsgMatchLazyOrder{
		AddressBuy:  addressBuy,
		OrderIdBuy:  orderIdBuy,
		AddressSell: addressSell,
		OrderIdSell: orderIdSell,
	}
}

func (msg *MsgMatchLazyOrder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.AddressBuy)
	if err != nil {
		return err
	}

	if len(msg.OrderIdBuy) == 0 {
		return ErrEmptyOrderId
	}

	_, err = sdk.AccAddressFromBech32(msg.AddressSell)
	if err != nil {
		return err
	}

	if len(msg.OrderIdSell) == 0 {
		return ErrEmptyOrderId
	}

	return nil
}

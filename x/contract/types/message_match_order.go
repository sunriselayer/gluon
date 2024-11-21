package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &MsgMatchOrder{}

func NewMsgMatchOrder(
	addressBuy string,
	orderIdBuy string,
	addressSell string,
	orderIdSell string,
) *MsgMatchOrder {
	return &MsgMatchOrder{
		AddressBuy:  addressBuy,
		OrderIdBuy:  orderIdBuy,
		AddressSell: addressSell,
		OrderIdSell: orderIdSell,
	}
}

func (msg *MsgMatchOrder) ValidateBasic() error {
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

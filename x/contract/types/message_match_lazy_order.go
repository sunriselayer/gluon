package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &MsgMatchLazyOrder{}

func NewMsgMatchLazyOrder(
	earlierAddress string,
	earlierOrderId string,
	laterAddress string,
	laterOrderId string,
) *MsgMatchLazyOrder {
	return &MsgMatchLazyOrder{
		EarlierAddress: earlierAddress,
		EarlierOrderId: earlierOrderId,
		LaterAddress:   laterAddress,
		LaterOrderId:   laterOrderId,
	}
}

func (msg *MsgMatchLazyOrder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.EarlierAddress)
	if err != nil {
		return err
	}

	if len(msg.EarlierOrderId) == 0 {
		return ErrEmptyOrderId
	}

	_, err = sdk.AccAddressFromBech32(msg.LaterAddress)
	if err != nil {
		return err
	}

	if len(msg.LaterOrderId) == 0 {
		return ErrEmptyOrderId
	}

	return nil
}

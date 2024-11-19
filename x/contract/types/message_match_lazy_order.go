package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &MsgMatchLazyOrder{}

func NewMsgMatchLazyOrder(earlier Order, later Order) *MsgMatchLazyOrder {
	return &MsgMatchLazyOrder{
		Earlier: earlier,
		Later:   later,
	}
}

func (msg *MsgMatchLazyOrder) ValidateBasic() error {
	err := msg.Earlier.Validate()
	if err != nil {
		return err
	}
	err = msg.Later.Validate()
	if err != nil {
		return err
	}

	return nil
}

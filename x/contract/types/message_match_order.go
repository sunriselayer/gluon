package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &MsgMatchOrder{}

func NewMsgMatchOrder(earlier Order, later Order) *MsgMatchOrder {
	return &MsgMatchOrder{
		Earlier: earlier,
		Later:   later,
	}
}

func (msg *MsgMatchOrder) ValidateBasic() error {
	err := msg.Earlier.ValidateBasic()
	if err != nil {
		return err
	}
	err = msg.Later.ValidateBasic()
	if err != nil {
		return err
	}

	return nil
}

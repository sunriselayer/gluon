package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateOrder{}

func NewMsgCreateOrder(
	order Order,
) *MsgCreateOrder {
	return &MsgCreateOrder{
		Order: order,
	}
}

func (msg *MsgCreateOrder) ValidateBasic() error {
	err := msg.Order.ValidateBasic()
	if err != nil {
		return err
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteOrder{}

func NewMsgDeleteOrder(
	user string,
	id string,

) *MsgDeleteOrder {
	return &MsgDeleteOrder{
		User: user,
		Id:   id,
	}
}

func (msg *MsgDeleteOrder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.User)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid user address (%s)", err)
	}
	return nil
}

package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateOrder{}

func NewMsgCreateOrder(
	user string,
	index string,

) *MsgCreateOrder {
	return &MsgCreateOrder{
		User:  user,
		Index: index,
	}
}

func (msg *MsgCreateOrder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.User)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid user address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateOrder{}

func NewMsgUpdateOrder(
	user string,
	index string,

) *MsgUpdateOrder {
	return &MsgUpdateOrder{
		User:  user,
		Index: index,
	}
}

func (msg *MsgUpdateOrder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.User)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid user address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteOrder{}

func NewMsgDeleteOrder(
	user string,
	index string,

) *MsgDeleteOrder {
	return &MsgDeleteOrder{
		User:  user,
		Index: index,
	}
}

func (msg *MsgDeleteOrder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.User)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid user address (%s)", err)
	}
	return nil
}

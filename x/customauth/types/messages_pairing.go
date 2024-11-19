package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreatePairing{}

func NewMsgCreatePairing(user string, address string, publicKey string) *MsgCreatePairing {
	return &MsgCreatePairing{
		User:      user,
		Address:   address,
		PublicKey: publicKey,
	}
}

func (msg *MsgCreatePairing) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.User)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid user address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdatePairing{}

func NewMsgUpdatePairing(user string, id uint64, address string, publicKey string) *MsgUpdatePairing {
	return &MsgUpdatePairing{
		Id:        id,
		User:      user,
		Address:   address,
		PublicKey: publicKey,
	}
}

func (msg *MsgUpdatePairing) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.User)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid user address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeletePairing{}

func NewMsgDeletePairing(user string, id uint64) *MsgDeletePairing {
	return &MsgDeletePairing{
		Id:   id,
		User: user,
	}
}

func (msg *MsgDeletePairing) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.User)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid user address (%s)", err)
	}
	return nil
}

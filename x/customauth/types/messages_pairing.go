package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
)

var _ sdk.Msg = &MsgCreatePairing{}

func NewMsgCreatePairing(user string, publicKey codectypes.Any) *MsgCreatePairing {
	return &MsgCreatePairing{
		User:      user,
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

var _ sdk.Msg = &MsgDeletePairing{}

func NewMsgDeletePairing(user string, pairingIndex string) *MsgDeletePairing {
	return &MsgDeletePairing{
		User:         user,
		PairingIndex: pairingIndex,
	}
}

func (msg *MsgDeletePairing) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.User)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid user address (%s)", err)
	}
	return nil
}

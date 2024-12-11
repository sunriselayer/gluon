package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
)

var _ sdk.Msg = &MsgLazyRegisterOrder{}

func NewMsgLazyRegisterOrder(
	user string,
	order codectypes.Any,
	pairingIndex string,
	signature []byte,
) *MsgLazyRegisterOrder {
	return &MsgLazyRegisterOrder{
		User:         user,
		Order:        order,
		PairingIndex: pairingIndex,
		Signature:    signature,
	}
}

func (msg *MsgLazyRegisterOrder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.User)
	if err != nil {
		return err
	}
	if len(msg.Signature) == 0 {
		return errorsmod.Wrap(ErrEmptySignature, "signature must not be empty")
	}

	return nil
}

var _ sdk.Msg = &MsgCancelOrder{}

func NewMsgCancelOrder(
	user string,
	orderHash string,

) *MsgCancelOrder {
	return &MsgCancelOrder{
		User:      user,
		OrderHash: orderHash,
	}
}

func (msg *MsgCancelOrder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.User)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid user address (%s)", err)
	}
	return nil
}

package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgLazyRegisterOrder{}

func NewMsgLazyRegisterOrder(
	order Order,
	pairingId uint64,
	signature []byte,
) *MsgLazyRegisterOrder {
	return &MsgLazyRegisterOrder{
		Order:     order,
		PairingId: pairingId,
		Signature: signature,
	}
}

func (msg *MsgLazyRegisterOrder) ValidateBasic() error {
	err := msg.Order.ValidateBasic()
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
	orderId string,

) *MsgCancelOrder {
	return &MsgCancelOrder{
		User:    user,
		OrderId: orderId,
	}
}

func (msg *MsgCancelOrder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.User)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid user address (%s)", err)
	}
	return nil
}

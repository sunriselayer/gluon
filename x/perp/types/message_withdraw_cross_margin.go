package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgWithdrawCrossMargin{}

func NewMsgWithdrawCrossMargin(user string) *MsgWithdrawCrossMargin {
	return &MsgWithdrawCrossMargin{
		User: user,
	}
}

func (msg *MsgWithdrawCrossMargin) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.User)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid user address (%s)", err)
	}
	return nil
}

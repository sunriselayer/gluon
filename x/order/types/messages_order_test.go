package types

import (
	"testing"

	"gluon/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgLazyRegisterOrder_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgLazyRegisterOrder
		err  error
	}{
		{
			name: "invalid address",
			msg:  MsgLazyRegisterOrder{
				// Order: Order{
				// 	Address: "invalid address",
				// },
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg:  MsgLazyRegisterOrder{
				// Order: Order{
				// 	Address: sample.AccAddress(),
				// },
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgCancelOrder_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCancelOrder
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCancelOrder{
				User: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCancelOrder{
				User: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

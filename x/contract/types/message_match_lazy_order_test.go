package types

import (
	"testing"

	"gluon/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgMatchLazyOrder_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgMatchLazyOrder
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgMatchLazyOrder{
				AddressBuy:  "invalid_address",
				AddressSell: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgMatchLazyOrder{
				AddressBuy:  sample.AccAddress(),
				AddressSell: sample.AccAddress(),
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

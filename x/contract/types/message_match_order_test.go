package types

import (
	"testing"

	"gluon/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgMatchOrder_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgMatchOrder
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgMatchOrder{
				Earlier: Order{
					Body: OrderBody{
						Address: "invalid_address",
					},
				},
				Later: Order{
					Body: OrderBody{
						Address: "invalid_address",
					},
				},
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgMatchOrder{
				Earlier: Order{
					Body: OrderBody{
						Address: sample.AccAddress(),
					},
				},
				Later: Order{
					Body: OrderBody{
						Address: sample.AccAddress(),
					},
				},
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

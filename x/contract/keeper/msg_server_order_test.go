package keeper_test

import (
	"strconv"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "gluon/testutil/keeper"
	"gluon/x/contract/keeper"
	"gluon/x/contract/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestOrderMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.ContractKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	user := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateOrder{
			Order: types.Order{
				Id:      strconv.Itoa(i),
				Address: user,
			},
		}
		_, err := srv.CreateOrder(ctx, expected)
		require.NoError(t, err)
		rst, found := k.GetOrder(ctx,
			expected.Order.Id,
		)
		require.True(t, found)
		require.Equal(t, expected.Order.Address, rst.Address)
	}
}

func TestOrderMsgServerDelete(t *testing.T) {
	user := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteOrder
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteOrder{User: user,
				Id: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteOrder{User: "B",
				Id: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteOrder{User: user,
				Id: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.ContractKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			_, err := srv.CreateOrder(ctx, &types.MsgCreateOrder{
				Order: types.Order{
					Id:      strconv.Itoa(0),
					Address: user,
				},
			})
			require.NoError(t, err)
			_, err = srv.DeleteOrder(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetOrder(ctx,
					tc.request.Id,
				)
				require.False(t, found)
			}
		})
	}
}

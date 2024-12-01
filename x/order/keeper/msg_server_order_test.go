package keeper_test

import (
	"strconv"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "gluon/testutil/keeper"
	"gluon/x/order/keeper"
	"gluon/x/order/types"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestOrderMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.OrderKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	user := "A"
	for i := 0; i < 5; i++ {
		orderAny := codectypes.Any{}
		hash, err := types.GetOrderHash(orderAny)
		require.NoError(t, err)
		expected := &types.MsgLazyRegisterOrder{
			User:  user,
			Order: orderAny,
		}
		_, err = srv.LazyRegisterOrder(ctx, expected)
		require.NoError(t, err)
		rst, found := k.GetOrder(ctx,
			hash,
		)
		require.True(t, found)
		require.Equal(t, expected.User, rst.Owner)
	}
}

func TestOrderMsgServerDelete(t *testing.T) {
	user := "A"

	tests := []struct {
		desc    string
		request *types.MsgCancelOrder
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgCancelOrder{User: user,
				OrderHash: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgCancelOrder{User: "B",
				OrderHash: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgCancelOrder{User: user,
				OrderHash: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.OrderKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			_, err := srv.LazyRegisterOrder(ctx, &types.MsgLazyRegisterOrder{
				User: user,
				// Hash: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.CancelOrder(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetOrder(ctx,
					tc.request.OrderHash,
				)
				require.False(t, found)
			}
		})
	}
}

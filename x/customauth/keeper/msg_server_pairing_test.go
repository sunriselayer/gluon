package keeper_test

import (
	"strconv"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "gluon/testutil/keeper"
	"gluon/x/customauth/keeper"
	"gluon/x/customauth/types"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestPairingMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.CustomauthKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	user := "A"
	for i := 0; i < 5; i++ {
		pairingAny := codectypes.Any{}
		hash, err := k.GetPairingIndex(pairingAny)
		require.NoError(t, err)
		expected := &types.MsgCreatePairing{
			User:      user,
			PublicKey: pairingAny,
		}
		_, err = srv.CreatePairing(ctx, expected)
		require.NoError(t, err)
		rst, found := k.GetPairing(ctx,
			user,
			hash,
		)
		require.True(t, found)
		require.Equal(t, expected.User, rst.Owner)
	}
}

func TestPairingMsgServerDelete(t *testing.T) {
	user := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeletePairing
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeletePairing{User: user,
				PairingIndex: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeletePairing{User: "B",
				PairingIndex: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeletePairing{User: user,
				PairingIndex: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.CustomauthKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			_, err := srv.CreatePairing(ctx, &types.MsgCreatePairing{
				User: user,
				// Hash: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeletePairing(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetPairing(ctx,
					tc.request.User,
					tc.request.PairingIndex,
				)
				require.False(t, found)
			}
		})
	}
}

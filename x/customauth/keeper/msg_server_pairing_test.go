package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"gluon/x/customauth/types"
)

func TestPairingMsgServerCreate(t *testing.T) {
	_, srv, ctx := setupMsgServer(t)
	wctx := sdk.UnwrapSDKContext(ctx)

	user := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreatePairing(wctx, &types.MsgCreatePairing{User: user})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.PairingId))
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
			desc:    "Completed",
			request: &types.MsgDeletePairing{User: user},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeletePairing{User: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeletePairing{User: user, PairingId: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			_, srv, ctx := setupMsgServer(t)
			wctx := sdk.UnwrapSDKContext(ctx)

			_, err := srv.CreatePairing(wctx, &types.MsgCreatePairing{User: user})
			require.NoError(t, err)
			_, err = srv.DeletePairing(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

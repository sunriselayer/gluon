package keeper_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "gluon/testutil/keeper"
	"gluon/testutil/nullify"
	"gluon/x/perp/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestCrossMarginQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.PerpKeeper(t)
	msgs := createNCrossMargin(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryCrossMarginRequest
		response *types.QueryCrossMarginResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryCrossMarginRequest{
				Owner: msgs[0].Owner,
			},
			response: &types.QueryCrossMarginResponse{CrossMargin: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryCrossMarginRequest{
				Owner: msgs[1].Owner,
			},
			response: &types.QueryCrossMarginResponse{CrossMargin: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryCrossMarginRequest{
				Owner: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.CrossMargin(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

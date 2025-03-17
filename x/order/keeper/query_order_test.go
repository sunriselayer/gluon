package keeper_test

import (
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "gluon/testutil/keeper"
	"gluon/testutil/nullify"
	"gluon/x/order/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestOrderQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.OrderKeeper(t)
	msgs := createNOrder(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryOrderRequest
		response *types.QueryOrderResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryOrderRequest{
				Hash: msgs[0].Hash,
			},
			response: &types.QueryOrderResponse{Order: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryOrderRequest{
				Hash: msgs[1].Hash,
			},
			response: &types.QueryOrderResponse{Order: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryOrderRequest{
				Hash: strconv.Itoa(100000),
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
			response, err := keeper.Order(ctx, tc.request)
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

func TestOrderQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.OrderKeeper(t)
	msgs := createNOrder(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryOrdersRequest {
		return &types.QueryOrdersRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.Orders(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Orders), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Orders),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.Orders(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Orders), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Orders),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.Orders(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Orders),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.Orders(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}

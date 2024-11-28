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
	"gluon/x/contract/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestSortedOrderQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.ContractKeeper(t)
	msgs := createNSortedOrder(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetSortedOrderRequest
		response *types.QueryGetSortedOrderResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetSortedOrderRequest{
				Expiry: msgs[0].Expiry,
				Id:     msgs[0].Id,
			},
			response: &types.QueryGetSortedOrderResponse{SortedOrders: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetSortedOrderRequest{
				Expiry: msgs[1].Expiry,
				Id:     msgs[1].Id,
			},
			response: &types.QueryGetSortedOrderResponse{SortedOrders: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetSortedOrderRequest{
				Expiry: uint64(100000),
				Id:     strconv.Itoa(100000),
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
			response, err := keeper.SortedOrder(ctx, tc.request)
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

func TestSortedOrderQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.ContractKeeper(t)
	msgs := createNSortedOrder(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QuerySortedOrdersRequest {
		return &types.QuerySortedOrdersRequest{
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
			resp, err := keeper.SortedOrders(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.SortedOrders), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.SortedOrders),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.SortedOrders(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.SortedOrders), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.SortedOrders),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.SortedOrders(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.SortedOrders),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.SortedOrders(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}

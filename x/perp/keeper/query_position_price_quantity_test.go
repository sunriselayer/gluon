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
	"gluon/x/perp/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestPositionPriceQuantityQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.PerpKeeper(t)
	msgs := createNPositionPriceQuantity(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryPositionPriceQuantityRequest
		response *types.QueryPositionPriceQuantityResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryPositionPriceQuantityRequest{
				Owner:             msgs[0].Owner,
				PositionOrderHash: msgs[0].PositionOrderHash,
				Price:             msgs[0].Price,
			},
			response: &types.QueryPositionPriceQuantityResponse{PositionPriceQuantity: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryPositionPriceQuantityRequest{
				Owner:             msgs[1].Owner,
				PositionOrderHash: msgs[1].PositionOrderHash,
				Price:             msgs[1].Price,
			},
			response: &types.QueryPositionPriceQuantityResponse{PositionPriceQuantity: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryPositionPriceQuantityRequest{
				Owner:             strconv.Itoa(100000),
				PositionOrderHash: strconv.Itoa(100000),
				Price:             strconv.Itoa(100000),
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
			response, err := keeper.PositionPriceQuantity(ctx, tc.request)
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

func TestPositionPriceQuantityQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.PerpKeeper(t)
	msgs := createNPositionPriceQuantity(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryPositionPriceQuantitiesRequest {
		return &types.QueryPositionPriceQuantitiesRequest{
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
			resp, err := keeper.PositionPriceQuantities(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.PositionPriceQuantity), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.PositionPriceQuantity),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.PositionPriceQuantities(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.PositionPriceQuantity), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.PositionPriceQuantity),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.PositionPriceQuantities(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.PositionPriceQuantity),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.PositionPriceQuantities(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
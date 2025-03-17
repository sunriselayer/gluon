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

func TestFundingRateQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.PerpKeeper(t)
	msgs := createNFundingRate(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryFundingRateRequest
		response *types.QueryFundingRateResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryFundingRateRequest{
				DenomBase:  msgs[0].DenomBase,
				DenomQuote: msgs[0].DenomQuote,
				Seconds:    msgs[0].Seconds,
				Nanos:      msgs[0].Nanos,
			},
			response: &types.QueryFundingRateResponse{FundingRate: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryFundingRateRequest{
				DenomBase:  msgs[1].DenomBase,
				DenomQuote: msgs[1].DenomQuote,
				Seconds:    msgs[1].Seconds,
				Nanos:      msgs[1].Nanos,
			},
			response: &types.QueryFundingRateResponse{FundingRate: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryFundingRateRequest{
				DenomBase:  strconv.Itoa(100000),
				DenomQuote: strconv.Itoa(100000),
				Seconds:    100000,
				Nanos:      100000,
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
			response, err := keeper.FundingRate(ctx, tc.request)
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

func TestFundingRateQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.PerpKeeper(t)
	msgs := createNFundingRate(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryFundingRatesRequest {
		return &types.QueryFundingRatesRequest{
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
			resp, err := keeper.FundingRates(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.FundingRates), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.FundingRates),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.FundingRates(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.FundingRates), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.FundingRates),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.FundingRates(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.FundingRates),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.FundingRates(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}

package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "gluon/testutil/keeper"
	"gluon/testutil/nullify"
	"gluon/x/perp/keeper"
	"gluon/x/perp/types"

	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNFundingRate(keeper keeper.Keeper, ctx context.Context, n int) []types.FundingRate {
	items := make([]types.FundingRate, n)
	for i := range items {
		items[i].DenomBase = strconv.Itoa(i)
		items[i].DenomQuote = strconv.Itoa(i)
		items[i].Seconds = uint64(i)
		items[i].Nanos = uint32(i)

		keeper.SetFundingRate(ctx, items[i])
	}
	return items
}

func TestFundingRateGet(t *testing.T) {
	keeper, ctx := keepertest.PerpKeeper(t)
	items := createNFundingRate(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetFundingRate(ctx,
			item.DenomBase,
			item.DenomQuote,
			item.Seconds,
			item.Nanos,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestFundingRateRemove(t *testing.T) {
	keeper, ctx := keepertest.PerpKeeper(t)
	items := createNFundingRate(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveFundingRate(ctx,
			item.DenomBase,
			item.DenomQuote,
			item.Seconds,
			item.Nanos,
		)
		_, found := keeper.GetFundingRate(ctx,
			item.DenomBase,
			item.DenomQuote,
			item.Seconds,
			item.Nanos,
		)
		require.False(t, found)
	}
}

func TestFundingRateGetAll(t *testing.T) {
	keeper, ctx := keepertest.PerpKeeper(t)
	items := createNFundingRate(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllFundingRate(ctx)),
	)
}

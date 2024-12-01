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

func createNCrossMargin(keeper keeper.Keeper, ctx context.Context, n int) []types.CrossMargin {
	items := make([]types.CrossMargin, n)
	for i := range items {
		items[i].Owner = strconv.Itoa(i)

		keeper.SetCrossMargin(ctx, items[i])
	}
	return items
}

func TestCrossMarginGet(t *testing.T) {
	keeper, ctx := keepertest.PerpKeeper(t)
	items := createNCrossMargin(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetCrossMargin(ctx,
			item.Owner,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestCrossMarginRemove(t *testing.T) {
	keeper, ctx := keepertest.PerpKeeper(t)
	items := createNCrossMargin(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveCrossMargin(ctx,
			item.Owner,
		)
		_, found := keeper.GetCrossMargin(ctx,
			item.Owner,
		)
		require.False(t, found)
	}
}

func TestCrossMarginGetAll(t *testing.T) {
	keeper, ctx := keepertest.PerpKeeper(t)
	items := createNCrossMargin(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllCrossMargin(ctx)),
	)
}

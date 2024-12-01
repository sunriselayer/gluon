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

func createNPosition(keeper keeper.Keeper, ctx context.Context, n int) []types.Position {
	items := make([]types.Position, n)
	for i := range items {
		items[i].OrderHash = strconv.Itoa(i)

		keeper.SetPosition(ctx, items[i])
	}
	return items
}

func TestPositionGet(t *testing.T) {
	keeper, ctx := keepertest.PerpKeeper(t)
	items := createNPosition(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetPosition(ctx,
			item.Owner,
			item.OrderHash,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestPositionRemove(t *testing.T) {
	keeper, ctx := keepertest.PerpKeeper(t)
	items := createNPosition(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePosition(ctx,
			item.Owner,
			item.OrderHash,
		)
		_, found := keeper.GetPosition(ctx,
			item.Owner,
			item.OrderHash,
		)
		require.False(t, found)
	}
}

func TestPositionGetAll(t *testing.T) {
	keeper, ctx := keepertest.PerpKeeper(t)
	items := createNPosition(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPosition(ctx)),
	)
}

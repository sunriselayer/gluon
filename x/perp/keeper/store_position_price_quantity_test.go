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

func createNPositionPriceQuantity(keeper keeper.Keeper, ctx context.Context, n int) []types.PositionPriceQuantity {
	items := make([]types.PositionPriceQuantity, n)
	for i := range items {
		items[i].Owner = strconv.Itoa(i)
		items[i].PositionOrderHash = strconv.Itoa(i)
		items[i].Price = strconv.Itoa(i)

		keeper.SetPositionPriceQuantity(ctx, items[i])
	}
	return items
}

func TestPositionPriceQuantityGet(t *testing.T) {
	keeper, ctx := keepertest.PerpKeeper(t)
	items := createNPositionPriceQuantity(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetPositionPriceQuantity(ctx,
			item.Owner,
			item.PositionOrderHash,
			item.Price,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestPositionPriceQuantityRemove(t *testing.T) {
	keeper, ctx := keepertest.PerpKeeper(t)
	items := createNPositionPriceQuantity(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePositionPriceQuantity(ctx,
			item.Owner,
			item.PositionOrderHash,
			item.Price,
		)
		_, found := keeper.GetPositionPriceQuantity(ctx,
			item.Owner,
			item.PositionOrderHash,
			item.Price,
		)
		require.False(t, found)
	}
}

func TestPositionPriceQuantityGetAll(t *testing.T) {
	keeper, ctx := keepertest.PerpKeeper(t)
	items := createNPositionPriceQuantity(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPositionPriceQuantity(ctx)),
	)
}

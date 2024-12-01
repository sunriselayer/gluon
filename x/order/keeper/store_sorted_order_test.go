package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "gluon/testutil/keeper"
	"gluon/testutil/nullify"
	"gluon/x/order/keeper"
	"gluon/x/order/types"

	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNSortedOrder(keeper keeper.Keeper, ctx context.Context, n int) []types.SortedOrder {
	items := make([]types.SortedOrder, n)
	for i := range items {
		items[i].OrderHash = strconv.Itoa(i)

		keeper.SetSortedOrder(ctx, items[i])
	}
	return items
}

func TestSortedOrderGet(t *testing.T) {
	keeper, ctx := keepertest.OrderKeeper(t)
	items := createNSortedOrder(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetSortedOrder(ctx,
			item.Seconds,
			item.Nanos,
			item.OrderHash,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestSortedOrderRemove(t *testing.T) {
	keeper, ctx := keepertest.OrderKeeper(t)
	items := createNSortedOrder(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveSortedOrder(ctx,
			item.Seconds,
			item.Nanos,
			item.OrderHash,
		)
		_, found := keeper.GetSortedOrder(ctx,
			item.Seconds,
			item.Nanos,
			item.OrderHash,
		)
		require.False(t, found)
	}
}

func TestSortedOrderGetAll(t *testing.T) {
	keeper, ctx := keepertest.OrderKeeper(t)
	items := createNSortedOrder(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllSortedOrder(ctx)),
	)
}

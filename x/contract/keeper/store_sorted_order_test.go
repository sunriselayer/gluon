package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "gluon/testutil/keeper"
	"gluon/testutil/nullify"
	"gluon/x/contract/keeper"
	"gluon/x/contract/types"

	"github.com/stretchr/testify/require"

	"time"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNSortedOrder(keeper keeper.Keeper, ctx context.Context, n int) []types.SortedOrder {
	items := make([]types.SortedOrder, n)
	for i := range items {
		items[i].Expiry = uint64(i)
		items[i].Id = strconv.Itoa(i)

		keeper.SetSortedOrder(ctx, items[i])
	}
	return items
}

func TestSortedOrderGet(t *testing.T) {
	keeper, ctx := keepertest.ContractKeeper(t)
	items := createNSortedOrder(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetSortedOrder(ctx,
			time.UnixMilli(int64(item.Expiry)),
			item.Id,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestSortedOrderRemove(t *testing.T) {
	keeper, ctx := keepertest.ContractKeeper(t)
	items := createNSortedOrder(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveSortedOrder(ctx,
			item.Expiry,
			item.Id,
		)
		_, found := keeper.GetSortedOrder(ctx,
			time.UnixMilli(int64(item.Expiry)),
			item.Id,
		)
		require.False(t, found)
	}
}

func TestSortedOrderGetAll(t *testing.T) {
	keeper, ctx := keepertest.ContractKeeper(t)
	items := createNSortedOrder(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllSortedOrder(ctx)),
	)
}

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
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNOrder(keeper keeper.Keeper, ctx context.Context, n int) []types.Order {
	items := make([]types.Order, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetOrder(ctx, items[i])
	}
	return items
}

func TestOrderGet(t *testing.T) {
	keeper, ctx := keepertest.ContractKeeper(t)
	items := createNOrder(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetOrder(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestOrderRemove(t *testing.T) {
	keeper, ctx := keepertest.ContractKeeper(t)
	items := createNOrder(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveOrder(ctx,
			item.Index,
		)
		_, found := keeper.GetOrder(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestOrderGetAll(t *testing.T) {
	keeper, ctx := keepertest.ContractKeeper(t)
	items := createNOrder(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllOrder(ctx)),
	)
}

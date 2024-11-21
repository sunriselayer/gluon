package keeper_test

import (
	"context"
	"testing"

	keepertest "gluon/testutil/keeper"
	"gluon/testutil/nullify"
	"gluon/x/contract/keeper"
	"gluon/x/contract/types"

	"github.com/stretchr/testify/require"
)

func createNLazySettlement(keeper keeper.Keeper, ctx context.Context, n int) []types.LazySettlement {
	items := make([]types.LazySettlement, n)
	for i := range items {
		items[i].Id = keeper.AppendLazySettlement(ctx, items[i])
	}
	return items
}

func TestLazySettlementGet(t *testing.T) {
	keeper, ctx := keepertest.ContractKeeper(t)
	items := createNLazySettlement(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetLazySettlement(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestLazySettlementRemove(t *testing.T) {
	keeper, ctx := keepertest.ContractKeeper(t)
	items := createNLazySettlement(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveLazySettlement(ctx, item.Id)
		_, found := keeper.GetLazySettlement(ctx, item.Id)
		require.False(t, found)
	}
}

func TestLazySettlementGetAll(t *testing.T) {
	keeper, ctx := keepertest.ContractKeeper(t)
	items := createNLazySettlement(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllLazySettlement(ctx)),
	)
}

func TestLazySettlementCount(t *testing.T) {
	keeper, ctx := keepertest.ContractKeeper(t)
	items := createNLazySettlement(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetLazySettlementCount(ctx))
}

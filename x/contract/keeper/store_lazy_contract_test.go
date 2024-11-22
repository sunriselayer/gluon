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

func createNLazyContract(keeper keeper.Keeper, ctx context.Context, n int) []types.LazyContract {
	items := make([]types.LazyContract, n)
	for i := range items {
		items[i].Id = keeper.AppendLazyContract(ctx, items[i])
	}
	return items
}

func TestLazyContractGet(t *testing.T) {
	keeper, ctx := keepertest.ContractKeeper(t)
	items := createNLazyContract(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetLazyContract(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestLazyContractRemove(t *testing.T) {
	keeper, ctx := keepertest.ContractKeeper(t)
	items := createNLazyContract(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveLazyContract(ctx, item.Id)
		_, found := keeper.GetLazyContract(ctx, item.Id)
		require.False(t, found)
	}
}

func TestLazyContractGetAll(t *testing.T) {
	keeper, ctx := keepertest.ContractKeeper(t)
	items := createNLazyContract(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllLazyContract(ctx)),
	)
}

func TestLazyContractCount(t *testing.T) {
	keeper, ctx := keepertest.ContractKeeper(t)
	items := createNLazyContract(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetLazyContractCount(ctx))
}

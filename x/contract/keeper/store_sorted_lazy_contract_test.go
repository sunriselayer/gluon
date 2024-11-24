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

func createNSortedLazyContract(keeper keeper.Keeper, ctx context.Context, n int) []types.SortedLazyContract {
	items := make([]types.SortedLazyContract, n)
	for i := range items {
		items[i].Expiry = uint64(i)
		items[i].Id = uint64(i)

		keeper.SetSortedLazyContract(ctx, items[i])
	}
	return items
}

func TestSortedLazyContractGet(t *testing.T) {
	keeper, ctx := keepertest.ContractKeeper(t)
	items := createNSortedLazyContract(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetSortedLazyContract(ctx,
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
func TestSortedLazyContractRemove(t *testing.T) {
	keeper, ctx := keepertest.ContractKeeper(t)
	items := createNSortedLazyContract(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveSortedLazyContract(ctx,
			item.Expiry,
			item.Id,
		)
		_, found := keeper.GetSortedLazyContract(ctx,
			time.UnixMilli(int64(item.Expiry)),
			item.Id,
		)
		require.False(t, found)
	}
}

func TestSortedLazyContractGetAll(t *testing.T) {
	keeper, ctx := keepertest.ContractKeeper(t)
	items := createNSortedLazyContract(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllSortedLazyContract(ctx)),
	)
}

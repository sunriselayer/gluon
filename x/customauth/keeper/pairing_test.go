package keeper_test

import (
	"context"
	"testing"

	keepertest "gluon/testutil/keeper"
	"gluon/testutil/nullify"
	"gluon/x/customauth/keeper"
	"gluon/x/customauth/types"

	"github.com/stretchr/testify/require"
)

func createNPairing(keeper keeper.Keeper, ctx context.Context, n int) []types.Pairing {
	items := make([]types.Pairing, n)
	for i := range items {
		items[i].Id = keeper.AppendPairing(ctx, items[i])
	}
	return items
}

func TestPairingGet(t *testing.T) {
	keeper, ctx := keepertest.CustomauthKeeper(t)
	items := createNPairing(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetPairing(ctx, item.Address, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestPairingRemove(t *testing.T) {
	keeper, ctx := keepertest.CustomauthKeeper(t)
	items := createNPairing(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePairing(ctx, item.Address, item.Id)
		_, found := keeper.GetPairing(ctx, item.Address, item.Id)
		require.False(t, found)
	}
}

func TestPairingGetAll(t *testing.T) {
	keeper, ctx := keepertest.CustomauthKeeper(t)
	items := createNPairing(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPairing(ctx)),
	)
}

func TestPairingCount(t *testing.T) {
	keeper, ctx := keepertest.CustomauthKeeper(t)
	items := createNPairing(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetPairingCount(ctx))
}

package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "gluon/testutil/keeper"
	"gluon/testutil/nullify"
	"gluon/x/customauth/keeper"
	"gluon/x/customauth/types"

	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNPairing(keeper keeper.Keeper, ctx context.Context, n int) []types.Pairing {
	items := make([]types.Pairing, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetPairing(ctx, items[i])
	}
	return items
}

func TestPairingGet(t *testing.T) {
	keeper, ctx := keepertest.CustomauthKeeper(t)
	items := createNPairing(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetPairing(ctx,
			item.Owner,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestPairingRemove(t *testing.T) {
	keeper, ctx := keepertest.CustomauthKeeper(t)
	items := createNPairing(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePairing(ctx,
			item.Owner,
			item.Index,
		)
		_, found := keeper.GetPairing(ctx,
			item.Owner,
			item.Index,
		)
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

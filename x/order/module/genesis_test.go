package order_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "gluon/testutil/keeper"
	"gluon/testutil/nullify"
	order "gluon/x/order/module"
	"gluon/x/order/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		Orders: []types.Order{
			{
				Hash: "0",
			},
			{
				Hash: "1",
			},
		},
		SortedOrders: []types.SortedOrder{
			{
				OrderHash: "0",
			},
			{
				OrderHash: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.OrderKeeper(t)
	order.InitGenesis(ctx, k, genesisState)
	got := order.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.Orders, got.Orders)
	require.ElementsMatch(t, genesisState.SortedOrders, got.SortedOrders)
	// this line is used by starport scaffolding # genesis/test/assert
}

package contract_test

import (
	"testing"

	keepertest "gluon/testutil/keeper"
	"gluon/testutil/nullify"
	contract "gluon/x/contract/module"
	"gluon/x/contract/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		Orders: []types.Order{
			{
				Id: "0",
			},
			{
				Id: "1",
			},
		},
		SortedOrders: []types.SortedOrder{
			{
				Expiry: 0,
				Id:     "0",
			},
			{
				Expiry: 1,
				Id:     "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ContractKeeper(t)
	contract.InitGenesis(ctx, k, genesisState)
	got := contract.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.Orders, got.Orders)
	require.ElementsMatch(t, genesisState.SortedOrders, got.SortedOrders)
	// this line is used by starport scaffolding # genesis/test/assert
}

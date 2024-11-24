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
		PortId: types.PortID,
		OrderList: []types.Order{
			{
				Id: "0",
			},
			{
				Id: "1",
			},
		},
		SortedOrderList: []types.SortedOrder{
			{
				Expiry: 0,
				Id:     "0",
			},
			{
				Expiry: 1,
				Id:     "1",
			},
		},
		LazyContractList: []types.LazyContract{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		LazyContractCount: 2,
		SortedLazyContractList: []types.SortedLazyContract{
			{
				Expiry: 0,
				Id:     0,
			},
			{
				Expiry: 1,
				Id:     1,
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

	require.Equal(t, genesisState.PortId, got.PortId)

	require.ElementsMatch(t, genesisState.OrderList, got.OrderList)
	require.ElementsMatch(t, genesisState.SortedOrderList, got.SortedOrderList)
	require.ElementsMatch(t, genesisState.LazyContractList, got.LazyContractList)
	require.Equal(t, genesisState.LazyContractCount, got.LazyContractCount)
	require.ElementsMatch(t, genesisState.SortedLazyContractList, got.SortedLazyContractList)
	// this line is used by starport scaffolding # genesis/test/assert
}

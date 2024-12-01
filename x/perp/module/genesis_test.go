package perp_test

import (
	"testing"

	keepertest "gluon/testutil/keeper"
	"gluon/testutil/nullify"
	perp "gluon/x/perp/module"
	"gluon/x/perp/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		PositionList: []types.Position{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		PositionCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PerpKeeper(t)
	perp.InitGenesis(ctx, k, genesisState)
	got := perp.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.PositionList, got.PositionList)
	require.Equal(t, genesisState.PositionCount, got.PositionCount)
	// this line is used by starport scaffolding # genesis/test/assert
}

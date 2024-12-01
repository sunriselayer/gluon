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

		Positions: []types.Position{
			{
				Owner:     "0",
				OrderHash: "0",
			},
			{
				Owner:     "1",
				OrderHash: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PerpKeeper(t)
	perp.InitGenesis(ctx, k, genesisState)
	got := perp.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.Positions, got.Positions)
	// this line is used by starport scaffolding # genesis/test/assert
}

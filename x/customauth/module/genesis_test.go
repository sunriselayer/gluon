package customauth_test

import (
	"testing"

	keepertest "gluon/testutil/keeper"
	"gluon/testutil/nullify"
	customauth "gluon/x/customauth/module"
	"gluon/x/customauth/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		Pairings: []types.Pairing{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CustomauthKeeper(t)
	customauth.InitGenesis(ctx, k, genesisState)
	got := customauth.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.Pairings, got.Pairings)
	// this line is used by starport scaffolding # genesis/test/assert
}

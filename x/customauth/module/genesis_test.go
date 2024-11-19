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

		PairingList: []types.Pairing{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		PairingCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CustomauthKeeper(t)
	customauth.InitGenesis(ctx, k, genesisState)
	got := customauth.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.PairingList, got.PairingList)
	require.Equal(t, genesisState.PairingCount, got.PairingCount)
	// this line is used by starport scaffolding # genesis/test/assert
}

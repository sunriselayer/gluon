package types_test

import (
	"testing"

	"gluon/x/order/types"

	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

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
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated order",
			genState: &types.GenesisState{
				Orders: []types.Order{
					{
						Hash: "0",
					},
					{
						Hash: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated sortedOrder",
			genState: &types.GenesisState{
				SortedOrders: []types.SortedOrder{
					{
						OrderHash: "0",
					},
					{
						OrderHash: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}

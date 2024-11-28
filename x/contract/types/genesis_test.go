package types_test

import (
	"testing"

	"gluon/x/contract/types"

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
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated order",
			genState: &types.GenesisState{
				Orders: []types.Order{
					{
						Id: "0",
					},
					{
						Id: "0",
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
						Expiry: 0,
						Id:     "0",
					},
					{
						Expiry: 0,
						Id:     "0",
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

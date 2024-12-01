package types_test

import (
	"testing"

	"gluon/x/perp/types"

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
				PositionPriceQuantities: []types.PositionPriceQuantity{
					{
						Owner:             "0",
						PositionOrderHash: "0",
						Price:             "0",
					},
					{
						Owner:             "1",
						PositionOrderHash: "1",
						Price:             "1",
					},
				},
				CrossMargins: []types.CrossMargin{
					{
						Owner: "0",
					},
					{
						Owner: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated position",
			genState: &types.GenesisState{
				Positions: []types.Position{
					{
						Owner:     "0",
						OrderHash: "0",
					},
					{
						Owner:     "0",
						OrderHash: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated positionPriceQuantity",
			genState: &types.GenesisState{
				PositionPriceQuantities: []types.PositionPriceQuantity{
					{
						Owner:             "0",
						PositionOrderHash: "0",
						Price:             "0",
					},
					{
						Owner:             "0",
						PositionOrderHash: "0",
						Price:             "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated crossMargin",
			genState: &types.GenesisState{
				CrossMargins: []types.CrossMargin{
					{
						Owner: "0",
					},
					{
						Owner: "0",
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

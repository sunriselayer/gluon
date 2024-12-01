package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Positions: []Position{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in position
	positionIndexMap := make(map[string]struct{})

	for _, elem := range gs.Positions {
		index := string(PositionKey(elem.Owner, elem.OrderHash))
		if _, ok := positionIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for position")
		}
		positionIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}

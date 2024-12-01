package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PositionList: []Position{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in position
	positionIdMap := make(map[uint64]bool)
	positionCount := gs.GetPositionCount()
	for _, elem := range gs.PositionList {
		if _, ok := positionIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for position")
		}
		if elem.Id >= positionCount {
			return fmt.Errorf("position id should be lower or equal than the last id")
		}
		positionIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}

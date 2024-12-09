package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Pairings: []Pairing{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in pairing
	pairingIdMap := make(map[string]bool)
	for _, elem := range gs.Pairings {
		if _, ok := pairingIdMap[elem.Index]; ok {
			return fmt.Errorf("duplicated index for pairing")
		}
		pairingIdMap[elem.Index] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}

package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PairingList: []Pairing{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in pairing
	pairingIdMap := make(map[uint64]bool)
	pairingCount := gs.GetPairingCount()
	for _, elem := range gs.PairingList {
		if _, ok := pairingIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for pairing")
		}
		if elem.Id >= pairingCount {
			return fmt.Errorf("pairing id should be lower or equal than the last id")
		}
		pairingIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}

package types

import (
	"fmt"
)

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Orders:       []Order{},
		SortedOrders: []SortedOrder{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in order
	orderIndexMap := make(map[string]struct{})

	for _, elem := range gs.Orders {
		index := string(OrderKey(elem.Id))
		if _, ok := orderIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for order")
		}
		orderIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in sortedOrder
	sortedOrderIndexMap := make(map[string]struct{})

	for _, elem := range gs.SortedOrders {
		index := string(SortedOrderKey(elem.Expiry, elem.Id))
		if _, ok := sortedOrderIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for sortedOrder")
		}
		sortedOrderIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}

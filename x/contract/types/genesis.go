package types

import (
	"fmt"

	host "github.com/cosmos/ibc-go/v8/modules/core/24-host"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PortId:                 PortID,
		OrderList:              []Order{},
		SortedOrderList:        []SortedOrder{},
		LazyContractList:       []LazyContract{},
		SortedLazyContractList: []SortedLazyContract{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	if err := host.PortIdentifierValidator(gs.PortId); err != nil {
		return err
	}
	// Check for duplicated index in order
	orderIndexMap := make(map[string]struct{})

	for _, elem := range gs.OrderList {
		index := string(OrderKey(elem.Id))
		if _, ok := orderIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for order")
		}
		orderIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in sortedOrder
	sortedOrderIndexMap := make(map[string]struct{})

	for _, elem := range gs.SortedOrderList {
		index := string(SortedOrderKey(elem.Expiry, elem.Id))
		if _, ok := sortedOrderIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for sortedOrder")
		}
		sortedOrderIndexMap[index] = struct{}{}
	}
	// Check for duplicated ID in lazyContract
	lazyContractIdMap := make(map[uint64]bool)
	lazyContractCount := gs.GetLazyContractCount()
	for _, elem := range gs.LazyContractList {
		if _, ok := lazyContractIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for lazyContract")
		}
		if elem.Id >= lazyContractCount {
			return fmt.Errorf("lazyContract id should be lower or equal than the last id")
		}
		lazyContractIdMap[elem.Id] = true
	}
	// Check for duplicated index in sortedLazyContract
	sortedLazyContractIndexMap := make(map[string]struct{})

	for _, elem := range gs.SortedLazyContractList {
		index := string(SortedLazyContractKey(elem.Expiry, elem.Index))
		if _, ok := sortedLazyContractIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for sortedLazyContract")
		}
		sortedLazyContractIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}

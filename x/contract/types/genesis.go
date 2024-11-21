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
		PortId:             PortID,
		OrderList:          []Order{},
		SortedOrderList:    []SortedOrder{},
		LazySettlementList: []LazySettlement{},
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
	// Check for duplicated ID in lazySettlement
	lazySettlementIdMap := make(map[uint64]bool)
	lazySettlementCount := gs.GetLazySettlementCount()
	for _, elem := range gs.LazySettlementList {
		if _, ok := lazySettlementIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for lazySettlement")
		}
		if elem.Id >= lazySettlementCount {
			return fmt.Errorf("lazySettlement id should be lower or equal than the last id")
		}
		lazySettlementIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}

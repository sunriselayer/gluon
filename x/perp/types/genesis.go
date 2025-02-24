package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Positions:               []Position{},
		PositionPriceQuantities: []PositionPriceQuantity{},
		CrossMargins:            []CrossMargin{},
		FundingRates:            []FundingRate{},
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
	// Check for duplicated index in positionPriceQuantity
	positionPriceQuantityIndexMap := make(map[string]struct{})

	for _, elem := range gs.PositionPriceQuantities {
		index := string(PositionPriceQuantityKey(elem.Owner, elem.PositionOrderHash, elem.Price))
		if _, ok := positionPriceQuantityIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for positionPriceQuantity")
		}
		positionPriceQuantityIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in crossMargin
	crossMarginIndexMap := make(map[string]struct{})

	for _, elem := range gs.CrossMargins {
		index := string(CrossMarginKey(elem.Owner))
		if _, ok := crossMarginIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for crossMargin")
		}
		crossMarginIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in fundingRate
	fundingRateIndexMap := make(map[string]struct{})

	for _, elem := range gs.FundingRates {
		index := string(FundingRateKey(elem.DenomBase, elem.DenomQuote, elem.Seconds, elem.Nanos))
		if _, ok := fundingRateIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for fundingRate")
		}
		fundingRateIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}

package perp

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"gluon/x/perp/keeper"
	"gluon/x/perp/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the position
	for _, elem := range genState.Positions {
		k.SetPosition(ctx, elem)
	}

	// Set all the positionPriceQuantity
	for _, elem := range genState.PositionPriceQuantities {
		k.SetPositionPriceQuantity(ctx, elem)
	}
	// Set all the crossMargin
	for _, elem := range genState.CrossMargins {
		k.SetCrossMargin(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.Positions = k.GetAllPosition(ctx)
	genesis.PositionPriceQuantities = k.GetAllPositionPriceQuantity(ctx)
	genesis.CrossMargins = k.GetAllCrossMargin(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}

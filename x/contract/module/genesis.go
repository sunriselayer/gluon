package contract

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"gluon/x/contract/keeper"
	"gluon/x/contract/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the order
	for _, elem := range genState.Orders {
		k.SetOrder(ctx, elem)
	}
	// Set all the sortedOrder
	for _, elem := range genState.SortedOrders {
		k.SetSortedOrder(ctx, elem)
	}
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.Orders = k.GetAllOrder(ctx)
	genesis.SortedOrders = k.GetAllSortedOrder(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}

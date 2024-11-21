package contract

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"gluon/x/contract/keeper"
	"gluon/x/contract/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the order
	for _, elem := range genState.OrderList {
		k.SetOrder(ctx, elem)
	}
	// Set all the sortedOrder
	for _, elem := range genState.SortedOrderList {
		k.SetSortedOrder(ctx, elem)
	}
	// Set all the lazySettlement
	for _, elem := range genState.LazySettlementList {
		k.SetLazySettlement(ctx, elem)
	}

	// Set lazySettlement count
	k.SetLazySettlementCount(ctx, genState.LazySettlementCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetPort(ctx, genState.PortId)
	// Only try to bind to port if it is not already bound, since we may already own
	// port capability from capability InitGenesis
	if k.ShouldBound(ctx, genState.PortId) {
		// module binds to the port on InitChain
		// and claims the returned capability
		err := k.BindPort(ctx, genState.PortId)
		if err != nil {
			panic("could not claim port capability: " + err.Error())
		}
	}
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.PortId = k.GetPort(ctx)
	genesis.OrderList = k.GetAllOrder(ctx)
	genesis.SortedOrderList = k.GetAllSortedOrder(ctx)
	genesis.LazySettlementList = k.GetAllLazySettlement(ctx)
	genesis.LazySettlementCount = k.GetLazySettlementCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}

package customauth

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"gluon/x/customauth/keeper"
	"gluon/x/customauth/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the pairing
	for _, elem := range genState.PairingList {
		k.SetPairing(ctx, elem)
	}

	// Set pairing count
	k.SetPairingCount(ctx, genState.PairingCount)
	// this line is used by starport scaffolding # genesis/module/init
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.PairingList = k.GetAllPairing(ctx)
	genesis.PairingCount = k.GetPairingCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}

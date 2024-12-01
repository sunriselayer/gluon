package order

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"gluon/testutil/sample"
	ordersimulation "gluon/x/order/simulation"
	"gluon/x/order/types"
)

// avoid unused import issue
var (
	_ = ordersimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgLazyRegisterOrder = "op_weight_msg_order"
	// TODO: Determine the simulation weight value
	defaultWeightMsgLazyRegisterOrder int = 100

	opWeightMsgCancelOrder = "op_weight_msg_order"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCancelOrder int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	contractGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		Orders: []types.Order{
			{
				Hash:  "0",
				Owner: sample.AccAddress(),
			},
			{
				Hash:  "1",
				Owner: sample.AccAddress(),
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&contractGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgLazyRegisterOrder int
	simState.AppParams.GetOrGenerate(opWeightMsgLazyRegisterOrder, &weightMsgLazyRegisterOrder, nil,
		func(_ *rand.Rand) {
			weightMsgLazyRegisterOrder = defaultWeightMsgLazyRegisterOrder
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgLazyRegisterOrder,
		ordersimulation.SimulateMsgLazyRegisterOrder(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCancelOrder int
	simState.AppParams.GetOrGenerate(opWeightMsgCancelOrder, &weightMsgCancelOrder, nil,
		func(_ *rand.Rand) {
			weightMsgCancelOrder = defaultWeightMsgCancelOrder
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCancelOrder,
		ordersimulation.SimulateMsgCancelOrder(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgLazyRegisterOrder,
			defaultWeightMsgLazyRegisterOrder,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				ordersimulation.SimulateMsgLazyRegisterOrder(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCancelOrder,
			defaultWeightMsgCancelOrder,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				ordersimulation.SimulateMsgCancelOrder(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}

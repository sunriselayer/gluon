package contract

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"gluon/testutil/sample"
	contractsimulation "gluon/x/contract/simulation"
	"gluon/x/contract/types"
)

// avoid unused import issue
var (
	_ = contractsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgMatchOrder = "op_weight_msg_match_order"
	// TODO: Determine the simulation weight value
	defaultWeightMsgMatchOrder int = 100

	opWeightMsgMatchLazyOrder = "op_weight_msg_match_lazy_order"
	// TODO: Determine the simulation weight value
	defaultWeightMsgMatchLazyOrder int = 100

	opWeightMsgCreateOrder = "op_weight_msg_order"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateOrder int = 100

	opWeightMsgUpdateOrder = "op_weight_msg_order"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateOrder int = 100

	opWeightMsgDeleteOrder = "op_weight_msg_order"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteOrder int = 100

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
		PortId: types.PortID,
		OrderList: []types.Order{
			{
				Id:      "0",
				Address: sample.AccAddress(),
			},
			{
				Id:      "1",
				Address: sample.AccAddress(),
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

	var weightMsgMatchOrder int
	simState.AppParams.GetOrGenerate(opWeightMsgMatchOrder, &weightMsgMatchOrder, nil,
		func(_ *rand.Rand) {
			weightMsgMatchOrder = defaultWeightMsgMatchOrder
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgMatchOrder,
		contractsimulation.SimulateMsgMatchOrder(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgMatchLazyOrder int
	simState.AppParams.GetOrGenerate(opWeightMsgMatchLazyOrder, &weightMsgMatchLazyOrder, nil,
		func(_ *rand.Rand) {
			weightMsgMatchLazyOrder = defaultWeightMsgMatchLazyOrder
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgMatchLazyOrder,
		contractsimulation.SimulateMsgMatchLazyOrder(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateOrder int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateOrder, &weightMsgCreateOrder, nil,
		func(_ *rand.Rand) {
			weightMsgCreateOrder = defaultWeightMsgCreateOrder
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateOrder,
		contractsimulation.SimulateMsgCreateOrder(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteOrder int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteOrder, &weightMsgDeleteOrder, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteOrder = defaultWeightMsgDeleteOrder
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteOrder,
		contractsimulation.SimulateMsgDeleteOrder(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgMatchOrder,
			defaultWeightMsgMatchOrder,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				contractsimulation.SimulateMsgMatchOrder(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgMatchLazyOrder,
			defaultWeightMsgMatchLazyOrder,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				contractsimulation.SimulateMsgMatchLazyOrder(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateOrder,
			defaultWeightMsgCreateOrder,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				contractsimulation.SimulateMsgCreateOrder(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteOrder,
			defaultWeightMsgDeleteOrder,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				contractsimulation.SimulateMsgDeleteOrder(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}

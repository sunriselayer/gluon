package perp

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"gluon/testutil/sample"
	perpsimulation "gluon/x/perp/simulation"
	"gluon/x/perp/types"
)

// avoid unused import issue
var (
	_ = perpsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgMatchOrder = "op_weight_msg_match_order"
	// TODO: Determine the simulation weight value
	defaultWeightMsgMatchOrder int = 100

	opWeightMsgDepositCrossMargin = "op_weight_msg_deposit_cross_margin"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDepositCrossMargin int = 100

	opWeightMsgWithdrawCrossMargin = "op_weight_msg_withdraw_cross_margin"
	// TODO: Determine the simulation weight value
	defaultWeightMsgWithdrawCrossMargin int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	perpGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&perpGenesis)
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
		perpsimulation.SimulateMsgMatchOrder(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDepositCrossMargin int
	simState.AppParams.GetOrGenerate(opWeightMsgDepositCrossMargin, &weightMsgDepositCrossMargin, nil,
		func(_ *rand.Rand) {
			weightMsgDepositCrossMargin = defaultWeightMsgDepositCrossMargin
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDepositCrossMargin,
		perpsimulation.SimulateMsgDepositCrossMargin(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgWithdrawCrossMargin int
	simState.AppParams.GetOrGenerate(opWeightMsgWithdrawCrossMargin, &weightMsgWithdrawCrossMargin, nil,
		func(_ *rand.Rand) {
			weightMsgWithdrawCrossMargin = defaultWeightMsgWithdrawCrossMargin
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgWithdrawCrossMargin,
		perpsimulation.SimulateMsgWithdrawCrossMargin(am.accountKeeper, am.bankKeeper, am.keeper),
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
				perpsimulation.SimulateMsgMatchOrder(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDepositCrossMargin,
			defaultWeightMsgDepositCrossMargin,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				perpsimulation.SimulateMsgDepositCrossMargin(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgWithdrawCrossMargin,
			defaultWeightMsgWithdrawCrossMargin,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				perpsimulation.SimulateMsgWithdrawCrossMargin(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}

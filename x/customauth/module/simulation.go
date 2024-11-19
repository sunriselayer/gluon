package customauth

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"gluon/testutil/sample"
	customauthsimulation "gluon/x/customauth/simulation"
	"gluon/x/customauth/types"
)

// avoid unused import issue
var (
	_ = customauthsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreatePairing = "op_weight_msg_pairing"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreatePairing int = 100

	opWeightMsgUpdatePairing = "op_weight_msg_pairing"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdatePairing int = 100

	opWeightMsgDeletePairing = "op_weight_msg_pairing"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeletePairing int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	customauthGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		PairingList: []types.Pairing{
			{
				Id:      0,
				Address: sample.AccAddress(),
			},
			{
				Id:      1,
				Address: sample.AccAddress(),
			},
		},
		PairingCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&customauthGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreatePairing int
	simState.AppParams.GetOrGenerate(opWeightMsgCreatePairing, &weightMsgCreatePairing, nil,
		func(_ *rand.Rand) {
			weightMsgCreatePairing = defaultWeightMsgCreatePairing
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreatePairing,
		customauthsimulation.SimulateMsgCreatePairing(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeletePairing int
	simState.AppParams.GetOrGenerate(opWeightMsgDeletePairing, &weightMsgDeletePairing, nil,
		func(_ *rand.Rand) {
			weightMsgDeletePairing = defaultWeightMsgDeletePairing
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeletePairing,
		customauthsimulation.SimulateMsgDeletePairing(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreatePairing,
			defaultWeightMsgCreatePairing,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				customauthsimulation.SimulateMsgCreatePairing(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeletePairing,
			defaultWeightMsgDeletePairing,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				customauthsimulation.SimulateMsgDeletePairing(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}

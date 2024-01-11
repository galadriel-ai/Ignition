package agent

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"agent/testutil/sample"
	agentsimulation "agent/x/agent/simulation"
	"agent/x/agent/types"
)

// avoid unused import issue
var (
	_ = agentsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgAskQuestion = "op_weight_msg_ask_question"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAskQuestion int = 100

	opWeightMsgAnswerQuestion = "op_weight_msg_answer_question"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAnswerQuestion int = 100

	opWeightMsgAddPrompt = "op_weight_msg_add_prompt"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddPrompt int = 100

	opWeightMsgRunAgent = "op_weight_msg_run_agent"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRunAgent int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	agentGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&agentGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgAskQuestion int
	simState.AppParams.GetOrGenerate(opWeightMsgAskQuestion, &weightMsgAskQuestion, nil,
		func(_ *rand.Rand) {
			weightMsgAskQuestion = defaultWeightMsgAskQuestion
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAskQuestion,
		agentsimulation.SimulateMsgAskQuestion(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAnswerQuestion int
	simState.AppParams.GetOrGenerate(opWeightMsgAnswerQuestion, &weightMsgAnswerQuestion, nil,
		func(_ *rand.Rand) {
			weightMsgAnswerQuestion = defaultWeightMsgAnswerQuestion
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAnswerQuestion,
		agentsimulation.SimulateMsgAnswerQuestion(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAddPrompt int
	simState.AppParams.GetOrGenerate(opWeightMsgAddPrompt, &weightMsgAddPrompt, nil,
		func(_ *rand.Rand) {
			weightMsgAddPrompt = defaultWeightMsgAddPrompt
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddPrompt,
		agentsimulation.SimulateMsgAddPrompt(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRunAgent int
	simState.AppParams.GetOrGenerate(opWeightMsgRunAgent, &weightMsgRunAgent, nil,
		func(_ *rand.Rand) {
			weightMsgRunAgent = defaultWeightMsgRunAgent
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRunAgent,
		agentsimulation.SimulateMsgRunAgent(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgAskQuestion,
			defaultWeightMsgAskQuestion,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				agentsimulation.SimulateMsgAskQuestion(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgAnswerQuestion,
			defaultWeightMsgAnswerQuestion,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				agentsimulation.SimulateMsgAnswerQuestion(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgAddPrompt,
			defaultWeightMsgAddPrompt,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				agentsimulation.SimulateMsgAddPrompt(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgRunAgent,
			defaultWeightMsgRunAgent,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				agentsimulation.SimulateMsgRunAgent(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}

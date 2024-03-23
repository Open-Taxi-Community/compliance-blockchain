package compliance

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"compliance/testutil/sample"
	compliancesimulation "compliance/x/compliance/simulation"
	"compliance/x/compliance/types"
)

// avoid unused import issue
var (
	_ = compliancesimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateApproval = "op_weight_msg_create_approval"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateApproval int = 100

	opWeightMsgUpdateApproval = "op_weight_msg_update_approval"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateApproval int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	complianceGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&complianceGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateApproval int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateApproval, &weightMsgCreateApproval, nil,
		func(_ *rand.Rand) {
			weightMsgCreateApproval = defaultWeightMsgCreateApproval
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateApproval,
		compliancesimulation.SimulateMsgCreateApproval(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateApproval int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateApproval, &weightMsgUpdateApproval, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateApproval = defaultWeightMsgUpdateApproval
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateApproval,
		compliancesimulation.SimulateMsgUpdateApproval(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateApproval,
			defaultWeightMsgCreateApproval,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				compliancesimulation.SimulateMsgCreateApproval(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateApproval,
			defaultWeightMsgUpdateApproval,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				compliancesimulation.SimulateMsgUpdateApproval(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}

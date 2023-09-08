package eywacontract

import (
	"math/rand"

	"eywa-contract/testutil/sample"
	eywacontractsimulation "eywa-contract/x/eywacontract/simulation"
	"eywa-contract/x/eywacontract/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = eywacontractsimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgRegister = "op_weight_msg_register"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRegister int = 100

	opWeightMsgSendChat = "op_weight_msg_send_chat"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSendChat int = 100

	opWeightMsgHandshake = "op_weight_msg_handshake"
	// TODO: Determine the simulation weight value
	defaultWeightMsgHandshake int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	eywacontractGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&eywacontractGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgRegister int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRegister, &weightMsgRegister, nil,
		func(_ *rand.Rand) {
			weightMsgRegister = defaultWeightMsgRegister
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRegister,
		eywacontractsimulation.SimulateMsgRegister(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSendChat int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSendChat, &weightMsgSendChat, nil,
		func(_ *rand.Rand) {
			weightMsgSendChat = defaultWeightMsgSendChat
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSendChat,
		eywacontractsimulation.SimulateMsgSendChat(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgHandshake int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgHandshake, &weightMsgHandshake, nil,
		func(_ *rand.Rand) {
			weightMsgHandshake = defaultWeightMsgHandshake
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgHandshake,
		eywacontractsimulation.SimulateMsgHandshake(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgRegister,
			defaultWeightMsgRegister,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				eywacontractsimulation.SimulateMsgRegister(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSendChat,
			defaultWeightMsgSendChat,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				eywacontractsimulation.SimulateMsgSendChat(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgHandshake,
			defaultWeightMsgHandshake,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				eywacontractsimulation.SimulateMsgHandshake(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}

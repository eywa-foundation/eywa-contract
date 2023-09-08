package eywacontract

import (
	"eywa-contract/x/eywacontract/keeper"
	"eywa-contract/x/eywacontract/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the registerType
	for _, elem := range genState.RegisterTypeList {
		k.SetRegisterType(ctx, elem)
	}
	// Set all the sendChatType
	for _, elem := range genState.SendChatTypeList {
		k.SetSendChatType(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.RegisterTypeList = k.GetAllRegisterType(ctx)
	genesis.SendChatTypeList = k.GetAllSendChatType(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}

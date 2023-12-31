package eywacontract_test

import (
	"testing"

	keepertest "eywa-contract/testutil/keeper"
	"eywa-contract/testutil/nullify"
	"eywa-contract/x/eywacontract"
	"eywa-contract/x/eywacontract/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		RegisterTypeList: []types.RegisterType{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		SendChatTypeList: []types.SendChatType{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		HandshakeTypeList: []types.HandshakeType{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.EywacontractKeeper(t)
	eywacontract.InitGenesis(ctx, *k, genesisState)
	got := eywacontract.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.RegisterTypeList, got.RegisterTypeList)
	require.ElementsMatch(t, genesisState.SendChatTypeList, got.SendChatTypeList)
	require.ElementsMatch(t, genesisState.HandshakeTypeList, got.HandshakeTypeList)
	// this line is used by starport scaffolding # genesis/test/assert
}

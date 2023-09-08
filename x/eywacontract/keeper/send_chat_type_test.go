package keeper_test

import (
	"strconv"
	"testing"

	keepertest "eywa-contract/testutil/keeper"
	"eywa-contract/testutil/nullify"
	"eywa-contract/x/eywacontract/keeper"
	"eywa-contract/x/eywacontract/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNSendChatType(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.SendChatType {
	items := make([]types.SendChatType, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetSendChatType(ctx, items[i])
	}
	return items
}

func TestSendChatTypeGet(t *testing.T) {
	keeper, ctx := keepertest.EywacontractKeeper(t)
	items := createNSendChatType(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetSendChatType(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestSendChatTypeRemove(t *testing.T) {
	keeper, ctx := keepertest.EywacontractKeeper(t)
	items := createNSendChatType(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveSendChatType(ctx,
			item.Index,
		)
		_, found := keeper.GetSendChatType(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestSendChatTypeGetAll(t *testing.T) {
	keeper, ctx := keepertest.EywacontractKeeper(t)
	items := createNSendChatType(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllSendChatType(ctx)),
	)
}

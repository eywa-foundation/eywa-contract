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

func createNHandshakeType(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.HandshakeType {
	items := make([]types.HandshakeType, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetHandshakeType(ctx, items[i])
	}
	return items
}

func TestHandshakeTypeGet(t *testing.T) {
	keeper, ctx := keepertest.EywacontractKeeper(t)
	items := createNHandshakeType(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetHandshakeType(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestHandshakeTypeRemove(t *testing.T) {
	keeper, ctx := keepertest.EywacontractKeeper(t)
	items := createNHandshakeType(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveHandshakeType(ctx,
			item.Index,
		)
		_, found := keeper.GetHandshakeType(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestHandshakeTypeGetAll(t *testing.T) {
	keeper, ctx := keepertest.EywacontractKeeper(t)
	items := createNHandshakeType(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllHandshakeType(ctx)),
	)
}

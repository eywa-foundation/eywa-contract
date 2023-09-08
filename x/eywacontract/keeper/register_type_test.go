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

func createNRegisterType(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.RegisterType {
	items := make([]types.RegisterType, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetRegisterType(ctx, items[i])
	}
	return items
}

func TestRegisterTypeGet(t *testing.T) {
	keeper, ctx := keepertest.EywacontractKeeper(t)
	items := createNRegisterType(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetRegisterType(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestRegisterTypeRemove(t *testing.T) {
	keeper, ctx := keepertest.EywacontractKeeper(t)
	items := createNRegisterType(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveRegisterType(ctx,
			item.Index,
		)
		_, found := keeper.GetRegisterType(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestRegisterTypeGetAll(t *testing.T) {
	keeper, ctx := keepertest.EywacontractKeeper(t)
	items := createNRegisterType(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllRegisterType(ctx)),
	)
}

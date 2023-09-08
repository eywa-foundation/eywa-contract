package keeper_test

import (
	"testing"

	testkeeper "eywa-contract/testutil/keeper"
	"eywa-contract/x/eywacontract/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.EywacontractKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}

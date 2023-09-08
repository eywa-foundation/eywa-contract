package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "eywa-contract/testutil/keeper"
	"eywa-contract/testutil/nullify"
	"eywa-contract/x/eywacontract/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestRegisterTypeQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.EywacontractKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNRegisterType(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetRegisterTypeRequest
		response *types.QueryGetRegisterTypeResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetRegisterTypeRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetRegisterTypeResponse{RegisterType: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetRegisterTypeRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetRegisterTypeResponse{RegisterType: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetRegisterTypeRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.RegisterType(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestRegisterTypeQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.EywacontractKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNRegisterType(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllRegisterTypeRequest {
		return &types.QueryAllRegisterTypeRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.RegisterTypeAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.RegisterType), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.RegisterType),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.RegisterTypeAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.RegisterType), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.RegisterType),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.RegisterTypeAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.RegisterType),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.RegisterTypeAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}

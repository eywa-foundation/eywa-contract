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

func TestHandshakeTypeQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.EywacontractKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNHandshakeType(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetHandshakeTypeRequest
		response *types.QueryGetHandshakeTypeResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetHandshakeTypeRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetHandshakeTypeResponse{HandshakeType: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetHandshakeTypeRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetHandshakeTypeResponse{HandshakeType: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetHandshakeTypeRequest{
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
			response, err := keeper.HandshakeType(wctx, tc.request)
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

func TestHandshakeTypeQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.EywacontractKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNHandshakeType(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllHandshakeTypeRequest {
		return &types.QueryAllHandshakeTypeRequest{
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
			resp, err := keeper.HandshakeTypeAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.HandshakeType), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.HandshakeType),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.HandshakeTypeAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.HandshakeType), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.HandshakeType),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.HandshakeTypeAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.HandshakeType),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.HandshakeTypeAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}

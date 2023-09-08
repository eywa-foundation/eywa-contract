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

func TestSendChatTypeQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.EywacontractKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNSendChatType(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetSendChatTypeRequest
		response *types.QueryGetSendChatTypeResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetSendChatTypeRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetSendChatTypeResponse{SendChatType: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetSendChatTypeRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetSendChatTypeResponse{SendChatType: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetSendChatTypeRequest{
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
			response, err := keeper.SendChatType(wctx, tc.request)
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

func TestSendChatTypeQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.EywacontractKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNSendChatType(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllSendChatTypeRequest {
		return &types.QueryAllSendChatTypeRequest{
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
			resp, err := keeper.SendChatTypeAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.SendChatType), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.SendChatType),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.SendChatTypeAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.SendChatType), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.SendChatType),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.SendChatTypeAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.SendChatType),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.SendChatTypeAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}

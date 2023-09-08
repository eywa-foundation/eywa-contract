package keeper

import (
	"context"

	"eywa-contract/x/eywacontract/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) SendChatTypeAll(goCtx context.Context, req *types.QueryAllSendChatTypeRequest) (*types.QueryAllSendChatTypeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var sendChatTypes []types.SendChatType
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	sendChatTypeStore := prefix.NewStore(store, types.KeyPrefix(types.SendChatTypeKeyPrefix))

	pageRes, err := query.Paginate(sendChatTypeStore, req.Pagination, func(key []byte, value []byte) error {
		var sendChatType types.SendChatType
		if err := k.cdc.Unmarshal(value, &sendChatType); err != nil {
			return err
		}

		sendChatTypes = append(sendChatTypes, sendChatType)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSendChatTypeResponse{SendChatType: sendChatTypes, Pagination: pageRes}, nil
}

func (k Keeper) SendChatType(goCtx context.Context, req *types.QueryGetSendChatTypeRequest) (*types.QueryGetSendChatTypeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetSendChatType(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetSendChatTypeResponse{SendChatType: val}, nil
}

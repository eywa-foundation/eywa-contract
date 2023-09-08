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

func (k Keeper) HandshakeTypeAll(goCtx context.Context, req *types.QueryAllHandshakeTypeRequest) (*types.QueryAllHandshakeTypeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var handshakeTypes []types.HandshakeType
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	handshakeTypeStore := prefix.NewStore(store, types.KeyPrefix(types.HandshakeTypeKeyPrefix))

	pageRes, err := query.Paginate(handshakeTypeStore, req.Pagination, func(key []byte, value []byte) error {
		var handshakeType types.HandshakeType
		if err := k.cdc.Unmarshal(value, &handshakeType); err != nil {
			return err
		}

		handshakeTypes = append(handshakeTypes, handshakeType)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllHandshakeTypeResponse{HandshakeType: handshakeTypes, Pagination: pageRes}, nil
}

func (k Keeper) HandshakeType(goCtx context.Context, req *types.QueryGetHandshakeTypeRequest) (*types.QueryGetHandshakeTypeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetHandshakeType(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetHandshakeTypeResponse{HandshakeType: val}, nil
}

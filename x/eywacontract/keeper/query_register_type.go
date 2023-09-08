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

func (k Keeper) RegisterTypeAll(goCtx context.Context, req *types.QueryAllRegisterTypeRequest) (*types.QueryAllRegisterTypeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var registerTypes []types.RegisterType
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	registerTypeStore := prefix.NewStore(store, types.KeyPrefix(types.RegisterTypeKeyPrefix))

	pageRes, err := query.Paginate(registerTypeStore, req.Pagination, func(key []byte, value []byte) error {
		var registerType types.RegisterType
		if err := k.cdc.Unmarshal(value, &registerType); err != nil {
			return err
		}

		registerTypes = append(registerTypes, registerType)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRegisterTypeResponse{RegisterType: registerTypes, Pagination: pageRes}, nil
}

func (k Keeper) RegisterType(goCtx context.Context, req *types.QueryGetRegisterTypeRequest) (*types.QueryGetRegisterTypeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetRegisterType(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetRegisterTypeResponse{RegisterType: val}, nil
}

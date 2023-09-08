package keeper

import (
	"context"

	"eywa-contract/x/eywacontract/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Handshake(goCtx context.Context, msg *types.MsgHandshake) (*types.MsgHandshakeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgHandshakeResponse{}, nil
}

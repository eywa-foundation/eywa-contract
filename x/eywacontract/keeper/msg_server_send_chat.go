package keeper

import (
	"context"

	"eywa-contract/x/eywacontract/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SendChat(goCtx context.Context, msg *types.MsgSendChat) (*types.MsgSendChatResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSendChatResponse{}, nil
}

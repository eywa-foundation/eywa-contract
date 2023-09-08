package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"time"

	"eywa-contract/x/eywacontract/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SendChat(goCtx context.Context, msg *types.MsgSendChat) (*types.MsgSendChatResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	currentTime := time.Now().Local()
	var currentTimeBytes = []byte(currentTime.Format("2006-01-02 15:04:05.999"))
	var IndexBytes = append(currentTimeBytes, append([]byte(msg.Creator), []byte(msg.Receiver)...)...)
	var IndexHash = sha256.Sum256(IndexBytes)
	var IndexHashString = hex.EncodeToString(IndexHash[:])

	var chat = types.SendChatType{
		Index:    IndexHashString,
		Sender:   msg.Creator,
		Receiver: msg.Receiver,
		Chat:     msg.Chat,
	}

	k.SetSendChatType(ctx, chat)

	return &types.MsgSendChatResponse{}, nil
}

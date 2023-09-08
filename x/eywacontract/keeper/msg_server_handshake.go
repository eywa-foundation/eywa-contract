package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"time"

	"eywa-contract/x/eywacontract/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Handshake(goCtx context.Context, msg *types.MsgHandshake) (*types.MsgHandshakeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	currentTime := time.Now().Local()
	var currentTimeBytes = []byte(currentTime.Format("2006-01-02 15:04:05.999"))
	var IndexBytes = append(currentTimeBytes, append([]byte(msg.Creator), []byte(msg.Receiver)...)...)
	var IndexHash = sha256.Sum256(IndexBytes)
	var IndexHashString = hex.EncodeToString(IndexHash[:])

	var handshake = types.HandshakeType{
		Index:         IndexHashString,
		Sender:        msg.Creator,
		Receiver:      msg.Receiver,
		ServerAddress: msg.ServerAddress,
		Time:          msg.Time,
	}

	k.SetHandshakeType(ctx, handshake)

	return &types.MsgHandshakeResponse{}, nil
}

package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"time"

	"eywa-contract/x/eywacontract/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Register(goCtx context.Context, msg *types.MsgRegister) (*types.MsgRegisterResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	currentTime := time.Now().Local()
	var currentTimeBytes = []byte(currentTime.Format("2006-01-02"))
	var currentTimeHash = sha256.Sum256(currentTimeBytes)
	var currentTimeHashString = hex.EncodeToString(currentTimeHash[:])

	var register_type = types.RegisterType{
		Index:     currentTimeHashString,
		Submitter: msg.Creator,
		Pubkey:    msg.Pubkey,
	}

	// TODO: if the register_type already exists, return an error ??

	_, isFound := k.GetRegisterType(ctx, register_type.Index)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "The register_type already exists")
	}

	k.SetRegisterType(ctx, register_type)

	// TODO: Handling the message

	return &types.MsgRegisterResponse{}, nil
}

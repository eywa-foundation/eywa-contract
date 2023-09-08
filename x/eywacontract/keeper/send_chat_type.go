package keeper

import (
	"eywa-contract/x/eywacontract/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetSendChatType set a specific sendChatType in the store from its index
func (k Keeper) SetSendChatType(ctx sdk.Context, sendChatType types.SendChatType) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SendChatTypeKeyPrefix))
	b := k.cdc.MustMarshal(&sendChatType)
	store.Set(types.SendChatTypeKey(
		sendChatType.Index,
	), b)
}

// GetSendChatType returns a sendChatType from its index
func (k Keeper) GetSendChatType(
	ctx sdk.Context,
	index string,

) (val types.SendChatType, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SendChatTypeKeyPrefix))

	b := store.Get(types.SendChatTypeKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSendChatType removes a sendChatType from the store
func (k Keeper) RemoveSendChatType(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SendChatTypeKeyPrefix))
	store.Delete(types.SendChatTypeKey(
		index,
	))
}

// GetAllSendChatType returns all sendChatType
func (k Keeper) GetAllSendChatType(ctx sdk.Context) (list []types.SendChatType) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SendChatTypeKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.SendChatType
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

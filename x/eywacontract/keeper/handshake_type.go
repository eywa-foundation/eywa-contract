package keeper

import (
	"eywa-contract/x/eywacontract/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetHandshakeType set a specific handshakeType in the store from its index
func (k Keeper) SetHandshakeType(ctx sdk.Context, handshakeType types.HandshakeType) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HandshakeTypeKeyPrefix))
	b := k.cdc.MustMarshal(&handshakeType)
	store.Set(types.HandshakeTypeKey(
		handshakeType.Index,
	), b)
}

// GetHandshakeType returns a handshakeType from its index
func (k Keeper) GetHandshakeType(
	ctx sdk.Context,
	index string,

) (val types.HandshakeType, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HandshakeTypeKeyPrefix))

	b := store.Get(types.HandshakeTypeKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveHandshakeType removes a handshakeType from the store
func (k Keeper) RemoveHandshakeType(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HandshakeTypeKeyPrefix))
	store.Delete(types.HandshakeTypeKey(
		index,
	))
}

// GetAllHandshakeType returns all handshakeType
func (k Keeper) GetAllHandshakeType(ctx sdk.Context) (list []types.HandshakeType) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HandshakeTypeKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.HandshakeType
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

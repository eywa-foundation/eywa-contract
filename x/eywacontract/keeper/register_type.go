package keeper

import (
	"eywa-contract/x/eywacontract/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetRegisterType set a specific registerType in the store from its index
func (k Keeper) SetRegisterType(ctx sdk.Context, registerType types.RegisterType) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RegisterTypeKeyPrefix))
	b := k.cdc.MustMarshal(&registerType)
	store.Set(types.RegisterTypeKey(
		registerType.Index,
	), b)
}

// GetRegisterType returns a registerType from its index
func (k Keeper) GetRegisterType(
	ctx sdk.Context,
	index string,

) (val types.RegisterType, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RegisterTypeKeyPrefix))

	b := store.Get(types.RegisterTypeKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveRegisterType removes a registerType from the store
func (k Keeper) RemoveRegisterType(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RegisterTypeKeyPrefix))
	store.Delete(types.RegisterTypeKey(
		index,
	))
}

// GetAllRegisterType returns all registerType
func (k Keeper) GetAllRegisterType(ctx sdk.Context) (list []types.RegisterType) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RegisterTypeKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.RegisterType
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

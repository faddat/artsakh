package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/faddat/artsakh/x/artsakh/types"
    "github.com/cosmos/cosmos-sdk/codec"
)

// CreateAdd creates a add
func (k Keeper) CreateAdd(ctx sdk.Context, add types.Add) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.AddPrefix + add.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(add)
	store.Set(key, value)
}

// GetAdd returns the add information
func (k Keeper) GetAdd(ctx sdk.Context, key string) (types.Add, error) {
	store := ctx.KVStore(k.storeKey)
	var add types.Add
	byteKey := []byte(types.AddPrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &add)
	if err != nil {
		return add, err
	}
	return add, nil
}

// SetAdd sets a add
func (k Keeper) SetAdd(ctx sdk.Context, add types.Add) {
	addKey := add.ID
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(add)
	key := []byte(types.AddPrefix + addKey)
	store.Set(key, bz)
}

// DeleteAdd deletes a add
func (k Keeper) DeleteAdd(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.AddPrefix + key))
}

//
// Functions used by querier
//

func listAdd(ctx sdk.Context, k Keeper) ([]byte, error) {
	var addList []types.Add
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.AddPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var add types.Add
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &add)
		addList = append(addList, add)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, addList)
	return res, nil
}

func getAdd(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	add, err := k.GetAdd(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, add)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// Get creator of the item
func (k Keeper) GetAddOwner(ctx sdk.Context, key string) sdk.AccAddress {
	add, err := k.GetAdd(ctx, key)
	if err != nil {
		return nil
	}
	return add.Creator
}


// Check if the key exists in the store
func (k Keeper) AddExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.AddPrefix + key))
}

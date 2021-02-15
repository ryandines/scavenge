package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"

	"github.com/starport/scavenge/x/scavenge/types"
    "github.com/cosmos/cosmos-sdk/codec"
)

// GetScavengeCount get the total number of scavenge
func (k Keeper) GetScavengeCount(ctx sdk.Context) int64 {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.ScavengeCountPrefix)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseInt(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to int64
		panic("cannot decode count")
	}

	return count
}

// SetScavengeCount set the total number of scavenge
func (k Keeper) SetScavengeCount(ctx sdk.Context, count int64)  {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.ScavengeCountPrefix)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

// CreateScavenge creates a scavenge
func (k Keeper) CreateScavenge(ctx sdk.Context, msg types.MsgCreateScavenge) {
	// Create the scavenge
	count := k.GetScavengeCount(ctx)
    var scavenge = types.Scavenge{
        Creator: msg.Creator,
        ID:      strconv.FormatInt(count, 10),
        Description: msg.Description,
        SolutionHash: msg.SolutionHash,
        Reward: msg.Reward,
        Solution: msg.Solution,
        Scavenger: msg.Scavenger,
    }

	store := ctx.KVStore(k.storeKey)
	key := []byte(types.ScavengePrefix + scavenge.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(scavenge)
	store.Set(key, value)

	// Update scavenge count
    k.SetScavengeCount(ctx, count+1)
}

// GetScavenge returns the scavenge information
func (k Keeper) GetScavenge(ctx sdk.Context, key string) (types.Scavenge, error) {
	store := ctx.KVStore(k.storeKey)
	var scavenge types.Scavenge
	byteKey := []byte(types.ScavengePrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &scavenge)
	if err != nil {
		return scavenge, err
	}
	return scavenge, nil
}

// SetScavenge sets a scavenge
func (k Keeper) SetScavenge(ctx sdk.Context, scavenge types.Scavenge) {
	scavengeKey := scavenge.ID
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(scavenge)
	key := []byte(types.ScavengePrefix + scavengeKey)
	store.Set(key, bz)
}

// DeleteScavenge deletes a scavenge
func (k Keeper) DeleteScavenge(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.ScavengePrefix + key))
}

//
// Functions used by querier
//

func listScavenge(ctx sdk.Context, k Keeper) ([]byte, error) {
	var scavengeList []types.Scavenge
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.ScavengePrefix))
	for ; iterator.Valid(); iterator.Next() {
		var scavenge types.Scavenge
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &scavenge)
		scavengeList = append(scavengeList, scavenge)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, scavengeList)
	return res, nil
}

func getScavenge(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	scavenge, err := k.GetScavenge(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, scavenge)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// Get creator of the item
func (k Keeper) GetScavengeOwner(ctx sdk.Context, key string) sdk.AccAddress {
	scavenge, err := k.GetScavenge(ctx, key)
	if err != nil {
		return nil
	}
	return scavenge.Creator
}


// Check if the key exists in the store
func (k Keeper) ScavengeExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.ScavengePrefix + key))
}

package keeper

import (
	"agent/x/agent/types"
	"cosmossdk.io/store/prefix"
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AppendPrompt(ctx sdk.Context, prompt types.Prompt) uint64 {
	count := k.GetPromptCount(ctx)
	prompt.Id = count
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PromptKey))
	appendedValue := k.cdc.MustMarshal(&prompt)
	store.Set(GetPromptIDBytes(prompt.Id), appendedValue)
	k.SetPromptCount(ctx, count+1)
	return count
}

func (k Keeper) GetPromptCount(ctx sdk.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.PromptCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

func GetPromptIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

func (k Keeper) SetPromptCount(ctx sdk.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.PromptCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

func (k Keeper) GetPromptText(ctx sdk.Context, id uint64) (val types.Prompt, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PromptKey))
	b := store.Get(GetQuestionIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

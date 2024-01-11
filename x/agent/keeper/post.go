package keeper

import (
	"agent/x/agent/types"
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AppendPost(ctx sdk.Context, question types.Aiquery) uint64 {
	count := k.GetQuestionCount(ctx)
	question.Id = count
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.QuestionKey))
	appendedValue := k.cdc.MustMarshal(&question)
	store.Set(GetQuestionIDBytes(question.Id), appendedValue)
	k.SetQuestionCount(ctx, count+1)
	return count
}

func (k Keeper) GetQuestionCount(ctx sdk.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.QuestionCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

func GetQuestionIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

func (k Keeper) SetQuestionCount(ctx sdk.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.QuestionCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

func (k Keeper) GetAiquery(ctx sdk.Context, id uint64) (val types.Aiquery, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.QuestionKey))
	b := store.Get(GetQuestionIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) SetAnswer(ctx sdk.Context, aiquery types.Aiquery) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.QuestionKey))
	b := k.cdc.MustMarshal(&aiquery)
	store.Set(GetQuestionIDBytes(aiquery.Id), b)
}

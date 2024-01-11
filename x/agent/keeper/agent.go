package keeper

import (
	"agent/x/agent/types"
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AppendAgentRun(ctx sdk.Context, question types.Agentrun) uint64 {
	count := k.GetAgentRunCount(ctx)
	question.Id = count
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AgentrunKey))
	appendedValue := k.cdc.MustMarshal(&question)
	store.Set(GetAgentrunIDBytes(question.Id), appendedValue)
	k.SetAgentrunCount(ctx, count+1)
	return count
}

func (k Keeper) GetAgentRunCount(ctx sdk.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.AgentrunKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

func GetAgentrunIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

func (k Keeper) SetAgentrunCount(ctx sdk.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.AgentrunKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

func (k Keeper) GetAgentrun(ctx sdk.Context, id uint64) (val types.Agentrun, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AgentrunKey))
	b := store.Get(GetAgentrunIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// TODO: write logic
func (k Keeper) AddResponse(ctx sdk.Context, aiquery types.Agentrun) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AgentrunKey))
	b := k.cdc.MustMarshal(&aiquery)
	store.Set(GetAgentrunIDBytes(aiquery.Id), b)
}

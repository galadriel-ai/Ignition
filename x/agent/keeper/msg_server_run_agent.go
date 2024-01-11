package keeper

import (
	"context"

	"agent/x/agent/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RunAgent(goCtx context.Context, msg *types.MsgRunAgent) (*types.MsgRunAgentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var question = types.Agentrun{
		Query:      msg.Text,
		Isfinished: false,
		Prompt:     msg.Promptid,
	}
	id := k.AppendAgentRun(
		ctx,
		question,
	)
	return &types.MsgRunAgentResponse{
		Id: id,
	}, nil
}

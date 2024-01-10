package keeper

import (
	"context"

	"agent/x/agent/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AskQuestion(goCtx context.Context, msg *types.MsgAskQuestion) (*types.MsgAskQuestionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var question = types.Aiquery{
		Question: msg.Question,
	}
	id := k.AppendPost(
		ctx,
		question,
	)
	return &types.MsgAskQuestionResponse{
		Id: id,
	}, nil
}

package keeper

import (
	"context"

	"agent/x/agent/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AnswerQuestion(goCtx context.Context, msg *types.MsgAnswerQuestion) (*types.MsgAnswerQuestionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgAnswerQuestionResponse{}, nil
}

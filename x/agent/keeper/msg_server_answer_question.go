package keeper

import (
	"context"
	errorsmod "cosmossdk.io/errors"
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"agent/x/agent/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AnswerQuestion(goCtx context.Context, msg *types.MsgAnswerQuestion) (*types.MsgAnswerQuestionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	val, found := k.GetAiquery(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	var post = types.Aiquery{
		Id:       msg.Id,
		Question: val.Question,
		Answer:   msg.Answer,
	}
	//if msg.Creator != val.Creator {
	//	return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	//}
	k.SetAnswer(ctx, post)
	return &types.MsgAnswerQuestionResponse{}, nil
}

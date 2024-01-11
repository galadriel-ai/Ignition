package keeper

import (
	"context"

	"agent/x/agent/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AddPrompt(goCtx context.Context, msg *types.MsgAddPrompt) (*types.MsgAddPromptResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var prompt = types.Prompt{
		Text: msg.Text,
	}
	id := k.AppendPrompt(
		ctx,
		prompt,
	)
	return &types.MsgAddPromptResponse{
		Id: id,
	}, nil
}

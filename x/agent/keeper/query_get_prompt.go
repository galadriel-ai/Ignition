package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"agent/x/agent/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetPrompt(goCtx context.Context, req *types.QueryGetPromptRequest) (*types.QueryGetPromptResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	prompt, found := k.GetPromptText(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetPromptResponse{Text: prompt.Text}, nil
}

package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"agent/x/agent/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ShowQuestion(goCtx context.Context, req *types.QueryShowQuestionRequest) (*types.QueryShowQuestionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	question, found := k.GetAiquery(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryShowQuestionResponse{Aiquery: &question}, nil
	return &types.QueryShowQuestionResponse{}, nil
}

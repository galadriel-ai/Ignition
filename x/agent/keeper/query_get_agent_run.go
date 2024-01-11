package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"agent/x/agent/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetAgentRun(goCtx context.Context, req *types.QueryGetAgentRunRequest) (*types.QueryGetAgentRunResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	run, found := k.GetAgentrun(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	formattedRun := types.Agentrun{
		Query:           run.Query,
		Responses:       run.Responses,
		Functioncalls:   run.Functioncalls,
		Functionresults: run.Functionresults,
		Isfinished:      run.Isfinished,
		Id:              run.Id,
		Prompt:          run.Prompt,
	}
	return &types.QueryGetAgentRunResponse{Run: &formattedRun}, nil
}

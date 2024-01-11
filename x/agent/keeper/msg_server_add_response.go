package keeper

import (
	"context"
	errorsmod "cosmossdk.io/errors"
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"
	"strings"

	"agent/x/agent/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AddResponse(goCtx context.Context, msg *types.MsgAddResponse) (*types.MsgAddResponseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	val, found := k.GetAgentrun(ctx, msg.Runid)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Runid))
	}
	// Cannot un-finish run
	isFinished := val.Isfinished
	if !isFinished {
		isFinished = msg.Isfinished
	}

	functionCalls := val.Functioncalls
	newResponses, newFunctionCalls := k.GetInputs(msg.Response, val.Responses, functionCalls)
	functionResults := k.GetFunctionResults(val.Functionresults, functionCalls, newFunctionCalls)

	var post = types.Agentrun{
		Id:              val.Id,
		Query:           val.Query,
		Responses:       newResponses,
		Functioncalls:   newFunctionCalls,
		Functionresults: functionResults,
		Isfinished:      isFinished,
	}
	//if msg.Creator != val.Creator {
	//	return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	//}
	k.UpdateAgentRun(ctx, post)
	return &types.MsgAddResponseResponse{}, nil
}

// Returns responses and function calls, includes (k Keeper) for unit testing..
func (k Keeper) GetInputs(
	newResponse string,
	responses []string,
	functionCalls []string,
) ([]string, []string) {
	newFunctionCalls := functionCalls
	if strings.Contains(newResponse, "F_CALL: ") {
		split := strings.SplitAfter(newResponse, "F_CALL: ")
		if len(split) > 1 {
			newFunctionCalls = append(functionCalls, split[1])
		}
	} else {
		responses = append(responses, newResponse)
	}
	return responses, newFunctionCalls
}

func (k Keeper) GetFunctionResults(existingFunctionResults []string, functionCalls []string, newFunctionCalls []string) []string {
	newFunctionResults := existingFunctionResults
	if len(newFunctionCalls) > len(functionCalls) {
		newFunctionResult := k.CallFunction(newFunctionCalls[len(newFunctionCalls)-1])
		newFunctionResults = append(newFunctionResults, newFunctionResult)
	}
	return newFunctionResults
}

func (k Keeper) CallFunction(
	functionCall string,
) string {
	// TODO: error handling :)
	functionName := strings.Split(functionCall, "(")[0]
	functionParam := strings.Split(strings.Split(functionCall, "(\"")[1], "\")")[0]
	if functionName == "MultiplyBy8" {
		return k.MultiplyBy8(functionParam)
	}
	if functionName == "DivideBy2" {
		return k.DivideBy2(functionParam)
	}
	return ""
}

func (k Keeper) MultiplyBy8(input string) string {
	parseInt, _ := strconv.ParseInt(input, 10, 64)
	return strconv.FormatInt(parseInt*8, 10)
}

func (k Keeper) DivideBy2(input string) string {
	parseInt, _ := strconv.ParseInt(input, 10, 64)
	return strconv.FormatInt(parseInt/2, 10)
}

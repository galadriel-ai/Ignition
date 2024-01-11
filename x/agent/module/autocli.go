package agent

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "agent/api/agent/agent"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod:      "ShowQuestion",
					Use:            "show-question [id]",
					Short:          "Query show-question",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},

				{
					RpcMethod:      "GetPrompt",
					Use:            "get-prompt [id]",
					Short:          "Query get-prompt",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},

				{
					RpcMethod:      "GetAgentRun",
					Use:            "get-agent-run [id]",
					Short:          "Query get-agent-run",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},

				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "AskQuestion",
					Use:            "ask-question [question]",
					Short:          "Send a ask-question tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "question"}},
				},
				{
					RpcMethod:      "AnswerQuestion",
					Use:            "answer-question [answer] [id]",
					Short:          "Send a answer-question tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "answer"}, {ProtoField: "id"}},
				},
				{
					RpcMethod:      "AddPrompt",
					Use:            "add-prompt [text]",
					Short:          "Send a add-prompt tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "text"}},
				},
				{
					RpcMethod:      "RunAgent",
					Use:            "run-agent [text] [promptid]",
					Short:          "Send a run-agent tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "text"}, {ProtoField: "promptid"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}

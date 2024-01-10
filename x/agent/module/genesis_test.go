package agent_test

import (
	"testing"

	keepertest "agent/testutil/keeper"
	"agent/testutil/nullify"
	"agent/x/agent/module"
	"agent/x/agent/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AgentKeeper(t)
	agent.InitGenesis(ctx, k, genesisState)
	got := agent.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

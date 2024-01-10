package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "agent/testutil/keeper"
	"agent/x/agent/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.AgentKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}

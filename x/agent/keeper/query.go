package keeper

import (
	"agent/x/agent/types"
)

var _ types.QueryServer = Keeper{}

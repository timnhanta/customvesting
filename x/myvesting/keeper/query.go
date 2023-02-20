package keeper

import (
	"customvesting/x/myvesting/types"
)

var _ types.QueryServer = Keeper{}

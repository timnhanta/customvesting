package keeper

import (
	"github.com/timnhanta/customvesting/x/myvesting/types"
)

var _ types.QueryServer = Keeper{}

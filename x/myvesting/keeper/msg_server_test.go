package keeper_test

import (
	"context"
	"testing"

	keepertest "customvesting/testutil/keeper"
	"customvesting/x/myvesting/keeper"
	"customvesting/x/myvesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.MyvestingKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

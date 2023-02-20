package keeper_test

import (
	"testing"

	testkeeper "customvesting/testutil/keeper"
	"customvesting/x/myvesting/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.MyvestingKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}

package myvesting_test

import (
	"testing"

	keepertest "customvesting/testutil/keeper"
	"customvesting/testutil/nullify"
	"customvesting/x/myvesting"
	"customvesting/x/myvesting/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MyvestingKeeper(t)
	myvesting.InitGenesis(ctx, *k, genesisState)
	got := myvesting.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

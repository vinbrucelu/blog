package blog_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/vinbrucelu/blog/testutil/keeper"
	"github.com/vinbrucelu/blog/testutil/nullify"
	"github.com/vinbrucelu/blog/x/blog"
	"github.com/vinbrucelu/blog/x/blog/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BlogKeeper(t)
	blog.InitGenesis(ctx, *k, genesisState)
	got := blog.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

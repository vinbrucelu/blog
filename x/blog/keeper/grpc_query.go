package keeper

import (
	"github.com/vinbrucelu/blog/x/blog/types"
)

var _ types.QueryServer = Keeper{}

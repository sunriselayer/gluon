package keeper

import (
	"gluon/x/contract/types"
)

var _ types.QueryServer = Keeper{}

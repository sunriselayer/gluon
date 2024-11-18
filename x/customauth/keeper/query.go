package keeper

import (
	"gluon/x/customauth/types"
)

var _ types.QueryServer = Keeper{}

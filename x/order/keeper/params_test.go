package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

    keepertest "gluon/testutil/keeper"
    "gluon/x/order/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.OrderKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}

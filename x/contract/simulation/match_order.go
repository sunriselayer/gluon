package simulation

import (
	"math/rand"

	"gluon/x/contract/keeper"
	"gluon/x/contract/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgMatchOrder(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgMatchOrder{
			Earlier: types.Order{
				Address: simAccount.Address.String(),
			},
			Later: types.Order{
				Address: simAccount.Address.String(),
			},
		}

		// TODO: Handling the MatchOrder simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "MatchOrder simulation not implemented"), nil, nil
	}
}

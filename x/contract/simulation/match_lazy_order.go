package simulation

import (
	"math/rand"

	"gluon/x/contract/keeper"
	"gluon/x/contract/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgMatchLazyOrder(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgMatchLazyOrder{
			EarlierAddress: simAccount.Address.String(),
			LaterAddress:   simAccount.Address.String(),
		}

		// TODO: Handling the MatchLazyOrder simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "MatchLazyOrder simulation not implemented"), nil, nil
	}
}

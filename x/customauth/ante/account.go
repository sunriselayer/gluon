package ante

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
)

// SetAccountDecorator creates a new BaseAccount, if the signer is not exist on chain.
// This process is necessary when an account without balance sends tx with operator's signature.
// On Gluon, it is possible to send tx with the user's sequence and operator's signature without any fee.
type SetAccountDecorator struct {
	accountKeeper AccountKeeper
}

func NewSetAccountDecorator(ak AccountKeeper) SetAccountDecorator {
	return SetAccountDecorator{
		accountKeeper: ak,
	}
}

func (dfd SetAccountDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	sigTx, ok := tx.(authsigning.SigVerifiableTx)
	if !ok {
		return ctx, errorsmod.Wrap(sdkerrors.ErrTxDecode, "invalid transaction type")
	}

	signers, err := sigTx.GetSigners()
	if err != nil {
		return sdk.Context{}, err
	}

	for _, signer := range signers {
		accExists := dfd.accountKeeper.HasAccount(ctx, signer)

		if !accExists {
			acc := dfd.accountKeeper.NewAccountWithAddress(ctx, signer)
			err = acc.SetAccountNumber(0)
			if err != nil {
				return sdk.Context{}, err
			}
			dfd.accountKeeper.SetAccount(ctx, acc)
		}
	}

	return next(ctx, tx, simulate)
}

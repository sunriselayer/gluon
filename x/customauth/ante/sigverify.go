package ante

import (
	txsigning "cosmossdk.io/x/tx/signing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	sdkante "github.com/cosmos/cosmos-sdk/x/auth/ante"
)

// SigVerificationDecorator verifies all signatures for a tx and return an error if any are invalid. Note,
// the SigVerificationDecorator will not check signatures on ReCheck.
//
// CONTRACT: Pubkeys are set in context for all signers before this decorator runs
// CONTRACT: Tx must implement SigVerifiableTx interface
type SigVerificationDecorator struct {
	ak              AccountKeeper
	signModeHandler *txsigning.HandlerMap

	cak       CustomAuthKeeper
	customMap CustomSigVerifierMap
}

func NewSigVerificationDecorator(
	ak AccountKeeper,
	signModeHandler *txsigning.HandlerMap,
	cak CustomAuthKeeper,
	customMap CustomSigVerifierMap,
) SigVerificationDecorator {
	return SigVerificationDecorator{
		ak:              ak,
		signModeHandler: signModeHandler,

		cak:       cak,
		customMap: customMap,
	}
}

func (svd SigVerificationDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	if !svd.customMap.Match(tx) {
		return sdkante.NewSigVerificationDecorator(svd.ak, svd.signModeHandler).AnteHandle(ctx, tx, simulate, next)
	}

	return next(ctx, tx, simulate)
}

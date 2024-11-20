package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	sdkante "github.com/cosmos/cosmos-sdk/x/auth/ante"
)

// DeductFeeDecorator deducts fees from the fee payer. The fee payer is the fee granter (if specified) or first signer of the tx.
// If the fee payer does not have the funds to pay for the fees, return an InsufficientFunds error.
// Call next AnteHandler if fees successfully deducted.
// CONTRACT: Tx must implement FeeTx interface to use DeductFeeDecorator
type DeductFeeDecorator struct {
	accountKeeper  AccountKeeper
	bankKeeper     BankKeeper
	feegrantKeeper FeegrantKeeper
	tfc            sdkante.TxFeeChecker

	customMap CustomSigVerifierMap
}

func NewDeductFeeDecorator(ak AccountKeeper, bk BankKeeper, fk FeegrantKeeper, tfc sdkante.TxFeeChecker, customMap CustomSigVerifierMap) DeductFeeDecorator {
	return DeductFeeDecorator{
		accountKeeper:  ak,
		bankKeeper:     bk,
		feegrantKeeper: fk,
		tfc:            tfc,

		customMap: customMap,
	}
}

func (dfd DeductFeeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	if !dfd.customMap.Match(tx) {
		sdkante.NewDeductFeeDecorator(dfd.accountKeeper, dfd.bankKeeper, dfd.feegrantKeeper, dfd.tfc).AnteHandle(ctx, tx, simulate, next)
	}

	return next(ctx, tx, simulate)
}

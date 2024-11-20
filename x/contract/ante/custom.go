package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	customante "gluon/x/customauth/ante"
)

func MsgMatchOrderSigVerifier(cak customante.CustomAuthKeeper, tx sdk.Tx) error {

	return nil
}

func MsgMatchLazyOrderSigVerifier(cak customante.CustomAuthKeeper, tx sdk.Tx) error {
	return nil
}

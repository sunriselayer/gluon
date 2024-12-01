package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/perp module sentinel errors
var (
	ErrInvalidSigner    = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrInvalidOrderType = sdkerrors.Register(ModuleName, 1101, "must be perp order")
)

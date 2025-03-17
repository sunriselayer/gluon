package types

// DONTCOVER

import (
	errors "cosmossdk.io/errors/v2"
)

// x/perp module sentinel errors
var (
	ErrInvalidSigner              = errors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrInvalidOrderType           = errors.Register(ModuleName, 1101, "must be perp order")
	ErrInvalidMargin              = errors.Register(ModuleName, 1102, "invalid margin")
	ErrPositionCancelAmountExceed = errors.Register(ModuleName, 1113, "position cancel amount exceed")
)

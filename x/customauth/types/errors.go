package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/customauth module sentinel errors
var (
	ErrInvalidSigner      = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrPairingNotFound    = sdkerrors.Register(ModuleName, 1101, "pairing not found")
	ErrPairingDelayPeriod = sdkerrors.Register(ModuleName, 1102, "pairing key can't be used during delay period")
)

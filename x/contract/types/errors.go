package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/contract module sentinel errors
var (
	ErrInvalidSigner        = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrNotPositiveAmount    = sdkerrors.Register(ModuleName, 1101, "amount must be positive")
	ErrEmptySignature       = sdkerrors.Register(ModuleName, 1102, "signature must not be empty")
	ErrInvalidSignature     = sdkerrors.Register(ModuleName, 1103, "invalid signature")
	ErrPairingNotFound      = sdkerrors.Register(ModuleName, 1104, "pairing not found")
	ErrInvalidPacketTimeout = sdkerrors.Register(ModuleName, 1500, "invalid packet timeout")
	ErrInvalidVersion       = sdkerrors.Register(ModuleName, 1501, "invalid version")
)

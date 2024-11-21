package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/contract module sentinel errors
var (
	ErrInvalidSigner         = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrNotPositiveAmount     = sdkerrors.Register(ModuleName, 1101, "amount must be positive")
	ErrEmptyOrderId          = sdkerrors.Register(ModuleName, 1102, "order id must not be empty")
	ErrEmptySignature        = sdkerrors.Register(ModuleName, 1103, "signature must not be empty")
	ErrInvalidSignature      = sdkerrors.Register(ModuleName, 1104, "invalid signature")
	ErrPairingNotFound       = sdkerrors.Register(ModuleName, 1105, "pairing not found")
	ErrDenomMismatch         = sdkerrors.Register(ModuleName, 1106, "denom mismatch")
	ErrUnknownOrderDirection = sdkerrors.Register(ModuleName, 1107, "unknown order direction")
	ErrSameOrderDirection    = sdkerrors.Register(ModuleName, 1108, "same order direction")
	ErrBothMarketPriceOrder  = sdkerrors.Register(ModuleName, 1109, "both market price order")
	ErrPriceMismatch         = sdkerrors.Register(ModuleName, 1110, "price mismatch")
	ErrInvalidPacketTimeout  = sdkerrors.Register(ModuleName, 1500, "invalid packet timeout")
	ErrInvalidVersion        = sdkerrors.Register(ModuleName, 1501, "invalid version")
)

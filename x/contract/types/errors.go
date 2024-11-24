package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/contract module sentinel errors
var (
	ErrInvalidSigner                 = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrNotPositiveAmount             = sdkerrors.Register(ModuleName, 1101, "amount must be positive")
	ErrEmptyOrderId                  = sdkerrors.Register(ModuleName, 1102, "order id must not be empty")
	ErrEmptySignature                = sdkerrors.Register(ModuleName, 1103, "signature must not be empty")
	ErrInvalidSignature              = sdkerrors.Register(ModuleName, 1104, "invalid signature")
	ErrPairingNotFound               = sdkerrors.Register(ModuleName, 1105, "pairing not found")
	ErrOrderNotFound                 = sdkerrors.Register(ModuleName, 1106, "order not found")
	ErrSameAddress                   = sdkerrors.Register(ModuleName, 1106, "same address")
	ErrDenomMismatch                 = sdkerrors.Register(ModuleName, 1108, "denom mismatch")
	ErrInvalidOrderDirection         = sdkerrors.Register(ModuleName, 1109, "invalid order direction")
	ErrBothMarketPriceOrder          = sdkerrors.Register(ModuleName, 1110, "both market price order")
	ErrPriceMismatch                 = sdkerrors.Register(ModuleName, 1111, "price mismatch")
	ErrOrderExpired                  = sdkerrors.Register(ModuleName, 1112, "order expired")
	ErrContractAmountExceed          = sdkerrors.Register(ModuleName, 1113, "contract amount exceed")
	ErrLazyContractNotAllowed        = sdkerrors.Register(ModuleName, 1114, "lazy contract not allowed")
	ErrOrderExpireBeforeLazyContract = sdkerrors.Register(ModuleName, 1115, "order expire before lazy contract")
	ErrNotBuyerNeitherSeller         = sdkerrors.Register(ModuleName, 1116, "not buyer neither seller")
	ErrInvalidPacketTimeout          = sdkerrors.Register(ModuleName, 1500, "invalid packet timeout")
	ErrInvalidVersion                = sdkerrors.Register(ModuleName, 1501, "invalid version")
)

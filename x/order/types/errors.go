package types

// DONTCOVER

import (
	errors "cosmossdk.io/errors/v2"
)

// x/order module sentinel errors
var (
	ErrInvalidSigner                 = errors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrNotPositive                   = errors.Register(ModuleName, 1101, "value must be positive")
	ErrEmptyOrderHash                = errors.Register(ModuleName, 1102, "order hash must not be empty")
	ErrEmptySignature                = errors.Register(ModuleName, 1103, "signature must not be empty")
	ErrInvalidSignature              = errors.Register(ModuleName, 1104, "invalid signature")
	ErrPairingNotFound               = errors.Register(ModuleName, 1105, "pairing not found")
	ErrOrderNotFound                 = errors.Register(ModuleName, 1106, "order not found")
	ErrSameAddress                   = errors.Register(ModuleName, 1107, "same address")
	ErrDenomMismatch                 = errors.Register(ModuleName, 1108, "denom mismatch")
	ErrInvalidOrderDirection         = errors.Register(ModuleName, 1109, "invalid order direction")
	ErrBothMarketPriceOrder          = errors.Register(ModuleName, 1110, "both market price order")
	ErrPriceMismatch                 = errors.Register(ModuleName, 1111, "price mismatch")
	ErrOrderExpired                  = errors.Register(ModuleName, 1112, "order expired")
	ErrContractAmountExceed          = errors.Register(ModuleName, 1113, "contract amount exceed")
	ErrLazyContractNotAllowed        = errors.Register(ModuleName, 1114, "lazy contract not allowed")
	ErrOrderExpireBeforeLazyContract = errors.Register(ModuleName, 1115, "order expire before lazy contract")
	ErrNotBuyerNeitherSeller         = errors.Register(ModuleName, 1116, "not buyer neither seller")
	ErrInvalidPacketTimeout          = errors.Register(ModuleName, 1500, "invalid packet timeout")
	ErrInvalidVersion                = errors.Register(ModuleName, 1501, "invalid version")
)

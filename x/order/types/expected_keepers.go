package types

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"

	customauthtypes "gluon/x/customauth/types"
)

// AccountKeeper defines the expected interface for the Account module.
type AccountKeeper interface {
	GetAccount(context.Context, sdk.AccAddress) sdk.AccountI // only used for simulation
	// Methods imported from account should be defined here
	GetModuleAddress(moduleName string) sdk.AccAddress
}

// BankKeeper defines the expected interface for the Bank module.
type BankKeeper interface {
	SpendableCoins(context.Context, sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
}

type CustomAuthKeeper interface {
	GetPairing(ctx context.Context, address string, id uint64) (val customauthtypes.Pairing, found bool)
	GetPairingPubKey(ctx context.Context, pairing customauthtypes.Pairing) (cryptotypes.PubKey, error)
}
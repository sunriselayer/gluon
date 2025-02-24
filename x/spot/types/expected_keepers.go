package types

import (
	"context"

	sdkmath "cosmossdk.io/math"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	customauthtypes "gluon/x/customauth/types"
	"gluon/x/customauth/types/pairing"

	ordertypes "gluon/x/order/types"
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
	SendCoins(ctx context.Context, fromAddr, toAddr sdk.AccAddress, amt sdk.Coins) error
}

type CustomAuthKeeper interface {
	GetPairing(ctx context.Context, address string, id uint64) (val customauthtypes.Pairing, found bool)
	UnpackPairingPubKey(pairing customauthtypes.Pairing) (cryptotypes.PubKey, error)
	GetPairingPubKeyInternal(ctx context.Context, user sdk.AccAddress, pairingIndex string) (pairing.PubKeyInternal, error)
}

type OrderKeeper interface {
	GetOrderAndBody(ctx sdk.Context, orderOwner string, orderHash string) (ordertypes.Order, ordertypes.OrderBody, error)
	AddContractedAmount(ctx context.Context, orderOwner string, orderHash string, additionalContractedAmount sdkmath.Int) error
}

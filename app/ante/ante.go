package ante

import (
	"cosmossdk.io/x/tx/signing"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	ibcante "github.com/cosmos/ibc-go/v8/modules/core/ante"
	ibckeeper "github.com/cosmos/ibc-go/v8/modules/core/keeper"

	customante "gluon/x/customauth/ante"
)

func NewAnteHandler(
	accountKeeper ante.AccountKeeper,
	bankKeeper customante.BankKeeper,
	feegrantKeeper ante.FeegrantKeeper,
	customAuthKeeper customante.CustomAuthKeeper,
	signModeHandler *signing.HandlerMap,
	sigGasConsumer ante.SignatureVerificationGasConsumer,
	channelKeeper *ibckeeper.Keeper,
	TxEncoder sdk.TxEncoder,
	operatorMsgMatchHandler customante.OperatorMsgMatchHandler,
) sdk.AnteHandler {
	return sdk.ChainAnteDecorators(
		// Set up the context with a gas meter.
		// Must be called before gas consumption occurs in any other decorator.
		ante.NewSetUpContextDecorator(),
		// Ensure the tx does not contain any extension options.
		ante.NewExtensionOptionsDecorator(nil),
		// Ensure the tx passes ValidateBasic.
		ante.NewValidateBasicDecorator(),
		// Ensure the tx has not reached a height timeout.
		ante.NewTxTimeoutHeightDecorator(),
		// Ensure the tx memo <= max memo characters.
		ante.NewValidateMemoDecorator(accountKeeper),
		// Ensure the tx's gas limit is > the gas consumed based on the tx size.
		// Side effect: consumes gas from the gas meter.
		ante.NewConsumeGasForTxSizeDecorator(accountKeeper),
		// Ensure the feepayer (fee granter or first signer) has enough funds to pay for the tx.
		// Side effect: deducts fees from the fee payer. Sets the tx priority in context.
		customante.NewDeductFeeDecorator(accountKeeper, bankKeeper, feegrantKeeper, nil, operatorMsgMatchHandler),
		// Set public keys in the context for fee-payer and all signers.
		// Contract: must be called before all signature verification decorators.
		customante.NewSetPubKeyDecorator(accountKeeper, customAuthKeeper, operatorMsgMatchHandler),
		// Ensure that the tx's count of signatures is <= the tx signature limit.
		ante.NewValidateSigCountDecorator(accountKeeper),
		// Ensure that the tx's gas limit is > the gas consumed based on signature verification.
		// Side effect: consumes gas from the gas meter.
		ante.NewSigGasConsumeDecorator(accountKeeper, sigGasConsumer),
		// Ensure that the tx's signatures are valid. For each signature, ensure
		// that the signature's sequence number (a.k.a nonce) matches the
		// account sequence number of the signer.
		// Note: does not consume gas from the gas meter.
		customante.NewSigVerificationDecorator(accountKeeper, signModeHandler, customAuthKeeper, operatorMsgMatchHandler),
		// Side effect: increment the nonce for all tx signers.
		ante.NewIncrementSequenceDecorator(accountKeeper),
		// Ensure that the tx is not a IBC packet or update message that has already been processed.
		ibcante.NewRedundantRelayDecorator(channelKeeper),
	)
}

var DefaultSigVerificationGasConsumer = ante.DefaultSigVerificationGasConsumer

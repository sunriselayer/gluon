package ante

import (
	"fmt"

	"google.golang.org/protobuf/types/known/anypb"

	errorsmod "cosmossdk.io/errors"
	txsigning "cosmossdk.io/x/tx/signing"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"

	"github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"

	sdkante "github.com/cosmos/cosmos-sdk/x/auth/ante"

	"gluon/x/customauth/types/operator"
	"gluon/x/customauth/types/pairing"
)

var (
	// simulation signature values used to estimate gas consumption
	key                = make([]byte, secp256k1.PubKeySize)
	simSecp256k1Pubkey = &secp256k1.PubKey{Key: key}
	simSecp256k1Sig    [64]byte
)

// Consume parameter-defined amount of gas for each signature according to the passed-in SignatureVerificationGasConsumer function
// before calling the next AnteHandler
// CONTRACT: Pubkeys are set in context for all signers before this decorator runs
// CONTRACT: Tx must implement SigVerifiableTx interface
type SigGasConsumeDecorator struct {
	ak             AccountKeeper
	sigGasConsumer sdkante.SignatureVerificationGasConsumer
}

func NewSigGasConsumeDecorator(ak AccountKeeper, sigGasConsumer sdkante.SignatureVerificationGasConsumer) SigGasConsumeDecorator {
	if sigGasConsumer == nil {
		sigGasConsumer = sdkante.DefaultSigVerificationGasConsumer
	}

	return SigGasConsumeDecorator{
		ak:             ak,
		sigGasConsumer: sigGasConsumer,
	}
}

func (sgcd SigGasConsumeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	sigTx, ok := tx.(authsigning.SigVerifiableTx)
	if !ok {
		return ctx, errorsmod.Wrap(sdkerrors.ErrTxDecode, "invalid transaction type")
	}

	params := sgcd.ak.GetParams(ctx)
	sigs, err := sigTx.GetSignaturesV2()
	if err != nil {
		return ctx, err
	}

	// stdSigs contains the sequence number, account number, and signatures.
	// When simulating, this would just be a 0-length slice.
	// signers, err := sigTx.GetSigners()
	// if err != nil {
	// 	return ctx, err
	// }

	for _, sig := range sigs {
		// signerAcc, err := sdkante.GetSignerAcc(ctx, sgcd.ak, signers[i])
		// if err != nil {
		// 	return ctx, err
		// }

		// <gluon>
		pubKey := sig.PubKey
		// </gluon>

		// In simulate mode the transaction comes with no signatures, thus if the
		// account's pubkey is nil, both signature verification and gasKVStore.Set()
		// shall consume the largest amount, i.e. it takes more gas to verify
		// secp256k1 keys than ed25519 ones.
		if simulate && pubKey == nil {
			pubKey = simSecp256k1Pubkey
		}

		// make a SignatureV2 with PubKey filled in from above
		sig = signing.SignatureV2{
			PubKey:   pubKey,
			Data:     sig.Data,
			Sequence: sig.Sequence,
		}

		err = sgcd.sigGasConsumer(ctx.GasMeter(), sig, params)
		if err != nil {
			return ctx, err
		}
	}

	return next(ctx, tx, simulate)
}

// SigVerificationDecorator verifies all signatures for a tx and return an error if any are invalid. Note,
// the SigVerificationDecorator will not check signatures on ReCheck.
//
// CONTRACT: Pubkeys are set in context for all signers before this decorator runs
// CONTRACT: Tx must implement SigVerifiableTx interface
type SigVerificationDecorator struct {
	ak              AccountKeeper
	signModeHandler *txsigning.HandlerMap

	cak             CustomAuthKeeper
	msgMatchHandler OperatorMsgMatchHandler
}

func NewSigVerificationDecorator(
	ak AccountKeeper,
	signModeHandler *txsigning.HandlerMap,
	cak CustomAuthKeeper,
	msgMatchHandler OperatorMsgMatchHandler,
) SigVerificationDecorator {
	return SigVerificationDecorator{
		ak:              ak,
		signModeHandler: signModeHandler,

		cak:             cak,
		msgMatchHandler: msgMatchHandler,
	}
}

func (svd SigVerificationDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	sigTx, ok := tx.(authsigning.Tx)
	if !ok {
		return ctx, errorsmod.Wrap(sdkerrors.ErrTxDecode, "invalid transaction type")
	}

	// stdSigs contains the sequence number, account number, and signatures.
	// When simulating, this would just be a 0-length slice.
	sigs, err := sigTx.GetSignaturesV2()
	if err != nil {
		return ctx, err
	}

	signers, err := sigTx.GetSigners()
	if err != nil {
		return ctx, err
	}

	// check that signer length and signature length are the same
	if len(sigs) != len(signers) {
		return ctx, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "invalid number of signer;  expected: %d, got %d", len(signers), len(sigs))
	}

	// <gluon>
	pubKeys, err := sigTx.GetPubKeys()
	if err != nil {
		return ctx, err
	}
	isOperatorTx := IsOperatorTx(tx, svd.msgMatchHandler)
	// </gluon>

	for i, sig := range sigs {
		acc, err := sdkante.GetSignerAcc(ctx, svd.ak, signers[i])
		if err != nil {
			return ctx, err
		}

		// <gluon>
		pubKey := pubKeys[i]
		var signerPubKey cryptotypes.PubKey
		switch pk := pubKey.(type) {
		case *operator.PubKey:
			if !isOperatorTx {
				return ctx, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "operator pubkey can be used only for designated messages")
			}
			// get operator pubkey
			signerPubKey, err = svd.cak.GetOperatorPubKey(ctx)
			if err != nil {
				return ctx, err
			}
			// Check account sequence number.
			if sig.Sequence != acc.GetSequence() {
				return ctx, errorsmod.Wrapf(
					sdkerrors.ErrWrongSequence,
					"account sequence mismatch, expected %d, got %d", acc.GetSequence(), sig.Sequence,
				)
			}

		case *pairing.PubKey:
			pubKeyInternal, err := svd.cak.GetPairingPubKeyInternal(ctx, acc.GetAddress(), pk.PairingIndex)
			if err != nil {
				return ctx, err
			}
			// create multisig pubkey
			pairingPubKey, err := svd.cak.UnpackPublicKey(pubKeyInternal.GetPairingPublicKey())
			if err != nil {
				return ctx, err
			}
			operatorPubKey, err := svd.cak.UnpackPublicKey(pubKeyInternal.GetOperatorPublicKey())
			if err != nil {
				return ctx, err
			}
			signerPubKey = multisig.NewLegacyAminoPubKey(2, []cryptotypes.PubKey{pairingPubKey, operatorPubKey})
			// Check account sequence number. 1 is added for two-step signatures.
			if sig.Sequence != acc.GetSequence()+1 {
				return ctx, errorsmod.Wrapf(
					sdkerrors.ErrWrongSequence,
					"account sequence mismatch, expected %d, got %d", acc.GetSequence(), sig.Sequence,
				)
			}

		default:
			// For gen-tx
			if ctx.BlockHeight() == 0 {
				return ctx, nil
			}
			// force error for EOA
			return ctx, errorsmod.Wrap(sdkerrors.ErrInvalidPubKey, "not permitted public key type")
		}
		// </gluon>

		// retrieve signer data
		genesis := ctx.BlockHeight() == 0
		chainID := ctx.ChainID()
		var accNum uint64
		if !genesis {
			accNum = acc.GetAccountNumber()
		}

		// no need to verify signatures on recheck tx
		if !simulate && !ctx.IsReCheckTx() && ctx.IsSigverifyTx() {
			anyPk, _ := codectypes.NewAnyWithValue(pubKey)

			signerData := txsigning.SignerData{
				Address:       acc.GetAddress().String(),
				ChainID:       chainID,
				AccountNumber: accNum,
				Sequence:      acc.GetSequence(),
				PubKey: &anypb.Any{
					TypeUrl: anyPk.TypeUrl,
					Value:   anyPk.Value,
				},
			}
			adaptableTx, ok := tx.(authsigning.V2AdaptableTx)
			if !ok {
				return ctx, fmt.Errorf("expected tx to implement V2AdaptableTx, got %T", tx)
			}
			txData := adaptableTx.GetSigningTxData()
			err = authsigning.VerifySignature(ctx, signerPubKey, signerData, sig.Data, svd.signModeHandler, txData)
			if err != nil {
				errMsg := fmt.Sprintf("signature verification failed; please verify account number (%d) and chain-id (%s): (%s)", accNum, chainID, err.Error())
				return ctx, errorsmod.Wrap(sdkerrors.ErrUnauthorized, errMsg)
			}
		}
	}

	return next(ctx, tx, simulate)
}

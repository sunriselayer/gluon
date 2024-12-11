package ante

import (
	"fmt"

	"google.golang.org/protobuf/types/known/anypb"

	errorsmod "cosmossdk.io/errors"
	txsigning "cosmossdk.io/x/tx/signing"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"

	sdkante "github.com/cosmos/cosmos-sdk/x/auth/ante"

	"gluon/x/customauth/types/operator"
	"gluon/x/customauth/types/pairing"
)

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
		var pubKey cryptotypes.PubKey
		switch pk := pubKeys[i].(type) {
		case *operator.PubKey:
			if !isOperatorTx {
				return ctx, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "operator pubkey can be used only for designated messages")
			}
			pubKey, err = svd.cak.GetOperatorPubKey(ctx)
			if err != nil {
				return ctx, err
			}

		case *pairing.PubKey:
			pubKeyInternal, err := svd.cak.GetPairingPubKeyInternal(ctx, acc.GetAddress(), pk.PairingIndex)
			if err != nil {
				return ctx, err
			}
			pubKey = &pubKeyInternal

		default:
			// force error for EOA
			return ctx, errorsmod.Wrap(sdkerrors.ErrInvalidPubKey, "not permitted public key type")
		}
		// </gluon>

		// Check account sequence number.
		if sig.Sequence != acc.GetSequence() {
			return ctx, errorsmod.Wrapf(
				sdkerrors.ErrWrongSequence,
				"account sequence mismatch, expected %d, got %d", acc.GetSequence(), sig.Sequence,
			)
		}

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
			err = authsigning.VerifySignature(ctx, pubKey, signerData, sig.Data, svd.signModeHandler, txData)
			if err != nil {
				errMsg := fmt.Sprintf("signature verification failed; please verify account number (%d) and chain-id (%s): (%s)", accNum, chainID, err.Error())
				return ctx, errorsmod.Wrap(sdkerrors.ErrUnauthorized, errMsg)
			}
		}
	}

	return next(ctx, tx, simulate)
}

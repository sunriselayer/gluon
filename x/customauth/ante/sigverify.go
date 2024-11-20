package ante

import (
	"fmt"

	"google.golang.org/protobuf/types/known/anypb"

	errorsmod "cosmossdk.io/errors"
	txsigning "cosmossdk.io/x/tx/signing"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdkante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	types "gluon/x/customauth/types"
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
	// <gluon>
	match := Match(tx, svd.msgMatchHandler)
	// </gluon>

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

	for i, sig := range sigs {
		acc, err := sdkante.GetSignerAcc(ctx, svd.ak, signers[i])
		if err != nil {
			return ctx, err
		}

		// <gluon>
		var pubKey cryptotypes.PubKey
		if !match {
			if true {
				// retrieve pubkey
				pubKey = acc.GetPubKey()
				if !simulate && pubKey == nil {
					return ctx, errorsmod.Wrap(sdkerrors.ErrInvalidPubKey, "pubkey on account is not set")
				}
			} else {
				pairingId := uint64(0)
				pairing, found := svd.cak.GetPairing(ctx, acc.GetAddress().String(), pairingId)
				if !found {
					return ctx, errorsmod.Wrapf(types.ErrPairingNotFound, "address: %s, pairing_id: %d", acc.GetAddress().String(), pairingId)
				}
				params := svd.cak.GetParams(ctx)

				if ctx.BlockTime().Compare(pairing.CreatedAt.Add(params.ParingDelay)) < 0 {
					return ctx, errorsmod.Wrapf(types.ErrPairingDelayPeriod, "address: %s, pairing_id: %d", acc.GetAddress().String(), pairingId)
				}

				pubKey, err = svd.cak.GetPairingPubKey(ctx, pairing)
				if err != nil {
					return ctx, err
				}
			}
		} else {
			pubKey, err = svd.cak.GetOperatorPubKey(ctx)
			if err != nil {
				return ctx, err
			}
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

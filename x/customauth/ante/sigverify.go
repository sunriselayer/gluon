package ante

import (
	"bytes"
	"encoding/base64"
	"fmt"

	"google.golang.org/protobuf/types/known/anypb"

	errorsmod "cosmossdk.io/errors"
	txsigning "cosmossdk.io/x/tx/signing"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"

	sdkante "github.com/cosmos/cosmos-sdk/x/auth/ante"
)

var (
	// simulation signature values used to estimate gas consumption
	key                = make([]byte, secp256k1.PubKeySize)
	simSecp256k1Pubkey = &secp256k1.PubKey{Key: key}
)

// SetPubKeyDecorator sets PubKeys in context for any signer which does not already have pubkey set
// PubKeys must be set in context for all signers before any other sigverify decorators run
// CONTRACT: Tx must implement SigVerifiableTx interface
type SetPubKeyDecorator struct {
	ak AccountKeeper

	cak             CustomAuthKeeper
	msgMatchHandler OperatorMsgMatchHandler
}

func NewSetPubKeyDecorator(ak AccountKeeper, cak CustomAuthKeeper, msgMatchHandler OperatorMsgMatchHandler) SetPubKeyDecorator {
	return SetPubKeyDecorator{
		ak:              ak,
		cak:             cak,
		msgMatchHandler: msgMatchHandler,
	}
}

func (spkd SetPubKeyDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	// <gluon>
	txWithMemo, ok := tx.(sdk.TxWithMemo)
	if !ok {
		return ctx, errorsmod.Wrap(sdkerrors.ErrTxDecode, "invalid transaction type")
	}

	isOperatorTx := IsOperatorTx(tx, spkd.msgMatchHandler)
	// skip operator tx
	if isOperatorTx {
		return next(ctx, tx, simulate)
	}

	metadata, err := GetPairingId(txWithMemo)
	if err != nil {
		return ctx, err
	}
	var pairingId map[string]uint64
	if metadata != nil {
		pairingId = metadata.PairingId
	} else {
		pairingId = make(map[string]uint64)
	}
	// </gluon>

	sigTx, ok := tx.(authsigning.SigVerifiableTx)
	if !ok {
		return ctx, errorsmod.Wrap(sdkerrors.ErrTxDecode, "invalid tx type")
	}

	pubkeys, err := sigTx.GetPubKeys()
	if err != nil {
		return ctx, err
	}

	signers, err := sigTx.GetSigners()
	if err != nil {
		return sdk.Context{}, err
	}

	signerStrs := make([]string, len(signers))
	for i, pk := range pubkeys {
		var err error
		signerStrs[i], err = spkd.ak.AddressCodec().BytesToString(signers[i])
		if err != nil {
			return sdk.Context{}, err
		}

		// <gluon>
		// If the signer is pairing key, skip setting pub key.
		if _, found := pairingId[signerStrs[i]]; found {
			continue
		}
		// </gluon>

		// PublicKey was omitted from slice since it has already been set in context
		if pk == nil {
			if !simulate {
				continue
			}
			pk = simSecp256k1Pubkey
		}
		// Only make check if simulate=false
		if !simulate && !bytes.Equal(pk.Address(), signers[i]) && ctx.IsSigverifyTx() {
			return ctx, errorsmod.Wrapf(sdkerrors.ErrInvalidPubKey,
				"pubKey does not match signer address %s with signer index: %d", signerStrs[i], i)
		}

		acc, err := sdkante.GetSignerAcc(ctx, spkd.ak, signers[i])
		if err != nil {
			return ctx, err
		}
		// account already has pubkey set,no need to reset
		if acc.GetPubKey() != nil {
			continue
		}
		err = acc.SetPubKey(pk)
		if err != nil {
			return ctx, errorsmod.Wrap(sdkerrors.ErrInvalidPubKey, err.Error())
		}
		spkd.ak.SetAccount(ctx, acc)
	}

	// Also emit the following events, so that txs can be indexed by these
	// indices:
	// - signature (via `tx.signature='<sig_as_base64>'`),
	// - concat(address,"/",sequence) (via `tx.acc_seq='cosmos1abc...def/42'`).
	sigs, err := sigTx.GetSignaturesV2()
	if err != nil {
		return ctx, err
	}

	var events sdk.Events
	for i, sig := range sigs {
		events = append(events, sdk.NewEvent(sdk.EventTypeTx,
			sdk.NewAttribute(sdk.AttributeKeyAccountSequence, fmt.Sprintf("%s/%d", signerStrs[i], sig.Sequence)),
		))

		sigBzs, err := signatureDataToBz(sig.Data)
		if err != nil {
			return ctx, err
		}
		for _, sigBz := range sigBzs {
			events = append(events, sdk.NewEvent(sdk.EventTypeTx,
				sdk.NewAttribute(sdk.AttributeKeySignature, base64.StdEncoding.EncodeToString(sigBz)),
			))
		}
	}

	ctx.EventManager().EmitEvents(events)

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

	// <gluon>
	isOperatorTx := IsOperatorTx(tx, svd.msgMatchHandler)
	metadata, err := GetPairingId(sigTx)
	if err != nil {
		return ctx, err
	}
	var pairingId map[string]uint64
	if metadata != nil {
		pairingId = metadata.PairingId
	} else {
		pairingId = make(map[string]uint64)
	}
	// </gluon>

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
		if isOperatorTx {
			pubKey, err = svd.cak.GetOperatorPubKey(ctx)
			if err != nil {
				return ctx, err
			}
		} else if id, found := pairingId[acc.GetAddress().String()]; found {
			pubKey, err = svd.cak.GetDoubleSigPubKey(ctx, acc.GetAddress().String(), id)
			if err != nil {
				return ctx, err
			}
		} else {
			// force error for EOA
			return ctx, errorsmod.Wrap(sdkerrors.ErrInvalidPubKey, "pairing pubkey on account is not set")
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

// signatureDataToBz converts a SignatureData into raw bytes signature.
// For SingleSignatureData, it returns the signature raw bytes.
// For MultiSignatureData, it returns an array of all individual signatures,
// as well as the aggregated signature.
func signatureDataToBz(data signing.SignatureData) ([][]byte, error) {
	if data == nil {
		return nil, fmt.Errorf("got empty SignatureData")
	}

	switch data := data.(type) {
	case *signing.SingleSignatureData:
		return [][]byte{data.Signature}, nil
	case *signing.MultiSignatureData:
		sigs := [][]byte{}
		var err error

		for _, d := range data.Signatures {
			nestedSigs, err := signatureDataToBz(d)
			if err != nil {
				return nil, err
			}
			sigs = append(sigs, nestedSigs...)
		}

		multiSignature := cryptotypes.MultiSignature{
			Signatures: sigs,
		}
		aggregatedSig, err := multiSignature.Marshal()
		if err != nil {
			return nil, err
		}
		sigs = append(sigs, aggregatedSig)

		return sigs, nil
	default:
		return nil, sdkerrors.ErrInvalidType.Wrapf("unexpected signature data type %T", data)
	}
}

package pairing

import (
	"bytes"
	"fmt"

	cometbytes "github.com/cometbft/cometbft/libs/bytes"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	multisig "github.com/cosmos/cosmos-sdk/crypto/types/multisig"
	txsigning "github.com/cosmos/cosmos-sdk/types/tx/signing"

	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ cryptotypes.PubKey = &PubKey{}

func (pk PubKey) Address() cometbytes.HexBytes {
	return pk.User
}

func (pk PubKey) Bytes() []byte {
	return nil
}

func (pk PubKey) VerifySignature(msg []byte, sig []byte) bool {
	return false
}

func (pk PubKey) Equals(other cryptotypes.PubKey) bool {
	switch other := other.(type) {
	case *PubKey:
		return bytes.Equal(pk.User, other.User) && pk.PairingId == other.PairingId
	default:
		return false
	}
}

func (pk PubKey) Type() string {
	return "pairing"
}

var _ multisig.PubKey = &PubKeyInternal{}

func (pk PubKeyInternal) Address() cometbytes.HexBytes {
	return pk.User
}

func (pk PubKeyInternal) Bytes() []byte {
	return nil
}

func (pk PubKeyInternal) VerifySignature(msg []byte, sig []byte) bool {
	return false
}

func (pk PubKeyInternal) Equals(other cryptotypes.PubKey) bool {
	return false
}

func (pk PubKeyInternal) Type() string {
	return "pairing"
}

func (pk PubKeyInternal) VerifyMultisignature(getSignBytes multisig.GetSignBytesFunc, sig *txsigning.MultiSignatureData) error {
	if !sig.BitArray.GetIndex(0) || sig.BitArray.GetIndex(1) {
		return errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "both user pairing key and operator key must sign")
	}
	if len(sig.Signatures) != 2 {
		return errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "must be only two signatures")
	}

	pubKeys := pk.GetPubKeys()

	for i, data := range sig.Signatures {
		switch data := data.(type) {
		case *txsigning.SingleSignatureData:
			msg, err := getSignBytes(data.SignMode)
			if err != nil {
				return err
			}
			if !pubKeys[i].VerifySignature(msg, data.Signature) {
				return fmt.Errorf("unable to verify signature at index %d", i)
			}

		case *txsigning.MultiSignatureData:
			nestedMultisigPk, ok := pubKeys[i].(multisig.PubKey)
			if !ok {
				return fmt.Errorf("unable to parse pubkey of index %d", i)
			}
			if err := nestedMultisigPk.VerifyMultisignature(getSignBytes, data); err != nil {
				return err
			}

		default:
			return fmt.Errorf("improper signature data type for index %d", i)
		}
	}

	return nil
}

func (pk PubKeyInternal) GetPubKeys() []cryptotypes.PubKey {
	pubKeys := make([]cryptotypes.PubKey, 2)

	pubKeys[0] = pk.PairingPublicKey.GetCachedValue().(cryptotypes.PubKey)
	pubKeys[1] = pk.PairingPublicKey.GetCachedValue().(cryptotypes.PubKey)

	return pubKeys
}

func (pk PubKeyInternal) GetThreshold() uint {
	return 2
}

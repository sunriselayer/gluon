package pairing

import (
	"bytes"

	cometbytes "github.com/cometbft/cometbft/libs/bytes"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	multisig "github.com/cosmos/cosmos-sdk/crypto/types/multisig"
	txsigning "github.com/cosmos/cosmos-sdk/types/tx/signing"
)

var _ cryptotypes.PubKey = &PubKeyNotVerifiable{}

func (pk PubKeyNotVerifiable) Address() cometbytes.HexBytes {
	return pk.User
}

func (pk PubKeyNotVerifiable) Bytes() []byte {
	return nil
}

func (pk PubKeyNotVerifiable) VerifySignature(msg []byte, sig []byte) bool {
	return false
}

func (pk PubKeyNotVerifiable) Equals(other cryptotypes.PubKey) bool {
	switch other := other.(type) {
	case *PubKeyNotVerifiable:
		return bytes.Equal(pk.User, other.User) && pk.PairingId == other.PairingId
	default:
		return false
	}
}

func (pk PubKeyNotVerifiable) Type() string {
	return "pairing"
}

var _ multisig.PubKey = &PubKey{}

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
	return false
}

func (pk PubKey) Type() string {
	return "pairing"
}

func (pk PubKey) VerifyMultisignature(getSignBytes multisig.GetSignBytesFunc, sig *txsigning.MultiSignatureData) error {
	return nil
}

func (pk PubKey) GetPubKeys() []cryptotypes.PubKey {
	return nil
}

func (pk PubKey) GetThreshold() uint {
	return 2
}

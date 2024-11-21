package operator

import (
	"bytes"

	cometbytes "github.com/cometbft/cometbft/libs/bytes"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
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
		return bytes.Equal(pk.User, other.User)
	default:
		return false
	}

}

func (pk PubKeyNotVerifiable) Type() string {
	return "operator"
}

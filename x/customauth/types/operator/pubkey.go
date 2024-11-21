package operator

import (
	"bytes"

	cometbytes "github.com/cometbft/cometbft/libs/bytes"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
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
		return bytes.Equal(pk.User, other.User)
	default:
		return false
	}

}

func (pk PubKey) Type() string {
	return "operator"
}

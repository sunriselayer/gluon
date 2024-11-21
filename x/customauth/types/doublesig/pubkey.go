package doublesig

import (
	cometbytes "github.com/cometbft/cometbft/libs/bytes"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	multisig "github.com/cosmos/cosmos-sdk/crypto/types/multisig"
	txsigning "github.com/cosmos/cosmos-sdk/types/tx/signing"
)

var _ multisig.PubKey = &PubKey{}

func (pk PubKey) Address() cometbytes.HexBytes {
	return nil
}

func (pk PubKey) Bytes() []byte {
	return nil
}

func (pk PubKey) VerifySignature(msg []byte, sig []byte) bool {
	return true
}

func (pk PubKey) Equals(other cryptotypes.PubKey) bool {
	return true
}

func (pk PubKey) Type() string {
	return "doublesig"
}

func (pk PubKey) VerifyMultisignature(getSignBytes multisig.GetSignBytesFunc, sig *txsigning.MultiSignatureData) error {
	return nil
}

func (pk PubKey) GetPubKeys() []cryptotypes.PubKey {
	return nil
}

func (pk PubKey) GetThreshold() uint {
	return 0
}

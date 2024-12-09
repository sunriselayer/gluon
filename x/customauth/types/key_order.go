package types

import (
	"encoding/binary"
	"fmt"
)

var _ binary.ByteOrder

const (
	// PairingKeyPrefix is the prefix to retrieve all Pairing
	PairingKeyPrefix = "Pairing/value/"
)

// PairingKey returns the store key to retrieve a Pairing from the index fields
func PairingKey(
	owner string,
	index string,
) []byte {
	var key []byte

	ownerBytes := []byte(owner)
	key = append(key, ownerBytes...)
	key = append(key, []byte("/")...)

	hashBytes := []byte(index)
	key = append(key, hashBytes...)
	key = append(key, []byte("/")...)

	return key
}

// PairingKeyPrefixByOwner
func PairingKeyPrefixByOwner(
	owner string,
) []byte {
	return KeyPrefix(fmt.Sprintf("%s%s/", PairingKeyPrefix, owner))
}

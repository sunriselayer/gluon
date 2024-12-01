package types

import (
	"encoding/binary"
	"fmt"
)

var _ binary.ByteOrder

const (
	// PositionKeyPrefix is the prefix to retrieve all Position
	PositionKeyPrefix = "Position/value/"
)

// PositionKey returns the store key to retrieve a Position from the index fields
func PositionKey(
	owner string,
	orderHash string,
) []byte {
	var key []byte

	ownerBytes := []byte(owner)
	key = append(key, ownerBytes...)
	key = append(key, []byte("/")...)

	orderHashBytes := []byte(orderHash)
	key = append(key, orderHashBytes...)
	key = append(key, []byte("/")...)

	return key
}

// PositionKeyPrefixByOwner
func PositionKeyPrefixByOwner(
	owner string,
) []byte {
	return KeyPrefix(fmt.Sprintf("%s%s/", PositionKeyPrefix, owner))
}

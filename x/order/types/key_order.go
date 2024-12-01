package types

import (
	"encoding/binary"
	"fmt"
)

var _ binary.ByteOrder

const (
	// OrderKeyPrefix is the prefix to retrieve all Order
	OrderKeyPrefix = "Order/value/"
)

// OrderKey returns the store key to retrieve a Order from the index fields
func OrderKey(
	owner string,
	hash string,
) []byte {
	var key []byte

	ownerBytes := []byte(owner)
	key = append(key, ownerBytes...)
	key = append(key, []byte("/")...)

	hashBytes := []byte(hash)
	key = append(key, hashBytes...)
	key = append(key, []byte("/")...)

	return key
}

// OrderKeyPrefix
func OrderKeyPrefixByOwner(
	owner string,
) []byte {
	return KeyPrefix(fmt.Sprintf("%s%s/", OrderKeyPrefix, owner))
}

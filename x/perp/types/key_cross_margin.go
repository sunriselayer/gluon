package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// CrossMarginKeyPrefix is the prefix to retrieve all CrossMargin
	CrossMarginKeyPrefix = "CrossMargin/value/"
)

// CrossMarginKey returns the store key to retrieve a CrossMargin from the index fields
func CrossMarginKey(
	owner string,
) []byte {
	var key []byte

	ownerBytes := []byte(owner)
	key = append(key, ownerBytes...)
	key = append(key, []byte("/")...)

	return key
}

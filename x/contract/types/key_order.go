package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// OrderKeyPrefix is the prefix to retrieve all Order
	OrderKeyPrefix = "Order/value/"
)

// OrderKey returns the store key to retrieve a Order from the index fields
func OrderKey(
	id string,
) []byte {
	var key []byte

	indexBytes := []byte(id)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}

package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// SortedOrderKeyPrefix is the prefix to retrieve all SortedOrder
	SortedOrderKeyPrefix = "SortedOrder/value/"
)

// SortedOrderKey returns the store key to retrieve a SortedOrder from the index fields
func SortedOrderKey(
	expiry uint64,
	id string,
) []byte {
	var key []byte

	var expiryBytes []byte
	binary.BigEndian.Uint64(expiryBytes)
	key = append(key, expiryBytes...)
	key = append(key, []byte("/")...)

	indexBytes := []byte(id)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}

package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// SortedOrderKeyPrefix is the prefix to retrieve all SortedOrder
	SortedOrderKeyPrefix = "SortedOrder/value/"
)

// SortedOrderKey returns the store key to retrieve a SortedOrder from the index fields
func SortedOrderKey(
	expiry string,
	index string,
) []byte {
	var key []byte

	expiryBytes := []byte(expiry)
	key = append(key, expiryBytes...)
	key = append(key, []byte("/")...)

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}

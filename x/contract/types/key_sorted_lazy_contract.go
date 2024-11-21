package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// SortedLazyContractKeyPrefix is the prefix to retrieve all SortedLazyContract
	SortedLazyContractKeyPrefix = "SortedLazyContract/value/"
)

// SortedLazyContractKey returns the store key to retrieve a SortedLazyContract from the index fields
func SortedLazyContractKey(
	expiry uint64,
	id uint64,
) []byte {
	var key []byte

	expiryBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(expiryBytes, expiry)
	key = append(key, expiryBytes...)
	key = append(key, []byte("/")...)

	indexBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(indexBytes, id)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}

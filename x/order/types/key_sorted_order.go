package types

import (
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ binary.ByteOrder

const (
	// SortedOrderKeyPrefix is the prefix to retrieve all SortedOrder
	SortedOrderKeyPrefix = "SortedOrder/value/"
)

// SortedOrderKey returns the store key to retrieve a SortedOrder from the index fields
func SortedOrderKey(
	seconds uint64,
	nanos uint32,
	orderHash string,
) []byte {
	var key []byte

	secondsBytes := sdk.Uint64ToBigEndian(seconds)
	key = append(key, secondsBytes...)
	key = append(key, []byte("/")...)

	nanosBytes := uint32ToBigEndian(nanos)
	key = append(key, nanosBytes...)
	key = append(key, []byte("/")...)

	orderHashBytes := []byte(orderHash)
	key = append(key, orderHashBytes...)
	key = append(key, []byte("/")...)

	return key
}

func uint32ToBigEndian(i uint32) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, i)
	return b
}

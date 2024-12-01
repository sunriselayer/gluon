package types

import (
	"encoding/binary"
	"fmt"
)

var _ binary.ByteOrder

const (
	// PositionPriceQuantityKeyPrefix is the prefix to retrieve all PositionPriceQuantity
	PositionPriceQuantityKeyPrefix = "PositionPriceQuantity/value/"
)

// PositionPriceQuantityKey returns the store key to retrieve a PositionPriceQuantity from the index fields
func PositionPriceQuantityKey(
	owner string,
	positionOrderHash string,
	price string,
) []byte {
	var key []byte

	ownerBytes := []byte(owner)
	key = append(key, ownerBytes...)
	key = append(key, []byte("/")...)

	positionOrderHashBytes := []byte(positionOrderHash)
	key = append(key, positionOrderHashBytes...)
	key = append(key, []byte("/")...)

	priceBytes := []byte(price)
	key = append(key, priceBytes...)
	key = append(key, []byte("/")...)

	return key
}

// PositionPriceQuantityKeyPrefixByOwnerAndPositionOrderHash
func PositionPriceQuantityKeyPrefixByOwnerAndPositionOrderHash(
	owner string,
	positionOrderHash string,
) []byte {
	return KeyPrefix(fmt.Sprintf("%s%s/%s/", PositionPriceQuantityKeyPrefix, owner, positionOrderHash))
}

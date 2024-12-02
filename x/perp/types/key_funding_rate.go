package types

import (
	"encoding/binary"
	"fmt"
)

var _ binary.ByteOrder

const (
	// FundingRateKeyPrefix is the prefix to retrieve all FundingRate
	FundingRateKeyPrefix = "FundingRate/value/"
)

// FundingRateKey returns the store key to retrieve a FundingRate from the index fields
func FundingRateKey(
	denomBase string,
	denomQuote string,
	seconds uint64,
	nanos uint32,
) []byte {
	var key []byte

	denomBaseBytes := []byte(denomBase)
	key = append(key, denomBaseBytes...)
	key = append(key, []byte("/")...)

	denomQuoteBytes := []byte(denomQuote)
	key = append(key, denomQuoteBytes...)
	key = append(key, []byte("/")...)

	secondsBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(secondsBytes, seconds)
	key = append(key, secondsBytes...)
	key = append(key, []byte("/")...)

	nanosBytes := uint32ToBigEndian(nanos)
	key = append(key, nanosBytes...)
	key = append(key, []byte("/")...)

	return key
}

func uint32ToBigEndian(i uint32) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, i)
	return b
}

// FundingRateKeyPrefixByDenoms
func FundingRateKeyPrefixByDenoms(
	denomBase string,
	denomQuote string,
) []byte {
	return KeyPrefix(fmt.Sprintf("%s%s/%s/", FundingRateKeyPrefix, denomBase, denomQuote))
}

package types

import (
	"encoding/binary"
	"fmt"
)

var _ binary.ByteOrder

const (
	PositionKey      = "Position/value/"
	PositionCountKey = "Position/count/"
)

// PositionKeyPrefix
func PositionKeyPrefix(
	owner string,
) []byte {
	return KeyPrefix(fmt.Sprintf("%s%s", PositionKey, owner))
}

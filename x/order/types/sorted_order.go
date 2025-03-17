package types

import (
	"time"

	gogo "github.com/gogo/protobuf/types"
)

func (sortedOrder SortedOrder) GetTime() time.Time {
	timestamp := gogo.Timestamp{Seconds: int64(sortedOrder.Seconds), Nanos: int32(sortedOrder.Nanos)}
	t, err := gogo.TimestampFromProto(&timestamp)
	if err != nil {
		panic(err)
	}

	return t
}

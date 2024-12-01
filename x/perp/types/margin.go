package types

import (
	"fmt"
)

func GetCrossMarginAddressModule(positionOwner string) string {
	return fmt.Sprintf("%s/margin/%s", ModuleName, positionOwner)
}

func GetIsolatedMarginAddressModule(positionOwner string, positionId uint64) string {
	return fmt.Sprintf("%s/margin/%s/%d", ModuleName, positionOwner, positionId)
}

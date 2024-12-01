package types

import (
	"fmt"
)

func GetCrossMarginAddressModule(positionOwner string) string {
	return fmt.Sprintf("%s/margin/%s", ModuleName, positionOwner)
}

func GetIsolatedMarginAddressModule(positionOwner string, orderHash string) string {
	return fmt.Sprintf("%s/margin/%s/%s", ModuleName, positionOwner, orderHash)
}

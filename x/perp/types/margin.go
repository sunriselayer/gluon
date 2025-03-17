package types

import (
	"fmt"
)

func GetMarginAddressModule() string {
	return fmt.Sprintf("%s/margin", ModuleName)
}

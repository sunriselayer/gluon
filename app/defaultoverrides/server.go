package defaultoverrides

import (
	"fmt"

	"gluon/app/consts"

	serverconfig "github.com/cosmos/cosmos-sdk/server/config"
)

func DefaultServerConfig() *serverconfig.Config {
	cfg := serverconfig.DefaultConfig()

	cfg.MinGasPrices = fmt.Sprintf("%v%s", consts.DefaultMinGasPrice, consts.BondDenom)

	return cfg
}

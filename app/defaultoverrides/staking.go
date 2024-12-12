package defaultoverrides

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"

	staking "github.com/cosmos/cosmos-sdk/x/staking"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"gluon/app/consts"
)

// StakingModuleBasic wraps the x/staking module in order to overwrite specific
// ModuleManager APIs.
type StakingModuleBasic struct {
	staking.AppModuleBasic
}

// DefaultGenesis returns custom x/staking module genesis state.
func (StakingModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	genState := stakingtypes.DefaultGenesisState()
	genState.Params.BondDenom = consts.BondDenom

	return cdc.MustMarshalJSON(genState)
}

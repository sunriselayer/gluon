package defaultoverrides

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/cosmos/cosmos-sdk/x/bank"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// BankModuleBasic defines a custom wrapper around the x/bank module's AppModuleBasic
// implementation to provide custom default genesis state.
type BankModuleBasic struct {
	bank.AppModuleBasic
}

// DefaultGenesis returns custom x/bank module genesis state.
func (BankModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	metadataBond := banktypes.Metadata{
		Description: "The native token of Gluon.",
		DenomUnits: []*banktypes.DenomUnit{
			{
				Denom:    "uglu",
				Exponent: 0,
				Aliases:  []string{"microglu"},
			},
			{
				Denom:    "glu",
				Exponent: 6,
			},
		},
		Base:    "uglu",
		Display: "glu",
		Name:    "Gluon",
		Symbol:  "GLU",
	}

	genState := banktypes.DefaultGenesisState()
	genState.DenomMetadata = append(genState.DenomMetadata, metadataBond)

	return cdc.MustMarshalJSON(genState)
}

package defaultoverrides

import (
	"encoding/base64"
	"encoding/json"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"

	customauthmodule "gluon/x/customauth/module"
	customauthtypes "gluon/x/customauth/types"
)

// CustomAuthModuleBasic defines a custom wrapper around the x/customauth module's AppModuleBasic
// implementation to provide custom default genesis state.
type CustomAuthModuleBasic struct {
	customauthmodule.AppModuleBasic
}

// DefaultGenesis returns custom x/tokenconverter module genesis state.
func (m CustomAuthModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	genState := customauthtypes.DefaultGenesis()
	base64Pubkey := "AjYDAuQ764+t7gn9XCcC4LwJQt8JmezVVoIrbokq2ZKn"
	pubkeyBytes, _ := base64.StdEncoding.DecodeString(base64Pubkey)
	pubkey := secp256k1.PubKey{Key: pubkeyBytes}
	pubkeyAny, _ := codectypes.NewAnyWithValue(&pubkey)
	genState.Params.ParingDelay = time.Duration(time.Hour * 24)
	genState.Params.OperatorPublicKey = *pubkeyAny

	return cdc.MustMarshalJSON(genState)
}

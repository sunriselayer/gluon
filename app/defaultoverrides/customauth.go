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
	operatorBase64Pubkey := "AjYDAuQ764+t7gn9XCcC4LwJQt8JmezVVoIrbokq2ZKn"
	operatorPubkeyBytes, _ := base64.StdEncoding.DecodeString(operatorBase64Pubkey)
	operatorPubkey := secp256k1.PubKey{Key: operatorPubkeyBytes}
	operatorPubkeyAny, _ := codectypes.NewAnyWithValue(&operatorPubkey)
	genState.Params.ParingDelay = time.Duration(time.Hour * 24)
	genState.Params.OperatorPublicKey = *operatorPubkeyAny

	validatorBase64Pubkey := "A8FoOqwoBQjTnEPIYg/gSstpIwRsruLj7sz+9raKiLKV"
	validatorPubkeyBytes, _ := base64.StdEncoding.DecodeString(validatorBase64Pubkey)
	validatorPubkey := secp256k1.PubKey{Key: validatorPubkeyBytes}
	validatorPubkeyAny, _ := codectypes.NewAnyWithValue(&validatorPubkey)
	genState.Pairings = []customauthtypes.Pairing{
		customauthtypes.Pairing{
			Owner:     "gluon1wd3g35hec5m6n4qfwaydndzelsuaafzj7ua7qg",
			Index:     "gluon1wd3g35hec5m6n4qfwaydndzelsuaafzj7ua7qg",
			PublicKey: *validatorPubkeyAny,
			CreatedAt: time.Unix(0, 0),
		}}

	return cdc.MustMarshalJSON(genState)
}

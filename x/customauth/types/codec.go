package types

import (
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"

	// this line is used by starport scaffolding # 1

	"gluon/x/customauth/types/operator"
	"gluon/x/customauth/types/pairing"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
)

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*cryptotypes.PubKey)(nil),
		&operator.PubKey{},
		&pairing.PubKey{},
	)

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreatePairing{},
		&MsgDeletePairing{},
	)
	// this line is used by starport scaffolding # 3

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

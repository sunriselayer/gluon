package types

import (
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	// this line is used by starport scaffolding # 1

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"gluon/x/customauth/types/operator"
	"gluon/x/customauth/types/pairing"
)

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*cryptotypes.PubKey)(nil),
		&operator.PubKeyNotVerifiable{},
		&pairing.PubKeyNotVerifiable{},
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

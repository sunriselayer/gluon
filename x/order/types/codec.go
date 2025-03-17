package types

import (
	"cosmossdk.io/core/registry"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	// this line is used by starport scaffolding # 1
)

func RegisterInterfaces(registrar registry.InterfaceRegistrar) {
	registrar.RegisterInterface("gluon.order.OrderBody", (*OrderBody)(nil))

	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgLazyRegisterOrder{},
		&MsgCancelOrder{},
	)
	// this line is used by starport scaffolding # 3

	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)
	msgservice.RegisterMsgServiceDesc(registrar, &_Msg_serviceDesc)
}

package types

import (
	"cosmossdk.io/core/registry"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"

	ordertypes "gluon/x/order/types"
	// this line is used by starport scaffolding # 1
)

func RegisterInterfaces(registrar registry.InterfaceRegistrar) {
	registrar.RegisterImplementations((*ordertypes.OrderBody)(nil),
		&PerpPositionCreateOrder{},
		&PerpPositionCancelOrder{},
	)

	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgMatchOrder{},
	)
	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDepositCrossMargin{},
	)
	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgWithdrawCrossMargin{},
	)
	// this line is used by starport scaffolding # 3

	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)
	msgservice.RegisterMsgServiceDesc(registrar, &_Msg_serviceDesc)
}

package contract

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "gluon/api/gluon/contract"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "OrderAll",
					Use:       "list-order",
					Short:     "List all order",
				},
				{
					RpcMethod:      "Order",
					Use:            "show-order [id]",
					Short:          "Shows a order",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod: "SortedOrderAll",
					Use:       "list-sorted-order",
					Short:     "List all sorted_order",
				},
				{
					RpcMethod:      "SortedOrder",
					Use:            "show-sorted-order [id]",
					Short:          "Shows a sorted_order",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "expiry"}, {ProtoField: "index"}},
				},
				{
					RpcMethod: "LazyContractAll",
					Use:       "list-lazy-settlement",
					Short:     "List all lazy-settlement",
				},
				{
					RpcMethod:      "LazyContract",
					Use:            "show-lazy-settlement [id]",
					Short:          "Shows a lazy-settlement by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}

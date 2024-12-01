package order

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "gluon/api/gluon/order"
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
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
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
				{
					RpcMethod:      "CreateOrder",
					Use:            "create-order [index]",
					Short:          "Create a new order",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod:      "UpdateOrder",
					Use:            "update-order [index]",
					Short:          "Update order",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod:      "DeleteOrder",
					Use:            "delete-order [index]",
					Short:          "Delete order",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod:      "CreateSortedOrder",
					Use:            "create-sorted-order [index]",
					Short:          "Create a new sorted_order",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod:      "UpdateSortedOrder",
					Use:            "update-sorted-order [index]",
					Short:          "Update sorted_order",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod:      "DeleteSortedOrder",
					Use:            "delete-sorted-order [index]",
					Short:          "Delete sorted_order",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}

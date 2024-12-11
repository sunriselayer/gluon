package perp

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "gluon/api/gluon/perp"
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
					RpcMethod:      "Positions",
					Use:            "list-positions [owner]",
					Short:          "List all positions",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "owner"}},
				},
				{
					RpcMethod:      "Position",
					Use:            "show-position [owner] [order_hash]",
					Short:          "Shows a position by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "owner"}, {ProtoField: "order_hash"}},
				},
				{
					RpcMethod: "PositionPriceQuantities",
					Use:       "list-position-price-quantities",
					Short:     "List all position-price-quantities",
				},
				{
					RpcMethod:      "PositionPriceQuantity",
					Use:            "show-position-price-quantity [id]",
					Short:          "Shows a position-price-quantity",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "owner"}, {ProtoField: "position_order_hash"}, {ProtoField: "price"}},
				},
				{
					RpcMethod:      "CrossMargin",
					Use:            "show-cross-margin [owner]",
					Short:          "Shows a cross-margin",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "owner"}},
				},
				{
					RpcMethod: "FundingRates",
					Use:       "list-funding-rates",
					Short:     "List all funding-rates",
				},
				{
					RpcMethod:      "FundingRate",
					Use:            "show-funding-rate [id]",
					Short:          "Shows a funding-rate",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "denom_base"}, {ProtoField: "denom_quote"}, {ProtoField: "seconds"}, {ProtoField: "nanos"}},
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
					RpcMethod:      "DepositCrossMargin",
					Use:            "deposit-cross-margin",
					Short:          "Send a deposit-cross-margin tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},
				{
					RpcMethod:      "WithdrawCrossMargin",
					Use:            "withdraw-cross-margin",
					Short:          "Send a withdraw-cross-margin tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}

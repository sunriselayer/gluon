package customauth

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "gluon/api/gluon/customauth"
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
					RpcMethod:      "Pairings",
					Use:            "list-pairings [owner]",
					Short:          "List all pairings",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "owner"}},
				},
				{
					RpcMethod:      "Pairing",
					Use:            "show-pairing [index]",
					Short:          "Shows a pairing by index (hex string of 20 bytes)",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "owner"}, {ProtoField: "index"}},
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
					RpcMethod:      "CreatePairing",
					Use:            "create-pairing [user] [publicKey]",
					Short:          "Create pairing",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "user"}, {ProtoField: "public_key"}},
				},
				{
					RpcMethod:      "DeletePairing",
					Use:            "delete-pairing [user] [index]",
					Short:          "Delete pairing",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "user"}, {ProtoField: "pairing_index"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}

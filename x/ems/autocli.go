// x/ems/autocli.go
package module

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
	modulev1 "github.com/rollchains/hackmoschain/api/ems/v1"
)

func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Query the current consensus parameters",
				},
				{
					RpcMethod: "GetEvent",
					Use:       "event [id]",
					Short:     "Get event by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "id"},
					},
				},
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Msg_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "CreateEvent",
					Use:       "create [name] [id] [description] [location] [date] [time] [ticket_price] [num_tickets] [ticket_categories]",
					Short:     "Create an event",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "name"},
						{ProtoField: "id"},
						{ProtoField: "description"},
						{ProtoField: "location"},
						{ProtoField: "date"},
						{ProtoField: "time"},
						{ProtoField: "ticket_price"},
						{ProtoField: "num_tickets"},       // Removed ProtoType: "uint64"
						{ProtoField: "ticket_categories"}, // Removed ProtoType: "string", IsRepeated: true
					},
				},
				{
					RpcMethod: "IssueEventNFT",
					Use:       "issue [id] [receiver]",
					Short:     "Issue an event NFT",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "id"},
						{ProtoField: "receiver"},
					},
				},
				{
					RpcMethod: "UpdateParams",
					Skip:      false,
				},
			},
		},
	}
}

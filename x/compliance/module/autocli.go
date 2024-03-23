package compliance

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "compliance/api/compliance/compliance"
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
					RpcMethod:      "ListPendingApprovals",
					Use:            "list-pending-approvals",
					Short:          "Query list-pending-approvals",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},

				{
					RpcMethod:      "ListApprovedApprovals",
					Use:            "list-approved-approvals",
					Short:          "Query list-approved-approvals",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},

				{
					RpcMethod:      "GetApprovalByID",
					Use:            "get-approval-by-id [approval-id]",
					Short:          "Query get-approval-by-id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "approvalId"}},
				},

				{
					RpcMethod:      "ListDriverApprovals",
					Use:            "list-driver-approvals [driver-id]",
					Short:          "Query list-driver-approvals",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "driverId"}},
				},

				{
					RpcMethod:      "ListRejectedApprovals",
					Use:            "list-rejected-approvals",
					Short:          "Query list-rejected-approvals",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},

				{
					RpcMethod:      "ListAllApprovals",
					Use:            "list-all-approvals",
					Short:          "Query list-all-approvals",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
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
					RpcMethod:      "CreateApproval",
					Use:            "create-approval [driver-id] [server-url]",
					Short:          "Send a create-approval tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "driverId"}, {ProtoField: "serverUrl"}},
				},
				{
					RpcMethod:      "UpdateApproval",
					Use:            "update-approval [approval-id] [tester-id] [tester-url] [approved]",
					Short:          "Send a update-approval tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "approvalId"}, {ProtoField: "testerId"}, {ProtoField: "testerUrl"}, {ProtoField: "approved"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}

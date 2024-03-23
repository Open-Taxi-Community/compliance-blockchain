package keeper

import (
	"context"

	"compliance/x/compliance/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateApproval(goCtx context.Context, msg *types.MsgCreateApproval) (*types.MsgCreateApprovalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	approval := types.Approval{
		ServerUrl: msg.ServerUrl,
		DriverID:  msg.DriverId,
		Status:    types.ApprovalStatus_PENDING,
		Creator:   msg.Creator,
		Results:   []*types.ApprovalResult{},
	}
	id := k.AppendApproval(
		ctx,
		approval,
	)

	return &types.MsgCreateApprovalResponse{
		Id: id,
	}, nil
}

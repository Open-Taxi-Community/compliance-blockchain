package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"compliance/x/compliance/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateApproval(goCtx context.Context, msg *types.MsgUpdateApproval) (*types.MsgUpdateApprovalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	approval, found := k.GetApproval(ctx, msg.ApprovalId)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.ApprovalId))
	}
	result := types.ApprovalResult{
		Id:        uint64(len(approval.Results) + 1),
		TesterUrl: msg.TesterUrl,
		TesterID:  msg.TesterId,
		Approved:  msg.Approved,
		Creator:   msg.Creator,
	}
	approval.Results = append(approval.Results, &result)

	for _, res := range approval.Results {
		if !res.Approved {
			approval.Status = types.ApprovalStatus_REJECTED
		}
	}
	// TODO should be more than two but im just testing
	if approval.Status != types.ApprovalStatus_REJECTED && len(approval.Results) >= 2 {
		approval.Status = types.ApprovalStatus_APPROVED
	} else {
		approval.Status = types.ApprovalStatus_PENDING
	}

	k.SetApproval(ctx, approval)

	return &types.MsgUpdateApprovalResponse{Id: approval.Id}, nil
}

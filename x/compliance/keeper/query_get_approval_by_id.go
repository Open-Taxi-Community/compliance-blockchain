package keeper

import (
	"context"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"compliance/x/compliance/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetApprovalByID(goCtx context.Context, req *types.QueryGetApprovalByIDRequest) (*types.QueryGetApprovalByIDResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	approval, found := k.GetApproval(ctx, req.ApprovalId)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetApprovalByIDResponse{Approval: &approval}, nil
}

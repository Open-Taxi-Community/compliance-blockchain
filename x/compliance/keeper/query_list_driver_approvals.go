package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"

	"compliance/x/compliance/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListDriverApprovals(goCtx context.Context, req *types.QueryListDriverApprovalsRequest) (*types.QueryListDriverApprovalsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(goCtx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ApprovalKey))

	var approvals []*types.Approval
	pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		var approval types.Approval
		if err := k.cdc.Unmarshal(value, &approval); err != nil {
			return err
		}
		if approval.DriverID == req.DriverId {
			approvals = append(approvals, &approval)
		}
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryListDriverApprovalsResponse{Approvals: approvals, Pagination: pageRes}, nil
}

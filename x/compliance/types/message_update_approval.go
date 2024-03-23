package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateApproval{}

func NewMsgUpdateApproval(creator string, approvalId uint64, testerId string, testerUrl string, approved bool) *MsgUpdateApproval {
	return &MsgUpdateApproval{
		Creator:    creator,
		ApprovalId: approvalId,
		TesterId:   testerId,
		TesterUrl:  testerUrl,
		Approved:   approved,
	}
}

func (msg *MsgUpdateApproval) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

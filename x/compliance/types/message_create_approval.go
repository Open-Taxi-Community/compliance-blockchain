package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateApproval{}

func NewMsgCreateApproval(creator string, driverId string, serverUrl string) *MsgCreateApproval {
	return &MsgCreateApproval{
		Creator:   creator,
		DriverId:  driverId,
		ServerUrl: serverUrl,
	}
}

func (msg *MsgCreateApproval) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

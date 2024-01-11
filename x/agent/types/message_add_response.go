package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgAddResponse{}

func NewMsgAddResponse(creator string, response string, isfinished bool, runid uint64) *MsgAddResponse {
	return &MsgAddResponse{
		Creator:    creator,
		Response:   response,
		Isfinished: isfinished,
		Runid:      runid,
	}
}

func (msg *MsgAddResponse) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

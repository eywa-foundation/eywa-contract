package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgHandshake = "handshake"

var _ sdk.Msg = &MsgHandshake{}

func NewMsgHandshake(creator string, receiver string, serverAddress string, roomId string, time string) *MsgHandshake {
	return &MsgHandshake{
		Creator:       creator,
		Receiver:      receiver,
		ServerAddress: serverAddress,
		RoomId:        roomId,
		Time:          time,
	}
}

func (msg *MsgHandshake) Route() string {
	return RouterKey
}

func (msg *MsgHandshake) Type() string {
	return TypeMsgHandshake
}

func (msg *MsgHandshake) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgHandshake) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgHandshake) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSendChat = "send_chat"

var _ sdk.Msg = &MsgSendChat{}

func NewMsgSendChat(creator string, chat string, receiver string, time string) *MsgSendChat {
	return &MsgSendChat{
		Creator:  creator,
		Chat:     chat,
		Receiver: receiver,
		Time:     time,
	}
}

func (msg *MsgSendChat) Route() string {
	return RouterKey
}

func (msg *MsgSendChat) Type() string {
	return TypeMsgSendChat
}

func (msg *MsgSendChat) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSendChat) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendChat) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

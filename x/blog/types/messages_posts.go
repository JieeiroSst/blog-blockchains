package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreatePosts{}

func NewMsgCreatePosts(creator string, title string, body string) *MsgCreatePosts {
	return &MsgCreatePosts{
		Creator: creator,
		Title:   title,
		Body:    body,
	}
}

func (msg *MsgCreatePosts) Route() string {
	return RouterKey
}

func (msg *MsgCreatePosts) Type() string {
	return "CreatePosts"
}

func (msg *MsgCreatePosts) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreatePosts) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreatePosts) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdatePosts{}

func NewMsgUpdatePosts(creator string, id string, title string, body string) *MsgUpdatePosts {
	return &MsgUpdatePosts{
		Id:      id,
		Creator: creator,
		Title:   title,
		Body:    body,
	}
}

func (msg *MsgUpdatePosts) Route() string {
	return RouterKey
}

func (msg *MsgUpdatePosts) Type() string {
	return "UpdatePosts"
}

func (msg *MsgUpdatePosts) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdatePosts) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdatePosts) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgCreatePosts{}

func NewMsgDeletePosts(creator string, id string) *MsgDeletePosts {
	return &MsgDeletePosts{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeletePosts) Route() string {
	return RouterKey
}

func (msg *MsgDeletePosts) Type() string {
	return "DeletePosts"
}

func (msg *MsgDeletePosts) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeletePosts) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeletePosts) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetAdd{}

type MsgSetAdd struct {
  ID      string      `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  News string `json:"news" yaml:"news"`
  Title string `json:"title" yaml:"title"`
  Body string `json:"body" yaml:"body"`
  Link string `json:"link" yaml:"link"`
}

func NewMsgSetAdd(creator sdk.AccAddress, id string, news string, title string, body string, link string) MsgSetAdd {
  return MsgSetAdd{
    ID: id,
		Creator: creator,
    News: news,
    Title: title,
    Body: body,
    Link: link,
	}
}

func (msg MsgSetAdd) Route() string {
  return RouterKey
}

func (msg MsgSetAdd) Type() string {
  return "SetAdd"
}

func (msg MsgSetAdd) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgSetAdd) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgSetAdd) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}
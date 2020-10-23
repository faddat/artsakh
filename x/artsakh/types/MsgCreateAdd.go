package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateAdd{}

type MsgCreateAdd struct {
  ID      string
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  News string `json:"news" yaml:"news"`
  Title string `json:"title" yaml:"title"`
  Body string `json:"body" yaml:"body"`
  Link string `json:"link" yaml:"link"`
}

func NewMsgCreateAdd(creator sdk.AccAddress, news string, title string, body string, link string) MsgCreateAdd {
  return MsgCreateAdd{
    ID: uuid.New().String(),
		Creator: creator,
    News: news,
    Title: title,
    Body: body,
    Link: link,
	}
}

func (msg MsgCreateAdd) Route() string {
  return RouterKey
}

func (msg MsgCreateAdd) Type() string {
  return "CreateAdd"
}

func (msg MsgCreateAdd) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateAdd) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgCreateAdd) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}
package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeleteAdd{}

type MsgDeleteAdd struct {
  ID      string         `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
}

func NewMsgDeleteAdd(id string, creator sdk.AccAddress) MsgDeleteAdd {
  return MsgDeleteAdd{
    ID: id,
		Creator: creator,
	}
}

func (msg MsgDeleteAdd) Route() string {
  return RouterKey
}

func (msg MsgDeleteAdd) Type() string {
  return "DeleteAdd"
}

func (msg MsgDeleteAdd) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgDeleteAdd) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgDeleteAdd) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}
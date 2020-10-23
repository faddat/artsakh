package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Add struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
    News string `json:"news" yaml:"news"`
    Title string `json:"title" yaml:"title"`
    Body string `json:"body" yaml:"body"`
    Link string `json:"link" yaml:"link"`
}
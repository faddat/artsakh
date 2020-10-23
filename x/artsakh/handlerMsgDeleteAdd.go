package artsakh

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/faddat/artsakh/x/artsakh/types"
	"github.com/faddat/artsakh/x/artsakh/keeper"
)

// Handle a message to delete name
func handleMsgDeleteAdd(ctx sdk.Context, k keeper.Keeper, msg types.MsgDeleteAdd) (*sdk.Result, error) {
	if !k.AddExists(ctx, msg.ID) {
		// replace with ErrKeyNotFound for 0.39+
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.ID)
	}
	if !msg.Creator.Equals(k.GetAddOwner(ctx, msg.ID)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	k.DeleteAdd(ctx, msg.ID)
	return &sdk.Result{}, nil
}

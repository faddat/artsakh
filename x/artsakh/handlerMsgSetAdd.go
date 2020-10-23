package artsakh

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/faddat/artsakh/x/artsakh/types"
	"github.com/faddat/artsakh/x/artsakh/keeper"
)

func handleMsgSetAdd(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetAdd) (*sdk.Result, error) {
	var add = types.Add{
		Creator: msg.Creator,
		ID:      msg.ID,
    	News: msg.News,
    	Title: msg.Title,
    	Body: msg.Body,
    	Link: msg.Link,
	}
	if !msg.Creator.Equals(k.GetAddOwner(ctx, msg.ID)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	k.SetAdd(ctx, add)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

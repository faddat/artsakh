package artsakh

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/faddat/artsakh/x/artsakh/types"
	"github.com/faddat/artsakh/x/artsakh/keeper"
)

func handleMsgCreateAdd(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateAdd) (*sdk.Result, error) {
	var add = types.Add{
		Creator: msg.Creator,
		ID:      msg.ID,
    	News: msg.News,
    	Title: msg.Title,
    	Body: msg.Body,
    	Link: msg.Link,
	}
	k.CreateAdd(ctx, add)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

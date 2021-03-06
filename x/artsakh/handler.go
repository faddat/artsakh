package artsakh

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/faddat/artsakh/x/artsakh/keeper"
	"github.com/faddat/artsakh/x/artsakh/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
    // this line is used by starport scaffolding # 1
		case types.MsgCreateAdd:
			return handleMsgCreateAdd(ctx, k, msg)
		case types.MsgSetAdd:
			return handleMsgSetAdd(ctx, k, msg)
		case types.MsgDeleteAdd:
			return handleMsgDeleteAdd(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

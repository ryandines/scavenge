package scavenge

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/starport/scavenge/x/scavenge/keeper"
	"github.com/starport/scavenge/x/scavenge/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
    // this line is used by starport scaffolding # 1
		case types.MsgCreateCommit:
			return handleMsgCreateCommit(ctx, k, msg)
		case types.MsgSetCommit:
			return handleMsgSetCommit(ctx, k, msg)
		case types.MsgDeleteCommit:
			return handleMsgDeleteCommit(ctx, k, msg)
		case types.MsgCreateScavenge:
			return handleMsgCreateScavenge(ctx, k, msg)
		case types.MsgSetScavenge:
			return handleMsgSetScavenge(ctx, k, msg)
		case types.MsgDeleteScavenge:
			return handleMsgDeleteScavenge(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

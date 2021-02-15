package scavenge

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/starport/scavenge/x/scavenge/types"
	"github.com/starport/scavenge/x/scavenge/keeper"
)

func handleMsgCreateCommit(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateCommit) (*sdk.Result, error) {
	k.CreateCommit(ctx, msg)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

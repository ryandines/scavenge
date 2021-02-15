package scavenge

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/starport/scavenge/x/scavenge/types"
	"github.com/starport/scavenge/x/scavenge/keeper"
)

func handleMsgCreateScavenge(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateScavenge) (*sdk.Result, error) {
	k.CreateScavenge(ctx, msg)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

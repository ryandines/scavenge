package scavenge

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/starport/scavenge/x/scavenge/types"
	"github.com/starport/scavenge/x/scavenge/keeper"
)

func handleMsgSetCommit(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetCommit) (*sdk.Result, error) {
	var commit = types.Commit{
		Creator: msg.Creator,
		ID:      msg.ID,
    	SolutionHash: msg.SolutionHash,
    	SolutionScavengerHash: msg.SolutionScavengerHash,
	}
	if !msg.Creator.Equals(k.GetCommitOwner(ctx, msg.ID)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	k.SetCommit(ctx, commit)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

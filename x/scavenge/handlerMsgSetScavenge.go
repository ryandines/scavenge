package scavenge

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/starport/scavenge/x/scavenge/types"
	"github.com/starport/scavenge/x/scavenge/keeper"
)

func handleMsgSetScavenge(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetScavenge) (*sdk.Result, error) {
	var scavenge = types.Scavenge{
		Creator: msg.Creator,
		ID:      msg.ID,
    	Description: msg.Description,
    	SolutionHash: msg.SolutionHash,
    	Reward: msg.Reward,
    	Solution: msg.Solution,
    	Scavenger: msg.Scavenger,
	}
	if !msg.Creator.Equals(k.GetScavengeOwner(ctx, msg.ID)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	k.SetScavenge(ctx, scavenge)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

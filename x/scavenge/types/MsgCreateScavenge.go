package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateScavenge{}

// MsgCreateScavenge is a type that creates scavenge messages
type MsgCreateScavenge struct {
	Creator      sdk.AccAddress `json:"creator" yaml:"creator"`
	Description  string         `json:"description" yaml:"description"`
	SolutionHash string         `json:"solutionHash" yaml:"solutionHash"`
	Reward       sdk.Coins      `json:"reward" yaml:"reward"`
}

// NewMsgCreateScavenge creates scavenge messages
func NewMsgCreateScavenge(creator sdk.AccAddress, description string, solutionHash string, reward sdk.Coins) MsgCreateScavenge {
	return MsgCreateScavenge{
		Creator:      creator,
		Description:  description,
		SolutionHash: solutionHash,
		Reward:       reward,
	}
}

// Route is for routing
func (msg MsgCreateScavenge) Route() string {
	return RouterKey
}

// Type is for typing
func (msg MsgCreateScavenge) Type() string {
	return "CreateScavenge"
}

// GetSigners is for signing
func (msg MsgCreateScavenge) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

// GetSignBytes is for signing bytes
func (msg MsgCreateScavenge) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic is for validation
func (msg MsgCreateScavenge) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	if msg.SolutionHash == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "solutionHash can't be empty")
	}
	return nil
}

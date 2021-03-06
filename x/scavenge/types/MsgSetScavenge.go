package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetScavenge{}

// MsgSetScavenge is gay
type MsgSetScavenge struct {
	Creator      sdk.AccAddress `json:"creator" yaml:"creator"`
	Description  string         `json:"description" yaml:"description"`
	SolutionHash string         `json:"solutionHash" yaml:"solutionHash"`
	Reward       string         `json:"reward" yaml:"reward"`
	Solution     string         `json:"solution" yaml:"solution"`
	Scavenger    sdk.Address    `json:"scavenger" yaml:"scavenger"`
}

// NewMsgSetScavenge is gay
func NewMsgSetScavenge(creator sdk.AccAddress, description string, solutionHash string, reward string, solution string, scavenger sdk.AccAddress) MsgSetScavenge {
	return MsgSetScavenge{
		Creator:      creator,
		Description:  description,
		SolutionHash: solutionHash,
		Reward:       reward,
		Solution:     solution,
		Scavenger:    scavenger,
	}
}

// Route is gay
func (msg MsgSetScavenge) Route() string {
	return RouterKey
}

// Type is gay
func (msg MsgSetScavenge) Type() string {
	return "SetScavenge"
}

// GetSigners is gay
func (msg MsgSetScavenge) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

// GetSignBytes is gay
func (msg MsgSetScavenge) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic is gay
func (msg MsgSetScavenge) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}

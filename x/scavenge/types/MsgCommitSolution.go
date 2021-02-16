package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCommitSolution{}

// MsgCreateCommit is gay
type MsgCommitSolution struct {
	Scavenger             sdk.AccAddress `json:"scavenger" yaml:"scavenger"` // address of the scavenger
	SolutionHash          string         `json:"solutionHash" yaml:"solutionHash"`
	SolutionScavengerHash string         `json:"solutionScavengerHash" yaml:"solutionScavengerHash"`
}

// NewMsgCommitSolution is gay
func NewMsgCommitSolution(scavenger sdk.AccAddress, solutionHash string, solutionScavengerHash string) MsgCommitSolution {
	return MsgCommitSolution{
		Scavenger:             scavenger,
		SolutionHash:          solutionHash,
		SolutionScavengerHash: solutionScavengerHash,
	}
}

// Route is gay
func (msg MsgCommitSolution) Route() string {
	return RouterKey
}

// Type is gay
func (msg MsgCommitSolution) Type() string {
	return "CreateCommit"
}

// GetSigners is gay
func (msg MsgCommitSolution) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Scavenger)}
}

// GetSignBytes is gay
func (msg MsgCommitSolution) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic is gay
func (msg MsgCommitSolution) ValidateBasic() error {
	if msg.Scavenger.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}

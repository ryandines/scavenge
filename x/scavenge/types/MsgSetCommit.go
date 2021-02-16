package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetCommit{}

// MsgSetCommit implements IMessage
type MsgSetCommit struct {
	Scavenger             sdk.AccAddress `json:"scavenger" yaml:"scavenger"`
	SolutionHash          string         `json:"solutionHash" yaml:"solutionHash"`
	SolutionScavengerHash string         `json:"solutionScavengerHash" yaml:"solutionScavengerHash"`
}

// NewMsgSetCommit implements IMessage
func NewMsgSetCommit(scavenger sdk.AccAddress, id string, solutionHash string, solutionScavengerHash string) MsgSetCommit {
	return MsgSetCommit{
		Scavenger:             scavenger,
		SolutionHash:          solutionHash,
		SolutionScavengerHash: solutionScavengerHash,
	}
}

// Route implements IMessage
func (msg MsgSetCommit) Route() string {
	return RouterKey
}

// Type implements IMessage
func (msg MsgSetCommit) Type() string {
	return "SetCommit"
}

// GetSigners implements IMessage
func (msg MsgSetCommit) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Scavenger)}
}

// GetSignBytes implements IMessage
func (msg MsgSetCommit) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements IMessage
func (msg MsgSetCommit) ValidateBasic() error {
	if msg.Scavenger.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}

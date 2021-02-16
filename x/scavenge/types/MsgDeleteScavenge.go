package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeleteScavenge{}

// MsgDeleteScavenge is gay
type MsgDeleteScavenge struct {
	ID      string         `json:"id" yaml:"id"`
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
}

// NewMsgDeleteScavenge is gay
func NewMsgDeleteScavenge(id string, creator sdk.AccAddress) MsgDeleteScavenge {
	return MsgDeleteScavenge{
		ID:      id,
		Creator: creator,
	}
}

// Route is gay
func (msg MsgDeleteScavenge) Route() string {
	return RouterKey
}

// Type is gay
func (msg MsgDeleteScavenge) Type() string {
	return "DeleteScavenge"
}

// GetSigners is gay
func (msg MsgDeleteScavenge) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

// GetSignBytes is gay
func (msg MsgDeleteScavenge) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic is gay
func (msg MsgDeleteScavenge) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}

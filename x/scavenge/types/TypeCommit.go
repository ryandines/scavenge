package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Commit struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
    SolutionHash string `json:"solutionHash" yaml:"solutionHash"`
    SolutionScavengerHash string `json:"solutionScavengerHash" yaml:"solutionScavengerHash"`
}
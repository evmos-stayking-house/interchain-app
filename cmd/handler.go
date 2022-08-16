package cmd

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// HandleDelegation unwraps the received EVMOS and delegates it to earn staking rewards.
func HandleDelegation(change DelegationChange) error {
	acc := sdk.AccAddress{}
	acc.String()
	return nil
}

func HandleUndelegation(change DelegationChange) error {
	return nil
}

package cmd

import (
	flag "github.com/spf13/pflag"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// HandleDelegation unwraps the received EVMOS and delegates it to earn staking rewards.
func HandleDelegation(ctx client.Context, flgs *flag.FlagSet, change DelegationChange) error {
	denom := "aevmos"
	amt := sdk.NewIntFromBigInt(change.Amount)
	coin := sdk.NewCoin(denom, amt)

	delAddr := ctx.GetFromAddress()
	valString, _ := flgs.GetString(flagValidator)
	valAddr, err := sdk.ValAddressFromBech32(valString)
	if err != nil {
		return err
	}

	msg := stakingtypes.NewMsgDelegate(delAddr, valAddr, coin)

	return tx.GenerateOrBroadcastTxCLI(ctx, flgs, msg)
}

package cmd

import (
	"context"
	"fmt"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	flag "github.com/spf13/pflag"
	"github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
)

// HandleDelegation unwraps the received EVMOS and delegates it to earn staking rewards.
func HandleDelegation(cliCtx client.Context, flgs *flag.FlagSet, change DelegationChange) error {
	amt := sdk.NewIntFromBigInt(change.Amount)
	coin := sdk.NewCoin(baseDenom, amt)

	delAddr := cliCtx.GetFromAddress()
	valString, _ := flgs.GetString(flagValidator)
	valAddr, err := sdk.ValAddressFromBech32(valString)
	if err != nil {
		return err
	}

	msg := stakingtypes.NewMsgDelegate(delAddr, valAddr, coin)

	return tx.GenerateOrBroadcastTxCLI(cliCtx, flgs, msg)
}

// HandleEpochEnd handles epoch ending by delegating the newly received coins
func HandleEpochEnd(cliCtx client.Context, flgs *flag.FlagSet) error {
	queryClient := distrtypes.NewQueryClient(cliCtx)

	// get delegator address & validator address
	delAddr := cliCtx.GetFromAddress()
	valString, err := flgs.GetString(flagValidator)
	if err != nil {
		return err
	}
	valAddr, err := sdk.ValAddressFromBech32(valString)
	if err != nil {
		return err
	}

	// construct delegation rewards query request
	params := &distrtypes.QueryDelegationRewardsRequest{
		DelegatorAddress: delAddr.String(),
		ValidatorAddress: valAddr.String(),
	}

	res, err := queryClient.DelegationRewards(context.Background(), params)
	if err != nil {
		return err
	}
	fmt.Println(res)

	stakingQueryClient := stakingtypes.NewQueryClient(cliCtx)
	stakingParamsResponse, err := stakingQueryClient.Params(context.Background(), &stakingtypes.QueryParamsRequest{})
	if err != nil {
		return err
	}
	bondDenom := stakingParamsResponse.Params.BondDenom

	toCompound, _ := res.GetRewards().TruncateDecimal()
	delegateCoin := sdk.NewCoin(bondDenom, toCompound.AmountOf(bondDenom))

	// construct tx with withdraw & delegate commands
	withdrawMsg := distrtypes.NewMsgWithdrawDelegatorReward(delAddr, valAddr)
	msgs := []sdk.Msg{withdrawMsg}
	delegateMsg := stakingtypes.NewMsgDelegate(delAddr, valAddr, delegateCoin)
	msgs = append(msgs, delegateMsg)
	return tx.GenerateOrBroadcastTxCLI(cliCtx, flgs, msgs...)
}

// HandleUndelegateComplete handles completed unbondings by sending the unlocked coins to unbonding contract
func HandleUndelegateComplete(ctx client.Context, flgs *flag.FlagSet, event types.Event) interface{} {
	return nil
}

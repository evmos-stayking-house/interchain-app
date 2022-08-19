package cmd

import (
	"context"
	"fmt"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	flag "github.com/spf13/pflag"

	abcitypes "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
)

// HandleDelegation unwraps the received EVMOS and delegates it to earn staking rewards.
func HandleDelegation(cliCtx client.Context, flgs *flag.FlagSet, change DelegationChange) error {
	denom := "aevmos"
	amt := sdk.NewIntFromBigInt(change.Amount)
	coin := sdk.NewCoin(denom, amt)

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
func HandleEpochEnd(cliCtx client.Context, flgs *flag.FlagSet, event abcitypes.Event) error {
	queryClient := distrtypes.NewQueryClient(cliCtx)

	delAddr := cliCtx.GetFromAddress()

	valString, err := flgs.GetString(flagValidator)
	if err != nil {
		return err
	}

	valAddr, err := sdk.ValAddressFromBech32(valString)
	if err != nil {
		return err
	}
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

	withdrawMsg := distrtypes.NewMsgWithdrawDelegatorReward(delAddr, valAddr)
	msgs := []sdk.Msg{withdrawMsg}
	delegateMsg := stakingtypes.NewMsgDelegate(delAddr, valAddr, delegateCoin)
	msgs = append(msgs, delegateMsg)
	return tx.GenerateOrBroadcastTxCLI(cliCtx, flgs, msgs...)
}

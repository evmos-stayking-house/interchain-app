package cmd

import (
	"context"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	flag "github.com/spf13/pflag"
)

func getBondDenom(cliCtx client.Context) (string, error) {
	stakingQueryClient := stakingtypes.NewQueryClient(cliCtx)
	stakingParamsResponse, err := stakingQueryClient.Params(context.Background(), &stakingtypes.QueryParamsRequest{})
	if err != nil {
		return "", err
	}
	bondDenom := stakingParamsResponse.Params.BondDenom
	return bondDenom, nil
}

func getStakingRewards(cliCtx client.Context, flgs *flag.FlagSet) (sdk.Coins, error) {
	distrQC := distrtypes.NewQueryClient(cliCtx)

	// get delegator address & validator address
	delAddr := cliCtx.GetFromAddress()
	valString, err := flgs.GetString(flagValidator)
	if err != nil {
		return nil, err
	}
	valAddr, err := sdk.ValAddressFromBech32(valString)
	if err != nil {
		return nil, err
	}

	// construct delegation rewards query request
	params := &distrtypes.QueryDelegationRewardsRequest{
		DelegatorAddress: delAddr.String(),
		ValidatorAddress: valAddr.String(),
	}
	resp, err := distrQC.DelegationRewards(context.Background(), params)
	if err != nil {
		return nil, err
	}
	res, _ := resp.GetRewards().TruncateDecimal()
	return res, nil
}

func getAccountBalances(cliCtx client.Context, flgs *flag.FlagSet) (sdk.Coins, error) {
	fromAddr := cliCtx.GetFromAddress()
	bankQC := banktypes.NewQueryClient(cliCtx)
	req := banktypes.QueryAllBalancesRequest{
		Address: fromAddr.String(),
	}
	resp, err := bankQC.AllBalances(context.Background(), &req)
	if err != nil {
		return sdk.Coins{}, err
	}
	return resp.Balances, nil
}

func getTotalUnbonding(cliCtx client.Context, flgs *flag.FlagSet) (sdk.Coins, error) {
	// TODO: handle mutliple validators
	// assume only 1 validator for now
	// get delegator address & validator address
	delAddr := cliCtx.GetFromAddress()
	valString, err := flgs.GetString(flagValidator)
	if err != nil {
		return nil, err
	}
	valAddr, err := sdk.ValAddressFromBech32(valString)
	if err != nil {
		return nil, err
	}

	// query unbonding entries for the validator(s)
	stakingQC := stakingtypes.NewQueryClient(cliCtx)
	req := stakingtypes.QueryUnbondingDelegationRequest{
		DelegatorAddr: delAddr.String(),
		ValidatorAddr: valAddr.String(),
	}
	resp, err := stakingQC.UnbondingDelegation(context.Background(), &req)
	if err != nil {
		return sdk.Coins{}, err
	}
	entries := resp.Unbond.Entries
	total := sdk.ZeroInt()
	for _, e := range entries {
		total = total.Add(e.Balance)
	}

	// construct coins
	bondDenom, err := getBondDenom(cliCtx)
	if err != nil {
		return sdk.Coins{}, err
	}
	res := sdk.NewCoins(sdk.NewCoin(bondDenom, total))

	return res, nil
}

func getTotalBonded(cliCtx client.Context, flgs *flag.FlagSet) (sdk.Coins, error) {
	// TODO: handle mutliple validators
	// assume only 1 validator for now
	// get delegator address & validator address
	delAddr := cliCtx.GetFromAddress()
	valString, err := flgs.GetString(flagValidator)
	if err != nil {
		return nil, err
	}
	valAddr, err := sdk.ValAddressFromBech32(valString)
	if err != nil {
		return nil, err
	}

	// query unbonding entries for the validator(s)
	stakingQC := stakingtypes.NewQueryClient(cliCtx)
	req := stakingtypes.QueryDelegationRequest{
		DelegatorAddr: delAddr.String(),
		ValidatorAddr: valAddr.String(),
	}
	resp, err := stakingQC.Delegation(context.Background(), &req)
	if err != nil {
		return sdk.Coins{}, err
	}
	res := sdk.NewCoins(resp.DelegationResponse.Balance)
	return res, nil
}

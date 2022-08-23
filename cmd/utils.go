package cmd

import (
	"context"
	"github.com/cosmos/cosmos-sdk/client"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
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

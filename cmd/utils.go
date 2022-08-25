package cmd

import (
	"context"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	ethermint "github.com/evmos/ethermint/types"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
	flag "github.com/spf13/pflag"
	"log"
	"math/big"
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

// ConstructEthTx constructs ethereum transaction from the context
func ConstructEthTx(cliCtx client.Context, flgs *flag.FlagSet, contractAddr string, value *big.Int, inputData []byte) (signing.Tx, error) {
	//inputData, err := contractABI.Pack(name, args...)
	//if err != nil {
	//	return nil, err
	//}
	ethEndPoint, err := flgs.GetString(flagEthEndpoint)
	if err != nil {
		return nil, err
	}

	ethClient, err := ethclient.Dial(ethEndPoint)
	if err != nil {
		return nil, err
	}

	fromAddr := cliCtx.GetFromAddress()
	fromAddrEth := common.BytesToAddress(fromAddr.Bytes())

	contAddr := common.BytesToAddress(common.Hex2Bytes(contractAddr))
	if err != nil {
		return nil, err
	}

	gas, err := ethClient.EstimateGas(context.Background(), ethereum.CallMsg{
		From:      fromAddrEth,
		To:        &contAddr,
		Gas:       0,
		GasPrice:  nil,
		GasFeeCap: nil,
		GasTipCap: nil,
		Value:     value,
		Data:      inputData,
	})
	if err != nil {
		return nil, err
	}

	gasPriceStr, err := flgs.GetString(flags.FlagGasPrices)
	if err != nil {
		return nil, err
	}
	gasPriceCoin, err := sdk.ParseCoinNormalized(gasPriceStr)
	if err != nil {
		return nil, err
	}
	gasPrice := gasPriceCoin.Amount.BigInt()

	evmChainID, err := ethermint.ParseChainID(cliCtx.ChainID)
	if err != nil {
		return nil, err
	}
	nonce, err := ethClient.NonceAt(context.Background(), fromAddrEth, nil)
	if err != nil {
		return nil, err
	}

	txMsg := evmtypes.NewTx(evmChainID, nonce, &contAddr, value, gas, gasPrice, nil, nil, inputData, nil)
	signer := ethtypes.NewEIP155Signer(evmChainID)
	txMsg.From = fromAddrEth.String()
	err = txMsg.Sign(signer, cliCtx.Keyring)
	if err != nil {
		return nil, err
	}

	rsp, err := evmtypes.NewQueryClient(cliCtx).Params(context.Background(), &evmtypes.QueryParamsRequest{})
	if err != nil {
		return nil, err
	}
	tx, err := txMsg.BuildTx(cliCtx.TxConfig.NewTxBuilder(), rsp.Params.EvmDenom)
	if err != nil {
		return nil, err
	}

	txBytes, err := cliCtx.TxConfig.TxEncoder()(tx)
	if err != nil {
		return nil, err
	}

	// broadcast to a Tendermint node right away, since ethereum Txs have to be in its own tx
	res, err := cliCtx.BroadcastTxSync(txBytes)
	if err != nil {
		return nil, err
	}
	log.Println(res)

	return tx, nil
}

package cmd

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/evmos-stayking-house/interchain-app/abis/stayking"
	ethermint "github.com/evmos/ethermint/types"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
	flag "github.com/spf13/pflag"
	"log"
	"math/big"
	"time"
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

	// construct delegation rewards query request

	params := &distrtypes.QueryDelegationTotalRewardsRequest{
		DelegatorAddress: delAddr.String(),
	}
	resp, err := distrQC.DelegationTotalRewards(context.Background(), params)
	if err != nil {
		return sdk.Coins{}, nil
	}

	res := sdk.NewCoins()
	for _, r := range resp.Rewards {
		c, _ := r.Reward.TruncateDecimal()
		res = res.Add(c...)
	}
	return res, nil
}

func getStakingParams(cliCtx client.Context, flgs *flag.FlagSet) (stakingtypes.Params, error) {
	stakingQC := stakingtypes.NewQueryClient(cliCtx)

	// construct delegation rewards query request
	params := &stakingtypes.QueryParamsRequest{}
	resp, err := stakingQC.Params(context.Background(), params)
	if err != nil {
		return stakingtypes.Params{}, err
	}
	return resp.GetParams(), nil
}

func getAccountBalances(cliCtx client.Context, flgs *flag.FlagSet) (sdk.Coins, error) {
	fromAddr := cliCtx.GetFromAddress()
	bankQC := banktypes.NewQueryClient(cliCtx)
	req := banktypes.QueryAllBalancesRequest{
		Address: fromAddr.String(),
	}
	resp, err := bankQC.AllBalances(context.Background(), &req)
	if err != nil {
		return nil, err
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
		return nil, err
	}
	entries := resp.Unbond.Entries
	total := sdk.ZeroInt()
	for _, e := range entries {
		total = total.Add(e.Balance)
	}

	// construct coins
	bondDenom, err := getBondDenom(cliCtx)
	if err != nil {
		return nil, err
	}
	res := sdk.NewCoins(sdk.NewCoin(bondDenom, total))

	return res, nil
}

func getTotalBonded(cliCtx client.Context, flgs *flag.FlagSet) sdk.Coins {
	delAddr := cliCtx.GetFromAddress()

	// query unbonding entries for the validator(s)
	stakingQC := stakingtypes.NewQueryClient(cliCtx)
	req := stakingtypes.QueryDelegatorDelegationsRequest{
		DelegatorAddr: delAddr.String(),
	}
	resp, err := stakingQC.DelegatorDelegations(context.Background(), &req)
	if err != nil {
		return sdk.Coins{}
	}
	totalCoins := sdk.NewCoins()
	for _, d := range resp.DelegationResponses {
		totalCoins = totalCoins.Add(d.Balance)
	}
	return totalCoins
}

// ConstructEthTx constructs ethereum transaction from the context
func ConstructEthTx(cliCtx client.Context, flgs *flag.FlagSet, contractAddr string, value *big.Int, inputData []byte) (signing.Tx, error) {
	ethEndPoint, err := flgs.GetString(flagEthEndpoint)
	if err != nil {
		return nil, err
	}

	ethClient, err := ethclient.Dial(ethEndPoint)
	if err != nil {
		return nil, err
	}

	fromAddr := cliCtx.GetFromAddress()
	fromAddrEth := common.BytesToAddress(fromAddr)

	contAddr := common.HexToAddress(contractAddr)
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

	gasPriceUint64, err := flgs.GetUint64(flagEthGasPrice)
	if err != nil {
		return nil, err
	}
	gasPrice := sdk.NewIntFromUint64(gasPriceUint64).BigInt()

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
	numRetry, err := flgs.GetUint64(flagMaxRetry)
	if err != nil {
		log.Fatal(err)
	}

	res, err := TrySubmitEthTxMaxRetry(numRetry, cliCtx, txBytes)
	if err != nil {
		return nil, err
	}
	log.Printf("successfully broadcasted EthTx: %s\n", res.TxHash)

	return tx, nil
}

func TrySubmitTxMaxRetry(numRetry uint64, cliCtx client.Context, flgs *flag.FlagSet, msgs ...sdk.Msg) (err error) {
	var i uint64
	for i = 0; i < numRetry; i++ {
		time.Sleep(time.Second * 1)
		if err = tx.GenerateOrBroadcastTxCLI(cliCtx, flgs, msgs...); err == nil {
			return
		}
		log.Printf("failed broadcasting tx. Retrying... %d\n", i)
	}
	return err
}

func TrySubmitEthTxMaxRetry(numRetry uint64, cliCtx client.Context, txBytes []byte) (res *sdk.TxResponse, err error) {
	var i uint64
	for i = 0; i < numRetry; i++ {
		time.Sleep(time.Second * 1)
		if res, err = cliCtx.BroadcastTxCommit(txBytes); err == nil {
			return
		}
		log.Printf("failed broadcasting tx. Retrying... %d\n", i)
	}
	return
}

func withdrawRewards(cliCtx client.Context, flgs *flag.FlagSet) error {
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

	// fire the withdraw rewards tx first so it's ready so we can send the rewards to the stayking contract.
	withdrawMsg := distrtypes.NewMsgWithdrawDelegatorReward(delAddr, valAddr)
	msgs := []sdk.Msg{withdrawMsg}
	numRetry, err := flgs.GetUint64(flagMaxRetry)
	if err != nil {
		log.Fatal(err)
	}
	err = TrySubmitTxMaxRetry(numRetry, cliCtx, flgs, msgs...)
	if err != nil {
		// just log the error since there might not be any ongoing staking
		log.Println(err)
	}
	return nil
}

// amt is the *big.Int of the total interest (rewards + leftover asset)
func accrueAsset(cliCtx client.Context, flgs *flag.FlagSet, amt *big.Int) error {
	totalStaked := GetStakedAsset(cliCtx, flgs)
	totalStakedInt := totalStaked.AmountOf(baseDenom)
	args := totalStakedInt.BigInt()

	contractABI, err := stayking.StaykingMetaData.GetAbi()
	if err != nil {
		fmt.Println("1")
		return err
	}
	inputData, err := contractABI.Pack(accrueFuncName, args)
	if err != nil {
		fmt.Println("2")
		return err
	}

	contractAddr, err := flgs.GetString(flagStaykingContAddr)
	if err != nil {
		fmt.Println("3")
		return err
	}
	_, err = ConstructEthTx(cliCtx, flgs, contractAddr, amt, inputData)
	fmt.Println("4")
	return err
}

// deprecated in favor of calling accrue function directly.
func getDistributionAmt(cliCtx client.Context, flgs *flag.FlagSet, amt sdk.Int) (*big.Int, error) {
	ethEndPoint, err := flgs.GetString(flagEthEndpoint)
	if err != nil {
		return nil, err
	}
	ethClient, err := ethclient.Dial(ethEndPoint)
	if err != nil {
		return nil, err
	}
	contAddrStr, err := flgs.GetString(flagStaykingContAddr)
	if err != nil {
		return nil, err
	}
	contAddr := common.HexToAddress(contAddrStr)
	contractABI, err := stayking.StaykingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	inputData, err := contractABI.Pack(getAccruedValueFuncName, amt.BigInt())
	callMsg := ethereum.CallMsg{
		From:      common.Address{},
		To:        &contAddr,
		Gas:       0,
		GasPrice:  nil,
		GasFeeCap: nil,
		GasTipCap: nil,
		Value:     nil,
		Data:      inputData,
	}

	res, err := ethClient.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		return nil, err
	}
	unpacked, err := contractABI.Unpack(getAccruedValueFuncName, res)
	if err != nil {
		return nil, err
	}
	return unpacked[0].(*big.Int), nil
}

package cmd

import (
	"context"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/evmos-stayking-house/scheduled-worker-golang/abis"
	"github.com/evmos-stayking-house/scheduled-worker-golang/abis/stayking"
	"github.com/evmos-stayking-house/scheduled-worker-golang/types"
	flag "github.com/spf13/pflag"
	tmtypes "github.com/tendermint/tendermint/abci/types"
	"log"
	"math/big"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
)

// GetMsgDelegation unwraps the received EVMOS and delegates it to earn staking rewards.
func GetMsgDelegation(cliCtx client.Context, flgs *flag.FlagSet, change *big.Int) (sdk.Msg, error) {
	amt := sdk.NewIntFromBigInt(change)
	coin := sdk.NewCoin(baseDenom, amt)

	delAddr := cliCtx.GetFromAddress()
	valString, _ := flgs.GetString(flagValidator)
	valAddr, err := sdk.ValAddressFromBech32(valString)
	if err != nil {
		return nil, err
	}
	return stakingtypes.NewMsgDelegate(delAddr, valAddr, coin), nil
}

func GetMsgNewBlockEvents(cliCtx client.Context, flgs *flag.FlagSet, events []tmtypes.Event, store *types.Storage) ([]sdk.Msg, error) {
	msgs := []sdk.Msg{}
	for _, event := range events {
		if event.Type == "epoch_end" {
			newMsgs, err := GetMsgsEpochEnd(cliCtx, flgs)
			if err != nil {
				return nil, err
			}
			msgs = append(msgs, newMsgs...)

			log.Printf("Epoch %s ended. Processing Epoch end handling messages...\n", event.String())
			// only execute once per block
			break
		}
	}
	return msgs, nil
}

func GetMsgEndBlockEvents(cliCtx client.Context, flgs *flag.FlagSet, events []tmtypes.Event, store *types.Storage) ([]sdk.Msg, error) {
	msgs := []sdk.Msg{}
	for _, event := range events {
		if event.Type == "complete_unbonding" {
			newMsgs, err := GetMsgsUndelegateComplete(cliCtx, flgs, event)
			if err != nil {
				return []sdk.Msg{}, err
			}
			msgs = append(msgs, newMsgs...)

			log.Println("Unbonding completed. Processing Epoch end handling messages...")
		}
	}

	// handle batch unbonding initiator at endblock message generator
	msg, err := issueBatchUnbonding(cliCtx, flgs, store)
	msgs = append(msgs, msg...)
	return msgs, err
}

func issueBatchUnbonding(cliCtx client.Context, flgs *flag.FlagSet, store *types.Storage) ([]sdk.Msg, error) {
	params, err := getStakingParams(cliCtx, flgs)
	if err != nil {
		return nil, err
	}
	intervalNano := params.UnbondingTime.Nanoseconds() / int64(params.MaxEntries)

	durationSinceLast := time.Now().Sub(store.LastUndelegationTime)

	if store.PendingUndelegations.Cmp(big.NewInt(0)) == 0 || durationSinceLast.Nanoseconds() <= intervalNano {
		return nil, nil
	}
	from := cliCtx.GetFromAddress()
	val, err := flgs.GetString(flagValidator)
	if err != nil {
		return nil, err
	}
	delCoins := sdk.NewCoin(params.BondDenom, sdk.NewIntFromBigInt(store.PendingUndelegations))
	valAddr, err := sdk.ValAddressFromBech32(val)
	if err != nil {
		return nil, err
	}
	store.PendingUndelegations = big.NewInt(0)

	return []sdk.Msg{stakingtypes.NewMsgUndelegate(from, valAddr, delCoins)}, nil
}

// GetMsgsEpochEnd handles epoch ending by delegating the newly received coins
func GetMsgsEpochEnd(cliCtx client.Context, flgs *flag.FlagSet) ([]sdk.Msg, error) {
	rewardsEarned, err := getStakingRewards(cliCtx, flgs)
	bondDenom, err := getBondDenom(cliCtx)
	if err != nil {
		return nil, err
	}

	// calculate delegation and distribution amounts.
	rewardAmt := rewardsEarned.AmountOf(bondDenom)
	totalAmt, err := GetTotalAsset(cliCtx, flgs)
	if err != nil {
		return nil, err
	}
	distributionInt, err := getDistributionAmt(cliCtx, flgs, totalAmt.AmountOf(bondDenom))
	if err != nil {
		return nil, err
	}
	log.Printf("Calculated distribution amount: %s\n", distributionInt.String())
	delegateAmt := rewardAmt.Sub(sdk.NewIntFromBigInt(distributionInt))
	var delegateCoin sdk.Coin
	if !delegateAmt.IsNegative() {
		delegateCoin = sdk.NewCoin(bondDenom, delegateAmt)
	} else {
		delegateCoin = sdk.NewCoin(bondDenom, sdk.ZeroInt())
	}

	// TODO: add Ethereum Tx to feed the total asset
	err = accrueAsset(cliCtx, flgs, distributionInt)

	if delegateCoin.Amount.Equal(sdk.ZeroInt()) {
		// no delegation
		return nil, nil
	}

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

	// construct tx with withdraw & delegate commands
	withdrawMsg := distrtypes.NewMsgWithdrawDelegatorReward(delAddr, valAddr)
	msgs := []sdk.Msg{withdrawMsg}
	delegateMsg := stakingtypes.NewMsgDelegate(delAddr, valAddr, delegateCoin)
	msgs = append(msgs, delegateMsg)

	return msgs, err
}

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

func accrueAsset(cliCtx client.Context, flgs *flag.FlagSet, amt *big.Int) error {
	totalAsset, err := GetTotalAsset(cliCtx, flgs)
	if err != nil {
		return err
	}
	totalAssetInt := totalAsset.AmountOf(baseDenom)
	args := totalAssetInt.BigInt()

	contractABI, err := stayking.StaykingMetaData.GetAbi()
	if err != nil {
		return err
	}
	inputData, err := contractABI.Pack(accrueFuncName, args)
	if err != nil {
		return err
	}

	contractAddr, err := flgs.GetString(flagStaykingContAddr)
	if err != nil {
		return err
	}
	_, err = ConstructEthTx(cliCtx, flgs, contractAddr, amt, inputData)
	return err
}

// GetMsgsUndelegateComplete handles completed unbondings by sending the unlocked coins to unbonding contract
func GetMsgsUndelegateComplete(cliCtx client.Context, flgs *flag.FlagSet, event tmtypes.Event) ([]sdk.Msg, error) {
	var unbondedCoin sdk.Coin
	args := event.GetAttributes()
	for _, a := range args {
		if string(a.Key) == "amount" {
			v := string(a.Value)
			log.Println(string(a.Key))
			log.Println(v)

			var err error
			unbondedCoin, err = sdk.ParseCoinNormalized(v)
			if err != nil {
				return nil, err
			}
		}
	}
	value := unbondedCoin.Amount.BigInt()

	uEVMOSContAddr, err := flgs.GetString(flagUnbondingEvmosContAddr)
	if err != nil {
		return nil, err
	}
	contractABI, err := abis.AbisMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	inputData, err := contractABI.Pack(supplyUnbondedTokenFuncName)
	if err != nil {
		return nil, err
	}

	_, err = ConstructEthTx(cliCtx, flgs, uEVMOSContAddr, value, inputData)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// GetTotalAsset gets the total asset on the account
// which consists of coins on account + staked amount + unbonding amount + rewards
func GetTotalAsset(cliCtx client.Context, flgs *flag.FlagSet) (sdk.Coins, error) {
	ch1 := make(chan sdk.Coins, 4)
	ch2 := make(chan error, 4)
	go func() {
		rewards, err := getStakingRewards(cliCtx, flgs)
		ch1 <- rewards
		ch2 <- err
	}()
	go func() {
		balances, err := getAccountBalances(cliCtx, flgs)
		ch1 <- balances
		ch2 <- err
	}()
	go func() {
		unbonding, err := getTotalUnbonding(cliCtx, flgs)
		ch1 <- unbonding
		ch2 <- err
	}()
	go func() {
		bonded, err := getTotalBonded(cliCtx, flgs)
		ch1 <- bonded
		ch2 <- err
	}()

	var total sdk.Coins
	for i := 0; i < 4; i++ {
		total = total.Add(<-ch1...)
	}
	close(ch1)
	return total, nil
}

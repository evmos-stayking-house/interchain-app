package cmd

import (
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/evmos-stayking-house/scheduled-worker-golang/abis"
	flag "github.com/spf13/pflag"
	"github.com/tendermint/tendermint/abci/types"
	"log"
	"math/big"

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

func GetMsgNewBlockEvents(cliCtx client.Context, flgs *flag.FlagSet, events []types.Event) ([]sdk.Msg, error) {
	msgs := []sdk.Msg{}
	for _, event := range events {
		if event.Type == "epoch_end" {
			newMsgs, err := GetMsgsEpochEnd(cliCtx, flgs)
			if err != nil {
				return []sdk.Msg{}, nil
			}
			msgs = append(msgs, newMsgs...)

			log.Printf("Epoch %s ended. Processing Epoch end handling messages...\n", event.String())
			// only execute once per block
			break
		}
	}
	return msgs, nil
}

func GetMsgEndBlockEvents(cliCtx client.Context, flgs *flag.FlagSet, events []types.Event) ([]sdk.Msg, error) {
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
	return msgs, nil
}

// GetMsgsEpochEnd handles epoch ending by delegating the newly received coins
func GetMsgsEpochEnd(cliCtx client.Context, flgs *flag.FlagSet) ([]sdk.Msg, error) {
	toCompound, err := getStakingRewards(cliCtx, flgs)
	bondDenom, err := getBondDenom(cliCtx)
	if err != nil {
		return []sdk.Msg{}, err
	}

	delegateCoin := sdk.NewCoin(bondDenom, toCompound.AmountOf(bondDenom))

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

	// TODO: add Ethereum Tx to feed the total asset
	// res, err := GetTotalAsset(cliCtx, flgs)

	return msgs, nil
}

// GetMsgsUndelegateComplete handles completed unbondings by sending the unlocked coins to unbonding contract
func GetMsgsUndelegateComplete(cliCtx client.Context, flgs *flag.FlagSet, event types.Event) ([]sdk.Msg, error) {
	uEVMOSContAddr, err := flgs.GetString(flagUnbondingEvmosContAddr)
	if err != nil {
		return nil, err
	}
	contractABI, err := abis.AbisMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	var unbondedCoin sdk.Coin
	args := event.GetAttributes()
	for _, a := range args {
		if string(a.Key) == "amount" {
			v := string(a.Value)
			log.Println(string(a.Key))
			log.Println(v)
			unbondedCoin, err = sdk.ParseCoinNormalized(v)
			if err != nil {
				return nil, err
			}
		}
	}
	value := unbondedCoin.Amount.BigInt()
	inputData, err := contractABI.Pack("supplyUnbondedToken")
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
	go func() {
		rewards, _ := getStakingRewards(cliCtx, flgs)
		ch1 <- rewards
	}()
	go func() {
		balances, _ := getAccountBalances(cliCtx, flgs)
		ch1 <- balances
	}()
	go func() {
		unbonding, _ := getTotalUnbonding(cliCtx, flgs)
		ch1 <- unbonding
	}()
	go func() {
		bonded, _ := getTotalBonded(cliCtx, flgs)
		ch1 <- bonded
	}()
	var total sdk.Coins
	for i := 0; i < 4; i++ {
		total = total.Add(<-ch1...)
	}
	close(ch1)
	return total, nil
}

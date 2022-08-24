package cmd

import (
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
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
	var epoch_handled bool
	for _, event := range events {
		if event.Type == "epoch_end" && !epoch_handled {
			newMsgs, err := GetMsgsEpochEnd(cliCtx, flgs)
			if err != nil {
				return []sdk.Msg{}, nil
			}
			msgs = append(msgs, newMsgs...)

			log.Printf("Epoch %s ended. Processing Epoch end handling messages...\n", event.String())
			// only execute once per block
			epoch_handled = true
		} else if event.Type == "complete_unbonding" {
			newMsgs, err := GetMsgsUndelegateComplete(cliCtx, flgs, event)
			if err != nil {
				return []sdk.Msg{}, nil
			}
			msgs = append(msgs, newMsgs...)

			log.Println("Unbonding completed. Processing Epoch end handling messages...")
		}
	}
	return []sdk.Msg{}, nil
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
	return msgs, nil
}

// GetMsgsUndelegateComplete handles completed unbondings by sending the unlocked coins to unbonding contract
func GetMsgsUndelegateComplete(ctx client.Context, flgs *flag.FlagSet, event types.Event) ([]sdk.Msg, error) {
	log.Println(event.String())
	return []sdk.Msg{}, nil
}

// GetTotalAsset gets the total asset on the account
// which consists of coins on account + staked amount + unbonding amount + rewards
func GetTotalAsset(cliCtx client.Context, flgs *flag.FlagSet) (sdk.Coins, error) {
	rewards, err := getStakingRewards(cliCtx, flgs)
	if err != nil {
		return sdk.Coins{}, err
	}
	balances, err := getAccountBalances(cliCtx, flgs)
	if err != nil {
		return sdk.Coins{}, err
	}
	unbonding, err := getTotalUnbonding(cliCtx, flgs)
	if err != nil {
		return sdk.Coins{}, err
	}
	bonded, err := getTotalBonded(cliCtx, flgs)
	if err != nil {
		return sdk.Coins{}, err
	}

	total := rewards.Add(balances...).Add(unbonding...).Add(bonded...)
	return total, nil
}

package cmd

import (
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/evmos-stayking-house/interchain-app/abis"
	"github.com/evmos-stayking-house/interchain-app/abis/stayking"
	"github.com/evmos-stayking-house/interchain-app/types"
	"github.com/evmos/ethermint/rpc/backend"
	flag "github.com/spf13/pflag"
	abcitypes "github.com/tendermint/tendermint/abci/types"
	tmtypes "github.com/tendermint/tendermint/abci/types"
	"log"
	"math/big"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func handleContractEvent(event abcitypes.Event, store *types.Storage) error {
	// prepare filters and queries
	contractAbi, err := stayking.StaykingMetaData.GetAbi()
	if err != nil {
		log.Fatal(err)
	}

	logs, err := backend.ParseTxLogsFromEvent(event)
	if err != nil {
		return err
	}
	for _, l := range logs {
		if l.Topics[0].Hex() == LogDelegateSigHash.Hex() {
			data, err := contractAbi.Unpack("Stake", l.Data)
			if err != nil {
				log.Fatal(err)
			}
			amtInt := data[0].(*big.Int)
			store.PendingDelegations = store.PendingDelegations.Add(store.PendingDelegations, amtInt)
			log.Printf("Detected delegation... %s\n", amtInt.String())
		} else if l.Topics[0].Hex() == LogUndelegateSigHash.Hex() {
			data, err := contractAbi.Unpack("Unstake", l.Data)
			if err != nil {
				log.Fatal(err)
			}
			amtInt := data[0].(*big.Int)
			store.PendingUndelegations = store.PendingUndelegations.Add(store.PendingUndelegations, amtInt)
			log.Printf("Detected undelegation... %s\n", amtInt.String())
		} else if l.Topics[0].Hex() == LogAccrueSigHash.Hex() {
			data, err := contractAbi.Unpack("Accrue", l.Data)
			if err != nil {
				log.Fatal(err)
			}
			amtInt := data[0].(*big.Int)
			store.PendingDelegations = store.PendingDelegations.Add(store.PendingDelegations, amtInt)
			log.Printf("Detected compound... %s\n", amtInt.String())
		} else {
			log.Println("received some other event from the contract. skipping...")
		}
	}
	return nil
}

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
			log.Println("Unbonding completed. Processing complete unbonding handling messages...")
			newMsgs, err := GetMsgsUndelegateComplete(cliCtx, flgs, event)
			if err != nil {
				return []sdk.Msg{}, err
			}
			msgs = append(msgs, newMsgs...)
		}
	}

	// handle batch unbonding initiator at endblock message generator
	msg, err := issueBatchUnbonding(cliCtx, flgs, store)
	if err != nil {
		return []sdk.Msg{}, err
	}
	msgs = append(msgs, msg...)

	return msgs, types.WriteStore(cliCtx, *store)
}

func issueBatchUnbonding(cliCtx client.Context, flgs *flag.FlagSet, store *types.Storage) ([]sdk.Msg, error) {
	params, err := getStakingParams(cliCtx, flgs)
	if err != nil {
		return nil, err
	}
	intervalNano := params.UnbondingTime.Nanoseconds() / int64(params.MaxEntries)

	durationSinceLast := time.Now().Sub(store.LastUndelegationTime)

	if store.PendingUndelegations.Cmp(big.NewInt(0)) == 0 || durationSinceLast.Nanoseconds() <= intervalNano {
		return []sdk.Msg{}, nil
	}
	log.Println("issuing batch unbonding...")
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
	store.LastUndelegationTime = time.Now()
	err = types.WriteStore(cliCtx, *store)

	return []sdk.Msg{stakingtypes.NewMsgUndelegate(from, valAddr, delCoins)}, err
}

// GetMsgsEpochEnd handles epoch ending by delegating the newly received coins
func GetMsgsEpochEnd(cliCtx client.Context, flgs *flag.FlagSet) ([]sdk.Msg, error) {
	revenueAmt, err := getRevenue(cliCtx, flgs)
	log.Printf("calculated rewards amount: %s\n", revenueAmt.String())
	rewardIsZero := revenueAmt.Cmp(big.NewInt(0)) == 0
	if err != nil || rewardIsZero {
		return []sdk.Msg{}, err
	}

	// fire and wait for withdraw rewards to complete
	if !rewardIsZero {
		log.Println("withdrawing rewards...")
		err = withdrawRewards(cliCtx, flgs)
		if err != nil {
			return []sdk.Msg{}, err
		}
	}

	err = accrueAsset(cliCtx, flgs, revenueAmt)
	if err != nil {
		return []sdk.Msg{}, err
	}

	// no msgs for now
	return []sdk.Msg{}, err
}

// GetMsgsUndelegateComplete handles completed unbondings by sending the unlocked coins to unbonding contract
func GetMsgsUndelegateComplete(cliCtx client.Context, flgs *flag.FlagSet, event tmtypes.Event) ([]sdk.Msg, error) {
	var unbondedCoin sdk.Coin
	args := event.GetAttributes()
	for _, a := range args {
		if string(a.Key) == "amount" {
			v := string(a.Value)
			log.Printf("Detected unbonding completion: %s\n", v)

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
		bonded := getTotalBonded(cliCtx, flgs)
		ch1 <- bonded
	}()

	var total sdk.Coins
	for i := 0; i < 4; i++ {
		total = total.Add(<-ch1...)
	}
	close(ch1)
	return total, nil
}

// GetStakedAsset gets the total staked asset
func GetStakedAsset(cliCtx client.Context, flgs *flag.FlagSet) sdk.Coins {
	return getTotalBonded(cliCtx, flgs)
}

// gets total staking rewards + any available asset.
// Leaves a small amount left for txs
func getRevenue(cliCtx client.Context, flgs *flag.FlagSet) (*big.Int, error) {
	ch1 := make(chan sdk.Coins, 2)
	go func() {
		rewards, _ := getStakingRewards(cliCtx, flgs)
		ch1 <- rewards
	}()
	go func() {
		balances, _ := getAccountBalances(cliCtx, flgs)
		ch1 <- balances
	}()
	var total sdk.Coins
	for i := 0; i < 2; i++ {
		total = total.Add(<-ch1...)
	}
	close(ch1)

	bondDenom, err := getBondDenom(cliCtx)
	if err != nil {
		return nil, err
	}

	totalAmt := total.AmountOf(bondDenom).BigInt()

	//leave some behind for tx fees. if total balance < min account balance, return 0
	res := totalAmt.Sub(totalAmt, minAccBalance)
	// if res < 0, return 0
	if res.Cmp(big.NewInt(0)) < 0 {
		return big.NewInt(0), nil
	}
	return res, nil
}

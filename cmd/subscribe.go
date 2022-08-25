package cmd

import (
	"context"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/evmos-stayking-house/scheduled-worker-golang/abis/stayking"
	"github.com/evmos-stayking-house/scheduled-worker-golang/types"
	"github.com/evmos/ethermint/rpc/backend"
	flag "github.com/spf13/pflag"
	tmclient "github.com/tendermint/tendermint/rpc/client/http"
	tmtypes "github.com/tendermint/tendermint/types"
	"log"
	"math/big"
)

// Subscribe handles listeners for delegation, unbonding completion, and epoch ending events.
func Subscribe(contAddr string, cliCtx client.Context, flgs *flag.FlagSet) error {
	// initialize epochClient
	epochClient, err := tmclient.New(cliCtx.NodeURI, "/websocket")
	if err != nil {
		return err
	}

	err = epochClient.Start()
	if err != nil {
		return err
	}

	// prepare filters and queries
	contractAbi, err := stayking.StaykingMetaData.GetAbi()
	if err != nil {
		log.Fatal(err)
	}
	query := "tm.event='Tx' AND ethereum_tx.recipient='" + contAddr + "'"
	// query := "tm.event='Tx'"
	log.Println("constructed query: ", query)

	// subscribe to new block events according to the query
	txChan, err := epochClient.Subscribe(context.Background(), "",
		query, 10000)
	if err != nil {
		return err
	}

	// subscribe to new block events according to the query for clearing out aggregated delegation
	blockChan, err := epochClient.Subscribe(context.Background(), "",
		"tm.event='NewBlock'", 10000)
	if err != nil {
		return err
	}

	var delegationTracker types.SafeTotal
	delegationTracker.Amount = big.NewInt(0)
	var UndelegationTracker types.SafeTotal
	UndelegationTracker.Amount = big.NewInt(0)

	for {
		select {
		case txn := <-txChan:
			// get tendermint transaction data in struct of ResponseDeliverTx
			txData, ok := txn.Data.(tmtypes.EventDataTx)
			if !ok {
				log.Fatal("couldn't cast txn event data")
			}
			// undelegate completion in endblocker in vanilla staking
			// https://github.com/evmos/evmos/blob/9aba6f4fd4c3bc6772c503a2c459111065aba3d8/x/epochs/keeper/abci.go#L14-L14
			res := txData.GetResult()
			for _, event := range res.Events {
				logs, err := backend.ParseTxLogsFromEvent(event)
				if err != nil {
					return err
				}
				for _, l := range logs {
					if l.Topics[0].Hex() == logDelegateSigHash.Hex() {
						amt, err := contractAbi.Unpack("Stake", l.Data)
						if err != nil {
							log.Fatal(err)
						}
						amtInt := amt[0].(*big.Int)
						delegationTracker.Add(amtInt)
					} else if l.Topics[0].Hex() == logUndelegateSigHash.Hex() {
						amt, err := contractAbi.Unpack("Unstake", l.Data)
						if err != nil {
							log.Fatal(err)
						}
						amtInt := amt[0].(*big.Int)
						UndelegationTracker.Add(amtInt)
					} else {
						log.Println("received something else! wow")
					}
				}
			}
		case nb := <-blockChan:
			// get tendermint transaction data in struct of ResponseDeliverTx
			blockData, ok := nb.Data.(tmtypes.EventDataNewBlock)
			if !ok {
				log.Fatal("couldn't cast block event data")
			}
			eventsBB := blockData.ResultBeginBlock.GetEvents()
			msgsNewBlock, err := GetMsgNewBlockEvents(cliCtx, flgs, eventsBB)
			if err != nil {
				log.Fatal(err)
			}
			eventsEB := blockData.ResultEndBlock.GetEvents()
			msgsEndBlock, err := GetMsgEndBlockEvents(cliCtx, flgs, eventsEB)
			if err != nil {
				log.Fatal(err)
			}

			// epoch starts/ends in beginblocker in evmos
			// https://github.com/evmos/evmos/blob/9aba6f4fd4c3bc6772c503a2c459111065aba3d8/x/epochs/keeper/abci.go#L14-L14

			msgs := append(msgsNewBlock, msgsEndBlock...)
			// if no delegation, wrap up tx construction here
			if delegationTracker.Amount.Cmp(big.NewInt(0)) == 0 {
				if len(msgs) == 0 {
					continue
				}
				err = tx.GenerateOrBroadcastTxCLI(cliCtx, flgs, msgs...)
				if err != nil {
					log.Fatal(err)
				}
				continue
			}
			log.Printf("New Delegation detected! Total delegation: %s\n", delegationTracker.Amount)
			toDelegate := delegationTracker.Clear()
			delMsg, err := GetMsgDelegation(cliCtx, flgs, toDelegate)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Delegating received tokens: %s\n", toDelegate.String())
			msgs = append(msgs, delMsg)
			err = tx.GenerateOrBroadcastTxCLI(cliCtx, flgs, msgs...)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	return nil
}

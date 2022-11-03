package cmd

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/evmos-stayking-house/interchain-app/types"
	flag "github.com/spf13/pflag"
	tmclient "github.com/tendermint/tendermint/rpc/client/http"
	tmtypes "github.com/tendermint/tendermint/types"
	"log"
	"math/big"
)

// Subscribe handles listeners for delegation, unbonding completion, and epoch ending events.
func Subscribe(cliCtx client.Context, flgs *flag.FlagSet) error {
	// initialize epochClient
	epochClient, err := tmclient.New(cliCtx.NodeURI, "/websocket")
	if err != nil {
		return err
	}

	err = epochClient.Start()
	if err != nil {
		return err
	}

	// construct tx event websocket query
	contractAddr, _ := flgs.GetString(flagStaykingContAddr)
	query := "tm.event='Tx' AND ethereum_tx.recipient='" + contractAddr + "'"

	// subscribe to tx events according to the query
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

	storage, err := types.ReadStore(cliCtx)
	if err != nil {
		_ = types.Init(cliCtx)
		storage, _ = types.ReadStore(cliCtx)
	}
	fmt.Printf("Loaded stored status: %s\n", storage)

	numRetry, err := flgs.GetUint64(flagMaxRetry)
	if err != nil {
		log.Fatal(err)
	}

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
				if err := handleContractEvent(event, &storage); err != nil {
					log.Fatal(err)
				}
			}
			log.Printf("Storing updated status... %s\n", storage)
			if err := types.WriteStore(cliCtx, storage); err != nil {
				log.Fatal(err)
			}
		case nb := <-blockChan:
			// get tendermint transaction data in struct of ResponseDeliverTx
			blockData, ok := nb.Data.(tmtypes.EventDataNewBlock)
			if !ok {
				log.Fatal("couldn't cast block event data")
			}

			eventsBB := blockData.ResultBeginBlock.GetEvents()
			msgsNewBlock, err := GetMsgNewBlockEvents(cliCtx, flgs, eventsBB, &storage)
			if err != nil {
				log.Println("errored while generating newblock event messages")
				log.Fatal(err)
			}
			eventsEB := blockData.ResultEndBlock.GetEvents()
			msgsEndBlock, err := GetMsgEndBlockEvents(cliCtx, flgs, eventsEB, &storage)
			if err != nil {
				log.Println("errored while generating endblock event messages")
				log.Fatal(err)
			}

			// epoch starts/ends in beginblocker in evmos
			// https://github.com/evmos/evmos/blob/9aba6f4fd4c3bc6772c503a2c459111065aba3d8/x/epochs/keeper/abci.go#L14-L14
			log.Printf("detected new block %d! current delegation pending: %s current undelegation pending: %s\n",
				blockData.Block.Height, storage.PendingDelegations.String(), storage.PendingUndelegations.String())
			msgs := append(msgsNewBlock, msgsEndBlock...)
			// if no delegation, wrap up tx construction here
			if storage.PendingDelegations.Cmp(big.NewInt(0)) == 0 {
				if len(msgs) == 0 {
					continue
				}
				err = TrySubmitTxMaxRetry(numRetry, cliCtx, flgs, msgs...)
				if err != nil {
					log.Fatal(err)
				}
				continue
			}

			if storage.PendingDelegations.Cmp(big.NewInt(0)) != 0 {
				log.Printf("New Delegation detected! Total delegation: %s\n", storage.PendingDelegations.String())
				log.Printf("Delegating tokens: %s\n", storage.PendingDelegations.String())
				delMsg, err := GetMsgDelegation(cliCtx, flgs, storage.PendingDelegations)
				if err != nil {
					log.Fatal(err)
				}
				msgs = append(msgs, delMsg)
			}

			err = TrySubmitTxMaxRetry(numRetry, cliCtx, flgs, msgs...)
			if err != nil {
				log.Fatal(err)
			}
			storage.PendingDelegations = big.NewInt(0)
			err = types.WriteStore(cliCtx, storage)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	return nil
}

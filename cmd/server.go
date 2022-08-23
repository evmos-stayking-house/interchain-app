package cmd

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/evmos-stayking-house/scheduled-worker-golang/abis"
	"github.com/evmos-stayking-house/scheduled-worker-golang/types"
	"github.com/evmos/ethermint/rpc/backend"
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
	tmclient "github.com/tendermint/tendermint/rpc/client/http"
	tmtypes "github.com/tendermint/tendermint/types"
	"log"
	"math/big"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
)

func ServeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "serve",
		Short:                      "Serve subcommands",
		SuggestionsMinimumDistance: 2,
	}

	cmd.AddCommand(
		NewSubscribeDelegationCmd(),
		NewSubscribeEpochCmd(),
		NewSubscribeUndelegateCmd(),
	)

	return cmd
}

func NewSubscribeUndelegateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "subscribe-undelegate",
		Short: `Subscribes to unbonding end event and executes contract call commands.`,
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			return SubscribeUndelegate(cliCtx, cmd.Flags())
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().Set(flags.FlagSkipConfirmation, "true")
	cmd.Flags().Set(flags.FlagBroadcastMode, "block")
	cmd.Flags().Set(flags.FlagKeyringBackend, "test")
	cmd.Flags().Set(flags.FlagGasAdjustment, "1.5")
	cmd.Flags().Set(flags.FlagGas, "auto")
	cmd.Flags().Set(flags.FlagGasPrices, "10000000000"+baseDenom)
	return cmd
}

func SubscribeUndelegate(cliCtx client.Context, flgs *flag.FlagSet) error {
	// initialize epochClient
	epochClient, err := tmclient.New(cliCtx.NodeURI, "/websocket")
	if err != nil {
		return err
	}

	err = epochClient.Start()
	if err != nil {
		return err
	}

	// subscribe to new block events according to the query
	txChan, err := epochClient.Subscribe(context.Background(), "",
		"tm.event='NewBlock' AND complete_unbonding.delegator='evmos10vvd5e9kdezyjtnyrld2nfq7v8482ajlaamwus'", 10000)
	if err != nil {
		return err
	}

	for {
		select {
		case tx := <-txChan:
			// get tendermint transaction data in struct of ResponseDeliverTx
			txData, ok := tx.Data.(tmtypes.EventDataNewBlock)
			if !ok {
				log.Fatal("received non-newblock event")
			}
			// undelegate completion in endblocker in vanilla staking
			// https://github.com/evmos/evmos/blob/9aba6f4fd4c3bc6772c503a2c459111065aba3d8/x/epochs/keeper/abci.go#L14-L14
			for _, event := range txData.ResultEndBlock.GetEvents() {
				if event.Type == "complete_unbonding" {
					err := HandleUndelegateComplete(cliCtx, flgs, event)
					if err != nil {
						panic(err)
					}

					fmt.Println(event.String())
					// only execute once per block
					break

				}
			}
		}
	}
	return nil
}

func NewSubscribeEpochCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "subscribe-epoch",
		Short: `Subscribes to epoch end event and executes staking commands.`,
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			return SubscribeEpoch(cliCtx, cmd.Flags())
		},
	}
	cmd.Flags().String(flagValidator, "evmosvaloper10vvd5e9kdezyjtnyrld2nfq7v8482ajlsn57ad", "The validator to undelegate from")

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().Set(flags.FlagSkipConfirmation, "true")
	cmd.Flags().Set(flags.FlagBroadcastMode, "block")
	cmd.Flags().Set(flags.FlagKeyringBackend, "test")
	cmd.Flags().Set(flags.FlagGasAdjustment, "1.5")
	cmd.Flags().Set(flags.FlagGas, "auto")
	cmd.Flags().Set(flags.FlagGasPrices, "10000000000"+baseDenom)
	return cmd
}

func SubscribeEpoch(cliCtx client.Context, flgs *flag.FlagSet) error {
	// initialize epochClient
	epochClient, err := tmclient.New(cliCtx.NodeURI, "/websocket")
	if err != nil {
		return err
	}

	err = epochClient.Start()
	if err != nil {
		return err
	}

	// subscribe to new block events according to the query
	blockChan, err := epochClient.Subscribe(context.Background(), "",
		"tm.event='NewBlock' AND epoch_end.epoch_number EXISTS", 10000)
	if err != nil {
		return err
	}

	for {
		select {
		case tx := <-blockChan:
			// get tendermint transaction data in struct of ResponseDeliverTx
			txData, ok := tx.Data.(tmtypes.EventDataNewBlock)
			if !ok {
				log.Fatal("received non-newblock event")
			}
			// epoch starts/ends in beginblocker in evmos
			// https://github.com/evmos/evmos/blob/9aba6f4fd4c3bc6772c503a2c459111065aba3d8/x/epochs/keeper/abci.go#L14-L14
			for _, event := range txData.ResultBeginBlock.GetEvents() {
				if event.Type == "epoch_end" {
					err := HandleEpochEnd(cliCtx, flgs)
					if err != nil {
						panic(err)
					}

					fmt.Println(event.String())
					// only execute once per block
					break

				}
			}
		}
	}
	return nil
}

func NewSubscribeDelegationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "subscribe-delegation",
		Short: `Subscribes to a delegation event in a specific contract address.`,
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			contractAddr, _ := cmd.Flags().GetString(flagContAddr)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			return SubscribeDelegation(contractAddr, clientCtx, cmd.Flags())
		},
	}

	cmd.Flags().String(flagContAddr, "", "The contract address to listen to")
	cmd.Flags().String(flagValidator, "", "The validator to delegate to")

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().Set(flags.FlagSkipConfirmation, "true")
	cmd.Flags().Set(flags.FlagBroadcastMode, "block")
	cmd.Flags().Set(flags.FlagKeyringBackend, "test")
	cmd.Flags().Set(flags.FlagGasAdjustment, "1.5")
	cmd.Flags().Set(flags.FlagGas, "auto")
	cmd.Flags().Set(flags.FlagGasPrices, "10000000000"+baseDenom)

	return cmd
}

type DelegationChange struct {
	Delegator common.Address
	Amount    *big.Int
}

func SubscribeDelegation(contAddr string, cliCtx client.Context, flgs *flag.FlagSet) error {
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
	contractAbi, err := abi.JSON(strings.NewReader(abis.EventsMetaData.ABI))
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

	for {
		select {
		case tx := <-txChan:
			// get tendermint transaction data in struct of ResponseDeliverTx
			txData, ok := tx.Data.(tmtypes.EventDataTx)
			if !ok {
				log.Fatal("received tx to contract address event")
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
					log.Println(l)
					amt, err := contractAbi.Unpack("Delegate", l.Data)
					if err != nil {
						log.Fatal(err)
					}
					amtInt := amt[0].(*big.Int)
					delegationTracker.Add(amtInt)
				}
			}
		case _ = <-blockChan:
			if delegationTracker.Amount.Cmp(big.NewInt(0)) == 0 {
				continue
			}
			log.Printf("New Delegation detected! Total delegation: %s\n", delegationTracker.Amount)
			toDelegate := delegationTracker.Clear()
			if err := HandleDelegation(cliCtx, flgs, toDelegate); err != nil {
				log.Fatal(err)
			}
			log.Println(toDelegate.String())
		}
	}
	return nil
}

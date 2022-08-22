package cmd

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	tmclient "github.com/tendermint/tendermint/rpc/client/http"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/evmos-stayking-house/scheduled-worker-golang/events"
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
	txChan, err := epochClient.Subscribe(context.Background(), "",
		"tm.event='NewBlock' AND epoch_end.epoch_number EXISTS", 10000)
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
			ethEndpoint, _ := cmd.Flags().GetString(flagEthEndpoint)
			contractAddr, _ := cmd.Flags().GetString(flagContAddr)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			return SubscribeDelegation(ethEndpoint, contractAddr, clientCtx, cmd.Flags())
		},
	}

	cmd.Flags().String(flagEthEndpoint, "ws://localhost:8546", "The ethereum websocket endpoint to subscribe to")
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

func SubscribeDelegation(ethEndpoint, contAddr string, ctx client.Context, flgs *flag.FlagSet) error {
	wsclient, err := ethclient.Dial(ethEndpoint)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("subscribing to ethereum endpoint: ", ethEndpoint)

	// set contract addr and ABI
	contractAddress := common.HexToAddress(contAddr)

	// TODO: make this a parameter
	contractAbi, err := abi.JSON(strings.NewReader(events.EventsMetaData.ABI))
	if err != nil {
		log.Fatal(err)
	}

	logDelegateSig := []byte("Delegate(address,uint256)")
	logUndelegateSig := []byte("Undelegate(address,uint256)")

	logDelegateSigHash := crypto.Keccak256Hash(logDelegateSig)
	logUndelegateSigHash := crypto.Keccak256Hash(logUndelegateSig)
	fmt.Println(logDelegateSigHash.String())
	fmt.Println(logUndelegateSigHash.String())

	// subscribe to the staking contract
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)

	sub, err := wsclient.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog)
			switch vLog.Topics[0].Hex() {
			case logDelegateSigHash.Hex():
				amt, err := contractAbi.Unpack("Delegate", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}
				var d DelegationChange
				d.Amount = amt[0].(*big.Int)
				d.Delegator = common.HexToAddress(vLog.Topics[1].String())
				log.Printf("Delegate detected!: From: %s   Amount: %d\n", d.Delegator, d.Amount)
				if err = HandleDelegation(ctx, flgs, d); err != nil {
					log.Fatal(err)
				}
			case logUndelegateSigHash.Hex():
				amt, err := contractAbi.Unpack("Undelegate", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}
				var d DelegationChange
				d.Amount = amt[0].(*big.Int)
				d.Delegator = common.HexToAddress(vLog.Topics[1].String())
				log.Printf("Undelegated detected!: From: %s   Amount: %d\n", d.Delegator, d.Amount)
				//
			default:
				log.Println("Not delegation or undelegation")
			}
		}
	}
}

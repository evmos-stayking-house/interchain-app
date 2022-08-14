package cmd

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/spf13/cobra"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/evmos-stayking-house/scheduled-worker-golang/events"
)

var (
	flagEthEndpoint = "eth-endpoint"
	flagContAddr    = "cont-addr"
)

func ServeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "serve",
		Short:                      "Serve subcommands",
		SuggestionsMinimumDistance: 2,
	}

	cmd.AddCommand(NewSubscribeDelegationCmd())

	return cmd
}

func NewSubscribeDelegationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "subscribe-delegation",
		Short: `Subscribes to a delegation event in a specific contract address.`,
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			ethEndpoint, _ := cmd.Flags().GetString(flagEthEndpoint)
			contractAddr, _ := cmd.Flags().GetString(flagContAddr)

			return SubscribeDelegation(ethEndpoint, contractAddr)
		},
	}

	cmd.Flags().String(flagEthEndpoint, "wss://ropsten.infura.io/ws/v3/d0383d521441488fb754735af7fe0c59", "The ethereum websocket endpoint to subscribe to")
	cmd.Flags().String(flagContAddr, "0x50fCe2E7426FFfEd8762e21bdf7E0Fe9188eD54A", "The contract address to listen to")

	return cmd
}

type DelegationChange struct {
	Delegator common.Address
	Amount    *big.Int
}

func SubscribeDelegation(ethEndpoint, contAddr string) error {
	wsclient, err := ethclient.Dial(ethEndpoint)
	if err != nil {
		log.Fatal(err)
	}

	// set contract addr and ABI
	// TODO: make this a parameter
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
			switch vLog.Topics[0].Hex() {
			case logDelegateSigHash.Hex():
				amt, err := contractAbi.Unpack("Delegate", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}
				var d DelegationChange
				d.Amount = amt[0].(*big.Int)
				d.Delegator = common.HexToAddress(vLog.Topics[1].String())
				fmt.Printf("Delegate detected!: From: %s   Amount: %d\n", d.Delegator, d.Amount)
				if err = HandleDelegation(d); err != nil {
					log.Fatal(err)
				}
				//
			case logUndelegateSigHash.Hex():
				amt, err := contractAbi.Unpack("Undelegate", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}
				var d DelegationChange
				d.Amount = amt[0].(*big.Int)
				d.Delegator = common.HexToAddress(vLog.Topics[1].String())
				fmt.Printf("Undelegated detected!: From: %s   Amount: %d\n", d.Delegator, d.Amount)
				if err = HandleUndelegation(d); err != nil {
					log.Fatal(err)
				}
				//
			}
		}
	}
}
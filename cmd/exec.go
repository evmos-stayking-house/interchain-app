package cmd

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"log"
	"math"
	"math/big"
)

func ExecuteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "exec",
		Short:                      "execution subcommands",
		SuggestionsMinimumDistance: 2,
	}

	cmd.AddCommand(NewExecuteUndelegationCmd())

	return cmd
}

func NewExecuteUndelegationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "execute-undelegation",
		Short: `batch queries any unbonding event and initiates an unbonding for the corresponding amount.`,
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			ethEndpoint, _ := cmd.Flags().GetString(flagEthEndpoint)
			contractAddr, _ := cmd.Flags().GetString(flagContAddr)
			fromBlock, _ := cmd.Flags().GetInt64(flagFromHeight)
			toBlock, _ := cmd.Flags().GetInt64(flagToHeight)

			return ExecuteUndelegation(ethEndpoint, contractAddr, fromBlock, toBlock)
		},
	}

	cmd.Flags().String(flagEthEndpoint, "https://ropsten.infura.io/v3/d0383d521441488fb754735af7fe0c59", "The ethereum https endpoint to query")
	cmd.Flags().String(flagContAddr, "0x50fCe2E7426FFfEd8762e21bdf7E0Fe9188eD54A", "The contract address to query")
	cmd.Flags().Int64(flagFromHeight, 0, "The earliest height to query")
	cmd.Flags().Int64(flagToHeight, math.MaxInt64, "The latest height to query")

	return cmd
}

func ExecuteUndelegation(ethEndpoint string, contAddr string, fromBlock, toBlock int64) error {
	ethclient, err := ethclient.Dial(ethEndpoint)
	if err != nil {
		log.Fatal(err)
	}

	// set contract addr and ABI
	// TODO: make this a parameter
	contractAddress := common.HexToAddress(contAddr)

	logUndelegateSig := []byte("Undelegate(address,uint256)")

	logUndelegateSigHash := crypto.Keccak256Hash(logUndelegateSig)

	// TODO: make this a parameter
	//contractAbi, err := abi.JSON(strings.NewReader(events.EventsMetaData.ABI))
	//if err != nil {
	//	log.Fatal(err)
	//}

	ethclient.BlockByNumber(context.Background(), nil)

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(fromBlock),
		ToBlock:   big.NewInt(toBlock),
		Addresses: []common.Address{
			contractAddress,
		},
		Topics: [][]common.Hash{
			{logUndelegateSigHash},
		},
	}

	logs, err := ethclient.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(logs)
	return nil
}

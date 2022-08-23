package cmd

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/evmos-stayking-house/scheduled-worker-golang/abis"
	"github.com/spf13/cobra"
	"log"
	"math"
	"math/big"
	"strings"
)

func ExecuteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "exec",
		Short:                      "execution subcommands",
		SuggestionsMinimumDistance: 2,
	}

	cmd.AddCommand(NewExecuteUndelegationCmd())
	// add cosmos tx related flags to be passed on while generating cosmos txs
	return cmd
}

func NewExecuteUndelegationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "process-undelegation",
		Short: `batch queries any undelegation event and initiates an undelegation for the corresponding amount.`,
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			ethEndpoint, _ := cmd.Flags().GetString(flagEthEndpoint)
			contractAddr, _ := cmd.Flags().GetString(flagContAddr)
			fromBlock, _ := cmd.Flags().GetInt64(flagFromHeight)
			toBlock, _ := cmd.Flags().GetInt64(flagToHeight)
			valString, _ := cmd.Flags().GetString(flagValidator)
			valAddr, err := sdk.ValAddressFromBech32(valString)
			udAmt, err := QueryUndelegationAmt(ethEndpoint, contractAddr, fromBlock, toBlock)
			if err != nil {
				return err
			}
			unbondingAmt := sdk.NewIntFromBigInt(udAmt)
			// TODO: refer to actual evmos denom
			unbondingCoin := sdk.NewCoin(baseDenom, unbondingAmt)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			delAddr := clientCtx.GetFromAddress()

			msg := stakingtypes.NewMsgUndelegate(delAddr, valAddr, unbondingCoin)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(flagEthEndpoint, "wss://localhost:8546", "The ethereum https endpoint to query")
	cmd.Flags().String(flagContAddr, "", "The contract address to query")
	cmd.Flags().Int64(flagFromHeight, 0, "The earliest height to query")
	cmd.Flags().Int64(flagToHeight, math.MaxInt64, "The latest height to query")
	cmd.Flags().String(flagValidator, "", "The validator to undelegate from")

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().Set(flags.FlagSkipConfirmation, "true")
	cmd.Flags().Set(flags.FlagBroadcastMode, "block")
	cmd.Flags().Set(flags.FlagKeyringBackend, "test")
	cmd.Flags().Set(flags.FlagGasAdjustment, "1.5")
	cmd.Flags().Set(flags.FlagGas, "auto")
	cmd.Flags().Set(flags.FlagGasPrices, "10000000000"+baseDenom)
	return cmd
}

func QueryUndelegationAmt(ethEndpoint, contAddr string, fromBlock, toBlock int64) (*big.Int, error) {
	log.Println("querying to ethereum endpoint: ", ethEndpoint)
	ethclient, err := ethclient.Dial(ethEndpoint)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	contractAddress := common.HexToAddress(contAddr)

	logUndelegateSig := []byte("Undelegate(address,uint256)")
	logUndelegateSigHash := crypto.Keccak256Hash(logUndelegateSig)

	// TODO: make this a parameter
	contractAbi, err := abi.JSON(strings.NewReader(abis.EventsMetaData.ABI))
	if err != nil {
		log.Fatal(err)
	}

	block, err := ethclient.BlockByNumber(ctx, nil)
	if err != nil {
		return nil, err
	}
	blockHeight := block.Number()
	fmt.Println(blockHeight)
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

	logs, err := ethclient.FilterLogs(ctx, query)
	if err != nil {
		return nil, err
	}
	var changes []DelegationChange
	totalUndelegation := big.NewInt(0)
	for _, log := range logs {
		amt, err := contractAbi.Unpack("Undelegate", log.Data)
		if err != nil {
			return nil, err
		}

		// TODO: figure out if we need the addresses
		var d DelegationChange
		d.Amount = amt[0].(*big.Int)
		d.Delegator = common.HexToAddress(log.Topics[1].String())
		changes = append(changes, d)

		totalUndelegation.Add(d.Amount, totalUndelegation)
	}
	fmt.Println(changes)
	fmt.Println(totalUndelegation.String())
	return totalUndelegation, nil
}

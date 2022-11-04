package cmd

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func ServeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "serve",
		Short:                      "Serve subcommands",
		SuggestionsMinimumDistance: 2,
	}

	cmd.AddCommand(
		NewSubscribeCmd(),
	)

	return cmd
}

func NewSubscribeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "subscribe",
		Short: `Subscribes to delegation,epoch end, unbonding events.`,
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("1here!!!")
			fmt.Println("1here!!!")

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			fmt.Println("2here!!!")

			return Subscribe(clientCtx, cmd.Flags())
		},
	}

	cmd.Flags().String(flagStaykingContAddr, "", "The stayking contract address to listen to")
	cmd.Flags().String(flagUnbondingEvmosContAddr, "", "The uEVMOS contract address to listen to")
	cmd.Flags().String(flagValidator, "", "The validator to delegate to")
	cmd.Flags().String(flagEthEndpoint, "http://localhost:8545", "The ethereum endpoint to connect to")
	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")
	// usually 2 is enough to recover from not-updated state reference, but leave some more for buffer.
	cmd.Flags().Uint64P(flagMaxRetry, "r", 10, "The number of retries to broadcast a tx")
	cmd.Flags().Uint64P(flagEthGasPrice, "g", 10000000000, "The ethereum gas price to input")

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().Set(flags.FlagSkipConfirmation, "true")
	cmd.Flags().Set(flags.FlagBroadcastMode, "block")
	cmd.Flags().Set(flags.FlagKeyringBackend, "test")
	//cmd.Flags().Set(flags.FlagGasAdjustment, "1.5")
	//cmd.Flags().Set(flags.FlagGas, "auto")
	//cmd.Flags().Set(flags.FlagGasPrices, "100000000"+baseDenom)

	return cmd
}

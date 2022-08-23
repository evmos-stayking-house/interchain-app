package cmd

import (
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
		NewSubscribeDelegationCmd(),
	)

	return cmd
}

func NewSubscribeDelegationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "subscribe",
		Short: `Subscribes to delegation,epoch end, unbonding events.`,
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			contractAddr, _ := cmd.Flags().GetString(flagContAddr)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			return Subscribe(contractAddr, clientCtx, cmd.Flags())
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

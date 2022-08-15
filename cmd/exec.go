package cmd

import "github.com/spf13/cobra"

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

			return SubscribeDelegation(ethEndpoint, contractAddr)
		},
	}

	cmd.Flags().String(flagEthEndpoint, "https://ropsten.infura.io/ws/v3/d0383d521441488fb754735af7fe0c59", "The ethereum https endpoint to query")
	cmd.Flags().String(flagContAddr, "0x50fCe2E7426FFfEd8762e21bdf7E0Fe9188eD54A", "The contract address to query")

	return cmd
}

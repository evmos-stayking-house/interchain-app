package cmd

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"

	tmcfg "github.com/tendermint/tendermint/config"
	tmcli "github.com/tendermint/tendermint/libs/cli"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "sworker",
		Short: "scheduled worker",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			// set the default command outputs
			cmd.SetOut(cmd.OutOrStdout())
			cmd.SetErr(cmd.ErrOrStderr())
			return nil
		},
	}
	initRootCmd(rootCmd)
	return rootCmd
}

// Execute is a copy of cosmos binary Execute without the server side commands
func Execute(rootCmd *cobra.Command, defaultHome string) error {
	ctx := context.Background()
	ctx = context.WithValue(ctx, client.ClientContextKey, &client.Context{})

	rootCmd.PersistentFlags().String(flags.FlagLogLevel, zerolog.InfoLevel.String(), "The logging level (trace|debug|info|warn|error|fatal|panic)")
	rootCmd.PersistentFlags().String(flags.FlagLogFormat, tmcfg.LogFormatPlain, "The logging format (json|plain)")

	executor := tmcli.PrepareBaseCmd(rootCmd, "", defaultHome)
	return executor.ExecuteContext(ctx)
}

func initRootCmd(command *cobra.Command) {
	command.AddCommand(ServeCommand())
	command.AddCommand(ExecuteCommand())
}

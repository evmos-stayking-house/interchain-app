package cmd

//// NewConvertCoinCmd returns a CLI command handler for converting a Cosmos coin
//func NewConvertCoinCmd() *cobra.Command {
//	cmd := &cobra.Command{
//		Use:   "convert-coin [coin] [receiver_hex]",
//		Short: "Convert a Cosmos coin to ERC20. When the receiver [optional] is omitted, the ERC20 tokens are transferred to the sender.",
//		Args:  cobra.RangeArgs(1, 2),
//		RunE: func(cmd *cobra.Command, args []string) error {
//			cliCtx, err := client.GetClientTxContext(cmd)
//			if err != nil {
//				return err
//			}
//
//			coin, err := sdk.ParseCoinNormalized(args[0])
//			if err != nil {
//				return err
//			}
//
//			var receiver string
//			sender := cliCtx.GetFromAddress()
//
//			if len(args) == 2 {
//				receiver = args[1]
//				if err := ethermint.ValidateAddress(receiver); err != nil {
//					return fmt.Errorf("invalid receiver hex address %w", err)
//				}
//			} else {
//				receiver = common.BytesToAddress(sender).Hex()
//			}
//
//			msg := &types.MsgConvertCoin{
//				Coin:     coin,
//				Receiver: receiver,
//				Sender:   sender.String(),
//			}
//
//			if err := msg.ValidateBasic(); err != nil {
//				return err
//			}
//
//			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
//		},
//	}
//
//	flags.AddTxFlagsToCmd(cmd)
//	return cmd
//}

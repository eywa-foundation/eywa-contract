package cli

import (
	"strconv"

	"eywa-contract/x/eywacontract/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdSendChat() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send-chat [chat] [receiver] [time]",
		Short: "Broadcast message send-chat",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argChat := args[0]
			argReceiver := args[1]
			argTime := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSendChat(
				clientCtx.GetFromAddress().String(),
				argChat,
				argReceiver,
				argTime,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

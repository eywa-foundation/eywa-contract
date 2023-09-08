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

func CmdHandshake() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "handshake [receiver] [server-address] [room-id] [time]",
		Short: "Broadcast message handshake",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argReceiver := args[0]
			argServerAddress := args[1]
			argRoomId := args[2]
			argTime := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgHandshake(
				clientCtx.GetFromAddress().String(),
				argReceiver,
				argServerAddress,
				argRoomId,
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

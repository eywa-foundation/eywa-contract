package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"eywa-contract/x/eywacontract/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group eywacontract queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdListRegisterType())
	cmd.AddCommand(CmdShowRegisterType())
	cmd.AddCommand(CmdListSendChatType())
	cmd.AddCommand(CmdShowSendChatType())
	cmd.AddCommand(CmdListHandshakeType())
	cmd.AddCommand(CmdShowHandshakeType())
	// this line is used by starport scaffolding # 1

	return cmd
}

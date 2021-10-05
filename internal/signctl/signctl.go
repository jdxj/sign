package signctl

import (
	"github.com/spf13/cobra"

	"github.com/jdxj/sign/internal/signctl/auth"
	"github.com/jdxj/sign/internal/signctl/consts"
	"github.com/jdxj/sign/internal/signctl/create"
	"github.com/jdxj/sign/internal/signctl/delete"
	"github.com/jdxj/sign/internal/signctl/update"
)

var (
	rootCmd *cobra.Command
)

func init() {
	rootCmd = NewRootCmd()
}

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "signctl",
		Aliases:                    nil,
		SuggestFor:                 nil,
		Short:                      "",
		Long:                       "",
		Example:                    "",
		ValidArgs:                  nil,
		ValidArgsFunction:          nil,
		Args:                       nil,
		ArgAliases:                 nil,
		BashCompletionFunction:     "",
		Deprecated:                 "",
		Annotations:                nil,
		Version:                    "",
		PersistentPreRun:           nil,
		PersistentPreRunE:          nil,
		PreRun:                     nil,
		PreRunE:                    nil,
		Run:                        nil,
		RunE:                       nil,
		PostRun:                    nil,
		PostRunE:                   nil,
		PersistentPostRun:          nil,
		PersistentPostRunE:         nil,
		FParseErrWhitelist:         cobra.FParseErrWhitelist{},
		CompletionOptions:          cobra.CompletionOptions{},
		TraverseChildren:           false,
		Hidden:                     false,
		SilenceErrors:              false,
		SilenceUsage:               false,
		DisableFlagParsing:         false,
		DisableAutoGenTag:          false,
		DisableFlagsInUseLine:      false,
		DisableSuggestions:         false,
		SuggestionsMinimumDistance: 0,
	}

	// flags
	flagSet := cmd.PersistentFlags()
	flagSet.StringP(consts.Host, "H", "http://127.0.0.1:8080", "apiserver address")
	flagSet.StringP(consts.Token, "T", "", "user token")

	// subcommands
	cmd.AddCommand(create.New())
	cmd.AddCommand(auth.New())
	cmd.AddCommand(update.New())
	cmd.AddCommand(delete.New())
	return cmd
}

func Execute() error {
	return rootCmd.Execute()
}

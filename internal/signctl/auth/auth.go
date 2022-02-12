package auth

import (
	"github.com/spf13/cobra"

	"github.com/jdxj/sign/internal/pkg/util"
	"github.com/jdxj/sign/internal/signctl/consts"
	"github.com/jdxj/sign/internal/signctl/model"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "auth",
		Aliases:                    nil,
		SuggestFor:                 nil,
		Short:                      "auth used for obtain token",
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
		Run:                        authCmdRun,
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
	flagSet := cmd.Flags()
	flagSet.String(consts.Nickname, "", "user nickname")
	flagSet.String(consts.Password, "", "user password")

	return cmd
}

func authCmdRun(cmd *cobra.Command, _ []string) {
	var (
		host, _     = cmd.Flags().GetString(consts.Host)
		nickname, _ = cmd.Flags().GetString(consts.Nickname)
		password, _ = cmd.Flags().GetString(consts.Password)
		err         error
	)
	if password == "" {
		password, err = util.GetPassword()
		if err != nil {
			cmd.PrintErrf("get password failed: %s", err)
			return
		}
	}

	req := &model.AuthReq{
		Nickname: nickname,
		Password: password,
	}
	rsp := &model.Response{}

	err = util.SendJson(host, req, rsp, util.WithJoin(consts.ApiToken))
	if err != nil {
		cmd.Printf("%s: %s", consts.ErrSendJson, err)
		return
	}
	cmd.Printf("%s", rsp)
}

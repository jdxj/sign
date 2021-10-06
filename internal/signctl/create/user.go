package create

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/jdxj/sign/internal/pkg/util"
	"github.com/jdxj/sign/internal/signctl/consts"
	"github.com/jdxj/sign/internal/signctl/model"
)

func newUserCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "user",
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
		Run:                        userCmdRun,
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

func userCmdRun(cmd *cobra.Command, args []string) {
	var (
		host     = cmd.Flag(consts.Host)
		nickname = cmd.Flag(consts.Nickname)
		passFlag = cmd.Flag(consts.Password)
		pass     = passFlag.Value.String()
		err      error
	)

	if pass == "" {
		pass, err = util.GetPassword()
		if err != nil {
			cmd.PrintErrf("get password failed: %s", err)
			return
		}
	}

	url := fmt.Sprintf("%s%s",
		strings.TrimSuffix(host.Value.String(), "/"), consts.CreateUser)
	req := &model.CreateUserReq{
		Nickname: nickname.Value.String(),
		Password: pass,
	}
	rsp := &model.Response{}

	err = util.PostJson(url, req, rsp)
	if err != nil {
		cmd.Printf("%s: %s", consts.ErrSendJson, err)
		return
	}

	cmd.Printf("%s\n", rsp)
}

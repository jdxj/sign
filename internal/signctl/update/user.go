package update

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
	flagSet.String(consts.Mail, "", "user email address")
	flagSet.Int64(consts.Telegram, 0, "user telegram id")
	return cmd
}

func userCmdRun(cmd *cobra.Command, args []string) {
	host := cmd.Flag(consts.Host)
	token := cmd.Flag(consts.Token)

	nickname, _ := cmd.Flags().GetString(consts.Nickname)
	password, _ := cmd.Flags().GetString(consts.Password)
	mail, _ := cmd.Flags().GetString(consts.Mail)
	telegram, _ := cmd.Flags().GetInt64(consts.Telegram)

	url := fmt.Sprintf("%s%s",
		strings.TrimSuffix(host.Value.String(), "/"), consts.UserUpdate)

	req := &model.Request{
		Token: token.Value.String(),
		Data: &model.UpdateUserReq{
			Nickname: nickname,
			Password: password,
			Mail:     mail,
			Telegram: telegram,
		},
	}
	err := util.PutJson(url, req, nil)
	if err != nil {
		cmd.PrintErrf("%s: put, %s\n", consts.ErrSendJson, err)
		return
	}

	cmd.Println("update user successfully")
}

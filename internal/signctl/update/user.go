package update

import (
	"net/http"

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

func userCmdRun(cmd *cobra.Command, _ []string) {
	var (
		host, _     = cmd.Flags().GetString(consts.Host)
		t, _        = cmd.Flags().GetString(consts.Token)
		debug, _    = cmd.Flags().GetBool(consts.Debug)
		nickname, _ = cmd.Flags().GetString(consts.Nickname)
		password, _ = cmd.Flags().GetString(consts.Password)
		mail, _     = cmd.Flags().GetString(consts.Mail)
		telegram, _ = cmd.Flags().GetInt64(consts.Telegram)
	)
	req := &model.UpdateUserReq{
		Nickname: nickname,
		Password: password,
		Mail:     mail,
		Telegram: telegram,
	}
	rsp := &model.Response{}

	err := util.SendJson(
		host,
		req,
		rsp,
		util.WithDebug(debug),
		util.WithJoin(consts.ApiUser),
		util.WithMethod(http.MethodPut),
		util.WithBearer(t),
	)
	if err != nil {
		cmd.PrintErrf("%s: put, %s\n", consts.ErrSendJson, err)
		return
	}
	cmd.Printf("%s\n", rsp)
}

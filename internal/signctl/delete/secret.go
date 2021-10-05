package delete

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/jdxj/sign/internal/pkg/util"
	"github.com/jdxj/sign/internal/signctl/consts"
	"github.com/jdxj/sign/internal/signctl/model"
)

func newSecretCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "secret",
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
		Run:                        secretCmdRun,
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
	flagSet.Int64(consts.SecretID, 0, "secret id")
	return cmd
}

func secretCmdRun(cmd *cobra.Command, args []string) {
	host := cmd.Flag(consts.Host)
	token := cmd.Flag(consts.Token)

	secretID := cmd.Flag(consts.SecretID)

	secretIDInt64, err := strconv.ParseInt(secretID.Value.String(), 10, 64)
	if err != nil {
		cmd.PrintErrf("%s: secret-id: %s", consts.ErrInvalidParam, secretID.Value)
		return
	}

	req := &model.Request{
		Token: token.Value.String(),
		Data: &model.DeleteSecretReq{
			SecretID: secretIDInt64,
		},
	}
	url := fmt.Sprintf("%s%s",
		strings.TrimSuffix(host.Value.String(), "/"), consts.DeleteSecret)

	err = util.DeleteJson(url, req, nil)
	if err != nil {
		cmd.PrintErrf("%s: delete, %s", consts.ErrSendJson, err)
		return
	}

	cmd.Println("delete secret successfully")
}

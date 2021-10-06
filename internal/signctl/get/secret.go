package get

import (
	"fmt"
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
	secretIDsSecret = flagSet.Int64Slice(consts.SecretID, nil, "specify multiple secret id for query")
	domains = flagSet.IntSlice(consts.Domain, nil, "specify multiple domain for query")
	return cmd
}

var (
	secretIDsSecret *[]int64
	domains         *[]int
)

func secretCmdRun(cmd *cobra.Command, args []string) {
	host := cmd.Flag(consts.Host)
	token := cmd.Flag(consts.Token)

	url := fmt.Sprintf("%s%s",
		strings.TrimSuffix(host.Value.String(), "/"), consts.GetSecrets)
	req := &model.Request{
		Token: token.Value.String(),
		Data: &model.GetSecretsReq{
			SecretIDs: *secretIDsSecret,
			Domains:   *domains,
		},
	}
	rsp := &model.Response{}

	err := util.PostJson(url, req, rsp)
	if err != nil {
		cmd.PrintErrf("%s: post, %s", consts.ErrSendJson, err)
		return
	}

	cmd.Printf("%s\n", rsp)
}

package create

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/jdxj/sign/internal/pkg/util"
	"github.com/jdxj/sign/internal/signctl/consts"
	"github.com/jdxj/sign/internal/signctl/help"
	"github.com/jdxj/sign/internal/signctl/model"
)

func newTaskCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "task",
		Aliases:                    nil,
		SuggestFor:                 nil,
		Short:                      "",
		Long:                       help.AvailableKind(),
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
		Run:                        taskCmdRun,
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
	flagSet.String(consts.Describe, "", "description of the task")
	flagSet.Int(consts.Kind, 0, "kind of the task")
	flagSet.String(consts.Spec, "", "crontab expression")
	flagSet.Int64(consts.SecretID, 0, "the id of the secret used by the task")
	return cmd
}

func taskCmdRun(cmd *cobra.Command, args []string) {
	host := cmd.Flag(consts.Host)
	token := cmd.Flag(consts.Token)

	describe := cmd.Flag(consts.Describe)
	kind := cmd.Flag(consts.Kind)
	spec := cmd.Flag(consts.Spec)
	secretID := cmd.Flag(consts.SecretID)

	kindInt, err := strconv.Atoi(kind.Value.String())
	if err != nil {
		cmd.PrintErrf("%s, kind: %s", consts.ErrInvalidParam, kind.Value.String())
		return
	}
	secretIDInt64, err := strconv.ParseInt(secretID.Value.String(), 10, 64)
	if err != nil {
		cmd.PrintErrf("%s, secretID: %s", consts.ErrInvalidParam, secretID.Value.String())
		return
	}

	url := fmt.Sprintf("%s%s",
		strings.TrimSuffix(host.Value.String(), "/"), consts.CreateTask)

	req := &model.Request{
		Token: token.Value.String(),
		Data: &model.CreateTaskReq{
			Describe: describe.Value.String(),
			Kind:     kindInt,
			Spec:     spec.Value.String(),
			SecretID: secretIDInt64,
		},
	}
	rsp := &model.Response{}
	err = util.PostJson(url, req, rsp)
	if err != nil {
		cmd.PrintErrf("%s: post, %s\n", consts.ErrSendJson, err)
		return
	}

	cmd.Printf("%s\n", rsp)
}

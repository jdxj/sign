package create

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/jdxj/sign/internal/pkg/util"
	"github.com/jdxj/sign/internal/proto/task"
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
	flagSet.String(consts.Description, "", "description of the task")
	flagSet.Int32(consts.Kind, 0, "kind of the task")
	flagSet.String(consts.Spec, "", "crontab expression")
	flagSet.StringSliceP(consts.Param, "p", nil, "the param required by the task can be specified multiple times (key=value)")
	return cmd
}

func taskCmdRun(cmd *cobra.Command, args []string) {
	host := cmd.Flag(consts.Host)
	token := cmd.Flag(consts.Token)
	describe := cmd.Flag(consts.Description)
	kind, _ := cmd.Flags().GetInt32(consts.Kind)
	spec := cmd.Flag(consts.Spec)
	params, err := cmd.Flags().GetStringSlice(consts.Param)
	if err != nil {
		cmd.PrintErrf("%s, params: %s", consts.ErrInvalidParam, cmd.Flag(consts.Param).Value)
		return
	}

	url := fmt.Sprintf("%s%s",
		strings.TrimSuffix(host.Value.String(), "/"), consts.TaskCreate)

	paramMap := make(map[string]string)
	for _, param := range params {
		pair := strings.Split(param, "=")
		if len(pair) != 2 {
			cmd.PrintErrf("%s, param: %s", consts.ErrInvalidParam, param)
			return
		}
		paramMap[pair[0]] = pair[1]
	}
	param, _ := json.Marshal(paramMap)

	req := &model.Request{
		Token: token.Value.String(),
		Data: &model.CreateTaskReq{
			Describe: describe.Value.String(),
			Kind:     task.Kind_name[kind],
			Spec:     spec.Value.String(),
			Param:    param,
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

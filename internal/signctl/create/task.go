package create

import (
	"encoding/json"
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

func taskCmdRun(cmd *cobra.Command, _ []string) {
	var (
		host, _     = cmd.Flags().GetString(consts.Host)
		t, _        = cmd.Flags().GetString(consts.Token)
		debug, _    = cmd.Flags().GetBool(consts.Debug)
		describe, _ = cmd.Flags().GetString(consts.Description)
		kind, _     = cmd.Flags().GetInt32(consts.Kind)
		spec, _     = cmd.Flags().GetString(consts.Spec)
		params, _   = cmd.Flags().GetStringSlice(consts.Param)
	)
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

	req := &model.CreateTaskReq{
		Describe: describe,
		Kind:     task.Kind_name[kind],
		Spec:     spec,
		Param:    param,
	}
	rsp := &model.Response{}

	err := util.SendJson(
		host,
		req,
		rsp,
		util.WithDebug(debug),
		util.WithJoin(consts.ApiTasks),
		util.WithBearer(t),
	)
	if err != nil {
		cmd.PrintErrf("%s: post, %s\n", consts.ErrSendJson, err)
		return
	}
	cmd.Printf("%s\n", rsp)
}

package update

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/spf13/cobra"

	"github.com/jdxj/sign/internal/pkg/util"
	"github.com/jdxj/sign/internal/signctl/consts"
	"github.com/jdxj/sign/internal/signctl/model"
)

func newTaskCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "task",
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
	flagSet.String(consts.TaskID, "", "task id")
	flagSet.String(consts.Description, "", "task description")
	flagSet.String(consts.Spec, "", "task spec")
	flagSet.StringSliceP(consts.Param, "p", nil, "task params")
	return cmd
}

func taskCmdRun(cmd *cobra.Command, _ []string) {
	var (
		host, _   = cmd.Flags().GetString(consts.Host)
		debug, _  = cmd.Flags().GetBool(consts.Debug)
		token, _  = cmd.Flags().GetString(consts.Token)
		taskID, _ = cmd.Flags().GetString(consts.TaskID)
		desc, _   = cmd.Flags().GetString(consts.Description)
		spec, _   = cmd.Flags().GetString(consts.Spec)
		params, _ = cmd.Flags().GetStringSlice(consts.Param)
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

	req := &model.UpdateTaskReq{
		Desc:  desc,
		Spec:  spec,
		Param: param,
	}
	rsp := &model.Response{}

	err := util.SendJson(
		host,
		req,
		rsp,
		util.WithDebug(debug),
		util.WithJoin(consts.ApiTasks),
		util.WithJoin(taskID),
		util.WithMethod(http.MethodPut),
		util.WithBearer(token),
	)
	if err != nil {
		cmd.PrintErrf("%s: put, %s\n", consts.ErrSendJson, err)
		return
	}
	cmd.Printf("%s\n", rsp)
}

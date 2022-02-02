package update

import (
	"encoding/json"
	"fmt"
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
	flagSet.Int64(consts.TaskID, 0, "task id")
	flagSet.String(consts.Description, "", "task description")
	flagSet.String(consts.Spec, "", "task spec")
	flagSet.StringSlice(consts.Param, nil, "task params")
	return cmd
}

func taskCmdRun(cmd *cobra.Command, args []string) {
	host := cmd.Flag(consts.Host)
	token := cmd.Flag(consts.Token)
	taskID, _ := cmd.Flags().GetInt64(consts.TaskID)
	desc, _ := cmd.Flags().GetString(consts.Description)
	spec, _ := cmd.Flags().GetString(consts.Spec)
	params, err := cmd.Flags().GetStringSlice(consts.Param)
	if err != nil {
		cmd.PrintErrf("%s, params: %s", consts.ErrInvalidParam, cmd.Flag(consts.Param).Value)
		return
	}

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

	url := fmt.Sprintf("%s%s",
		strings.TrimSuffix(host.Value.String(), "/"), consts.TaskUpdate)

	req := &model.Request{
		Token: token.Value.String(),
		Data: &model.UpdateTaskReq{
			TaskID: taskID,
			Desc:   desc,
			Spec:   spec,
			Param:  param,
		},
	}

	err = util.PutJson(url, req, nil)
	if err != nil {
		cmd.PrintErrf("%s: put, %s\n", consts.ErrSendJson, err)
		return
	}

	cmd.Println("update task successfully")
}

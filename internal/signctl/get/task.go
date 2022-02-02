package get

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/jdxj/sign/internal/pkg/util"
	"github.com/jdxj/sign/internal/proto/task"
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
	flagSet.Int32(consts.Kind, 0, "task kind")
	flagSet.String(consts.Spec, "", "task spec")
	flagSet.Int64(consts.CreatedAt, 0, "create timestamp")
	flagSet.Int64(consts.PageID, 1, "page id")
	flagSet.Int64(consts.PageSize, 10, "page size")
	return cmd
}

func taskCmdRun(cmd *cobra.Command, args []string) {
	host := cmd.Flag(consts.Host)
	token := cmd.Flag(consts.Token)
	taskID, _ := cmd.Flags().GetInt64(consts.TaskID)
	desc, _ := cmd.Flags().GetString(consts.Description)
	kind, _ := cmd.Flags().GetInt32(consts.Kind)
	spec, _ := cmd.Flags().GetString(consts.Spec)
	createdAt, _ := cmd.Flags().GetInt64(consts.CreatedAt)
	pageID, _ := cmd.Flags().GetInt64(consts.PageID)
	pageSize, _ := cmd.Flags().GetInt64(consts.PageSize)

	req := &model.Request{
		Token: token.Value.String(),
		Data: &model.GetTasksReq{
			TaskID:    taskID,
			Desc:      desc,
			Kind:      task.Kind_name[kind],
			Spec:      spec,
			CreatedAt: createdAt,
			PageID:    pageID,
			PageSize:  pageSize,
		},
	}
	rsp := &model.Response{}
	url := fmt.Sprintf("%s%s",
		strings.TrimSuffix(host.Value.String(), "/"), consts.TaskList)

	err := util.PostJson(url, req, rsp)
	if err != nil {
		cmd.PrintErrf("%s: post, %s", consts.ErrSendJson, err)
		return
	}

	cmd.Printf("%s\n", rsp)
}

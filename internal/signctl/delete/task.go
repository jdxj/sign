package delete

import (
	"net/http"

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
	flagSet.String(consts.TaskID, "", "the id of task")
	return cmd
}

func taskCmdRun(cmd *cobra.Command, _ []string) {
	var (
		host, _   = cmd.Flags().GetString(consts.Host)
		debug, _  = cmd.Flags().GetBool(consts.Debug)
		token, _  = cmd.Flags().GetString(consts.Token)
		taskID, _ = cmd.Flags().GetString(consts.TaskID)
	)
	rsp := &model.Response{}

	err := util.SendJson(
		host,
		nil,
		rsp,
		util.WithDebug(debug),
		util.WithJoin(consts.ApiTasks),
		util.WithJoin(taskID),
		util.WithMethod(http.MethodDelete),
		util.WithBearer(token),
	)
	if err != nil {
		cmd.PrintErrf("%s: delete, %s", consts.ErrSendJson, err)
		return
	}
	cmd.Printf("%s\n", rsp)
}

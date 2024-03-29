package ref

import (
	"go-micro.dev/v4/client"

	"github.com/jdxj/sign/internal/proto/notice"
	"github.com/jdxj/sign/internal/proto/task"
	"github.com/jdxj/sign/internal/proto/trigger"
	"github.com/jdxj/sign/internal/proto/user"
)

var (
	UserService    user.UserService
	TaskService    task.TaskService
	NoticeService  notice.NoticeService
	TriggerService trigger.TriggerService
)

func Init(cc client.Client) {
	UserService = user.NewUserService(user.ServiceName, cc)
	TaskService = task.NewTaskService(task.ServiceName, cc)
	NoticeService = notice.NewNoticeService(notice.ServiceName, cc)
	TriggerService = trigger.NewTriggerService(trigger.ServiceName, cc)
}

package trigger

import (
	"github.com/panjf2000/ants/v2"
	"google.golang.org/protobuf/proto"

	"github.com/jdxj/sign/internal/pkg/logger"
	"github.com/jdxj/sign/internal/pkg/mq"
	"github.com/jdxj/sign/internal/proto/crontab"
	"github.com/jdxj/sign/internal/trigger/dao/task"
)

type job struct {
	spec  string
	gPool *ants.Pool
	tq    *mq.TaskQueue
}

func (j *job) Run() {
	err := j.gPool.Submit(j.sendTaskToMq)
	if err != nil {
		logger.Errorf("submit func failed, err: %s", err)
	}
}

func (j *job) sendTaskToMq() {
	where := map[string]interface{}{
		"spec = ?": j.spec,
	}
	rows, err := task.Find(where)
	if err != nil {
		logger.Errorf("get tasks failed: %s", err)
		return
	}

	for _, row := range rows {
		t := &crontab.Task{
			TaskID:   row.TaskID,
			UserID:   row.UserID,
			Kind:     crontab.TaskKind(row.Kind),
			Spec:     row.Spec,
			SecretID: row.SecretID,
		}
		body, err := proto.Marshal(t)
		if err != nil {
			logger.Warnf("proto marshal failed: %+v", t)
			continue
		}

		err = j.tq.Publish(body)
		if err != nil {
			logger.Errorf("publish failed, task id: %d", t.TaskID)
		} else {
			logger.Debugf("publish task %d success", t.TaskID)
		}
	}
}

package model

import (
	"context"
	"time"

	"go-micro.dev/v4/broker"
	"google.golang.org/protobuf/proto"

	"github.com/jdxj/sign/internal/pkg/db"
	"github.com/jdxj/sign/internal/pkg/logger"
	"github.com/jdxj/sign/internal/pkg/util"
	pb "github.com/jdxj/sign/internal/proto/task"
	"github.com/jdxj/sign/internal/task/client"
)

type dispatchTask struct {
	TaskID int64 `gorm:"primaryKey"`
	UserID int64
	Kind   string
	Spec   string
	Param  []byte
}

func (dt *dispatchTask) TableName() string {
	return "task"
}

func NewJob(key []byte, spec string) *Job {
	return &Job{
		key:  key,
		spec: spec,
	}
}

type Job struct {
	key  []byte
	spec string
}

func (j *Job) DispatchTasks() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// todo: 分批
	var rows []dispatchTask
	err := db.WithCtx(ctx).
		Where("spec = ?", j.spec).
		Find(&rows).
		Error
	if err != nil {
		logger.Errorf("Find: %s", err)
		return
	}

	for _, row := range rows {
		task := &pb.Task{
			TaskId: row.TaskID,
			UserId: row.UserID,
			Kind:   row.Kind,
			Spec:   row.Spec,
			Param:  util.Decrypt(j.key, row.Param),
		}
		body, err := proto.Marshal(task)
		if err != nil {
			logger.Errorf("Marshal: %s", err)
			continue
		}
		err = client.MQ.Publish(pb.Topic, &broker.Message{Body: body})
		if err != nil {
			logger.Errorf("Publish: %s", err)
			continue
		}
	}
}

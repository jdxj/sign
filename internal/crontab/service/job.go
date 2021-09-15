package service

import (
	"fmt"

	"github.com/panjf2000/ants/v2"

	"github.com/jdxj/sign/internal/pkg/logger"
)

type job struct {
	specID int64
	gPool  *ants.Pool
}

func (j *job) Run() {
	err := j.gPool.Submit(func() {
		SendTaskToMq(j.specID)
	})
	if err != nil {
		logger.Errorf("submit func failed, err: %s", err)
	}
}

func SendTaskToMq(specID int64) {
	// todo:
	//   实现: 不断从数据库中读取任务,
	//        并将任务信息发送到 RabbitMq.
	fmt.Printf("send data to mq ok: %d\n", specID)
}

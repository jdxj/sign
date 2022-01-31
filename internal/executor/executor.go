package executor

import (
	"context"
	"net/rpc"
	"sync"
	"time"

	"github.com/panjf2000/ants/v2"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"

	"github.com/jdxj/sign/internal/pkg/logger"
	"github.com/jdxj/sign/internal/proto/notice"
	"github.com/jdxj/sign/internal/proto/user"
)

func New() *Executor {
	e := &Executor{
		wg: &sync.WaitGroup{},
	}
	gPool, err := ants.NewPool(1000)
	if err != nil {
		panic(err)
	}
	e.gPool = gPool

	tq, err := mq.NewTaskQueue()
	if err != nil {
		panic(err)
	}
	e.tq = tq

	rpc.NewClient(user.ServiceName, func(cc *grpc.ClientConn) {
		e.userClient = user.NewUserServiceClient(cc)
	})
	rpc.NewClient(secret.ServiceName, func(cc *grpc.ClientConn) {
		e.secretClient = secret.NewSecretServiceClient(cc)
	})
	rpc.NewClient(notice.ServiceName, func(cc *grpc.ClientConn) {
		e.noticeClient = notice.NewNoticeServiceClient(cc)
	})
	return e
}

type Executor struct {
	gPool *ants.Pool
	tq    *mq.TaskQueue
	wg    *sync.WaitGroup

	userClient   user.UserServiceClient
	secretClient secret.SecretServiceClient
	noticeClient notice.NoticeServiceClient
}

func (e *Executor) Start() {
	e.wg.Add(1)
	go func() {
		defer e.wg.Done()

		dataChan, err := e.tq.Consume()
		if err != nil {
			logger.Errorf("get data chan failed: %s", err)
			return
		}

		for data := range dataChan {
			t := &crontab.Task{}
			err := proto.Unmarshal(data, t)
			if err != nil {
				logger.Warnf("unmarshal task failed: %v", data)
				continue
			}

			err = e.gPool.Submit(func() {
				e.start(t)
			})
			if err != nil {
				logger.Warnf("submit task failed: %s", err)
			}
		}
	}()
}

func (e *Executor) Stop() {
	e.tq.Stop()
	e.gPool.Release()

	e.wg.Wait()
}

func (e *Executor) start(task *crontab.Task) {
	agent, ok := agents[task.Kind]
	if !ok {
		logger.Warnf("agent not register, kind: %d", task.Kind)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	secretList, err := e.secretClient.GetSecretList(ctx, &secret.GetSecretListReq{
		SecretIDs: []int64{task.SecretID},
		UserID:    task.UserID,
	})
	if err != nil {
		logger.Errorf("get secret list failed: %s", err)
		return
	}
	if len(secretList.List) < 1 {
		logger.Warnf("secret not found, taskID: %d, userID: %d, secretID: %d",
			task.TaskID, task.UserID, task.SecretID)
		return
	}
	secretRecord := secretList.List[0]
	e.tryExecute(agent, task, secretRecord.Key)
}

func (e *Executor) tryExecute(agent Agent, task *crontab.Task, key string) {
	var (
		retry    = 3
		interval = 3 * time.Second

		text string
		err  error
	)
	for i := 0; i < retry; i++ {
		text, err = agent.Execute(key)
		if err != nil {
			logger.Errorf("try execute failed: %d, userID: %d, taskID: %d, error: %s",
				i, task.UserID, task.TaskID, err)
		} else {
			break
		}
		// 最后一个不用等
		if i != retry-1 {
			time.Sleep(interval)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if text != "" {
		_, err = e.noticeClient.SendMessage(ctx, &notice.SendMessageReq{
			UserID: task.UserID,
			Text:   text,
		})
		if err != nil {
			logger.Errorf("send message failed, error: %s", err)
		}
	}
}

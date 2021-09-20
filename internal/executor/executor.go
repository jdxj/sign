package executor

import (
	"context"
	"sync"
	"time"

	"github.com/panjf2000/ants/v2"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"

	"github.com/jdxj/sign/internal/pkg/logger"
	"github.com/jdxj/sign/internal/pkg/mq"
	"github.com/jdxj/sign/internal/pkg/rpc"
	"github.com/jdxj/sign/internal/proto/crontab"
	"github.com/jdxj/sign/internal/proto/secret"
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
	return e
}

type Executor struct {
	gPool *ants.Pool
	tq    *mq.TaskQueue

	userClient   user.UserServiceClient
	secretClient secret.SecretServiceClient

	wg *sync.WaitGroup
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
	agent, ok := agents[int(task.Kind)]
	if !ok {
		logger.Warnf("agent not register, kind: %d", task.Kind)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// todo: 失败时发送通知
	secretRsp, err := e.secretClient.GetSecret(ctx, &secret.GetSecretReq{
		SecretID: task.SecretID,
	})
	if err != nil {
		logger.Errorf("get secret failed: %s", err)
		return
	}

	err = agent.SignIn(secretRsp.Key)
	if err != nil {
		logger.Errorf("sign failed: %s", err)
	}
}

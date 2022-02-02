package service

import (
	"fmt"
	"log"

	"github.com/asim/go-micro/plugins/broker/rabbitmq/v4"
	"github.com/panjf2000/ants/v2"
	"go-micro.dev/v4/broker"
	"go-micro.dev/v4/client"

	"github.com/jdxj/sign/internal/pkg/config"
	"github.com/jdxj/sign/internal/proto/notice"
	pb "github.com/jdxj/sign/internal/proto/task"
	"github.com/jdxj/sign/internal/proto/trigger"
)

var (
	triggerService trigger.TriggerService
	noticeService  notice.NoticeService

	gPool *ants.Pool
	mq    broker.Broker
)

func Init(cc client.Client, conf config.Rabbit) error {
	log.Printf(" init other service\n")
	// grpc client
	triggerService = trigger.NewTriggerService(trigger.ServiceName, cc)
	noticeService = notice.NewNoticeService(notice.ServiceName, cc)

	// broker
	mq = rabbitmq.NewBroker()
	endpoint := fmt.Sprintf("amqp://%s:%s@%s:%d", conf.User, conf.Pass, conf.Host, conf.Port)
	err := mq.Init(broker.Addrs(endpoint))
	if err != nil {
		return err
	}
	err = mq.Connect()
	if err != nil {
		return err
	}
	_, err = mq.Subscribe(pb.Topic, execute, broker.Queue(pb.Queue), broker.DisableAutoAck())
	if err != nil {
		return err
	}

	// gPool
	gPool, err = ants.NewPool(100)
	return err
}

func Close() error {
	log.Printf("stop other service\n")
	err := mq.Disconnect()
	if err != nil {
		return err
	}
	gPool.Release()
	return nil
}

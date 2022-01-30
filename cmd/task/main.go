package main

import (
	"fmt"
	"log"

	"github.com/asim/go-micro/plugins/broker/rabbitmq/v4"
	"github.com/asim/go-micro/plugins/registry/etcd/v4"
	"github.com/panjf2000/ants/v2"
	"github.com/urfave/cli/v2"
	"go-micro.dev/v4"
	"go-micro.dev/v4/broker"
	"go-micro.dev/v4/registry"

	"github.com/jdxj/sign/configs"
	"github.com/jdxj/sign/internal/pkg/config"
	"github.com/jdxj/sign/internal/pkg/db"
	"github.com/jdxj/sign/internal/pkg/logger"
	"github.com/jdxj/sign/internal/pkg/util"
	"github.com/jdxj/sign/internal/proto/notice"
	pb "github.com/jdxj/sign/internal/proto/task"
	"github.com/jdxj/sign/internal/proto/trigger"
	"github.com/jdxj/sign/internal/task/client"
	"github.com/jdxj/sign/internal/task/model/executor"
	impl "github.com/jdxj/sign/internal/task/service"
)

func main() {
	var root config.Root

	service := micro.NewService(
		micro.Name(pb.ServiceName),
		micro.Registry(etcd.NewRegistry()),
	)

	service.Init(
		micro.Action(func(cli *cli.Context) (err error) {
			path := cli.String("config")
			if path == "" {
				return fmt.Errorf("config not found")
			}
			log.Printf(" config path:[%s]\n", path)

			root = config.ReadConfigs(path)

			err = service.Options().
				Registry.Init(
				registry.Addrs(root.Etcd.Endpoints...),
				registry.TLSConfig(
					util.NewTLSConfig(root.Etcd.Ca, root.Etcd.Cert, root.Etcd.Key),
				),
			)
			if err != nil {
				return
			}

			err = db.InitGorm(root.DB)
			if err != nil {
				return
			}

			logger.Init("")
			return nil
		}),
	)

	client.TriggerService = trigger.NewTriggerService(trigger.ServiceName, service.Client())
	client.NoticeService = notice.NewNoticeService(notice.ServiceName, service.Client())

	var err error
	client.GPool, err = ants.NewPool(100)
	if err != nil {
		log.Fatalln(err)
	}

	client.MQ = rabbitmq.NewBroker()
	err = client.MQ.Init(broker.Addrs(configs.RabbitMQEndpoint))
	if err != nil {
		log.Fatalln(err)
	}
	err = client.MQ.Connect()
	if err != nil {
		log.Fatalln(err)
	}
	_, err = client.MQ.Subscribe(pb.Topic, executor.Execute, broker.Queue(pb.Queue))

	err = pb.RegisterTaskServiceHandler(service.Server(), impl.New(root.Secret))
	if err != nil {
		log.Fatalln(err)
	}
	if err := service.Run(); err != nil {
		log.Fatalln(err)
	}
}

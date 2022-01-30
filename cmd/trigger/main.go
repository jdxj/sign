package main

import (
	"fmt"
	"log"

	"github.com/asim/go-micro/plugins/registry/etcd/v4"
	"github.com/urfave/cli/v2"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"

	"github.com/jdxj/sign/internal/pkg/config"
	"github.com/jdxj/sign/internal/pkg/db"
	"github.com/jdxj/sign/internal/pkg/logger"
	"github.com/jdxj/sign/internal/pkg/util"
	"github.com/jdxj/sign/internal/proto/task"
	pb "github.com/jdxj/sign/internal/proto/trigger"
	impl "github.com/jdxj/sign/internal/trigger/service"
)

const (
	serviceName = "trigger"
)

func main() {
	service := micro.NewService(
		micro.Name(serviceName),
		micro.Registry(etcd.NewRegistry()),
	)

	service.Init(
		micro.Action(func(cli *cli.Context) (err error) {
			path := cli.String("config")
			if path == "" {
				return fmt.Errorf("config not found")
			}
			log.Printf(" config path:[%s]\n", path)

			root := config.ReadConfigs(path)

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

	// todo: const service name
	impl.TaskService = task.NewTaskService("task", service.Client())

	iService := impl.New()
	err := iService.Init()
	if err != nil {
		log.Fatalln(err)
	}

	err = pb.RegisterTriggerServiceHandler(service.Server(), iService)
	if err != nil {
		log.Fatalln(err)
	}

	if err := service.Run(); err != nil {
		log.Printf("Run: %s\n", err)
	}
}

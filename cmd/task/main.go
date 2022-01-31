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
	pb "github.com/jdxj/sign/internal/proto/task"
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

		micro.BeforeStart(func() error {
			return impl.Init(service.Client(), root.Rabbit)
		}),
		micro.AfterStop(func() error {
			return impl.Close()
		}),
	)

	err := pb.RegisterTaskServiceHandler(service.Server(), impl.New(root.Secret))
	if err != nil {
		log.Fatalln(err)
	}
	if err := service.Run(); err != nil {
		log.Fatalln(err)
	}
}
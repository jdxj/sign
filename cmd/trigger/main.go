package main

import (
	"errors"
	"log"

	"github.com/asim/go-micro/plugins/registry/etcd/v4"
	"github.com/urfave/cli/v2"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"

	"github.com/jdxj/sign/internal/pkg/config"
	"github.com/jdxj/sign/internal/pkg/db"
	"github.com/jdxj/sign/internal/pkg/logger"
	"github.com/jdxj/sign/internal/pkg/util"
	pb "github.com/jdxj/sign/internal/proto/trigger"
	impl "github.com/jdxj/sign/internal/trigger/service"
)

var (
	ErrConfigNotFound = errors.New("config not found")
)

func main() {
	service := micro.NewService(
		micro.Name(pb.ServiceName),
		micro.Registry(etcd.NewRegistry()),
	)

	service.Init(
		micro.Action(func(cli *cli.Context) (err error) {
			path := cli.String("config")
			if path == "" {
				return ErrConfigNotFound
			}

			root := config.ReadConfigs(path)
			logger.Init(root.Logger.Path, pb.ServiceName)

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

			return db.InitGorm(root.DB)
		}),

		micro.BeforeStart(func() error {
			return impl.Init(service.Client())
		}),
	)

	srv := impl.New()
	err := pb.RegisterTriggerServiceHandler(service.Server(), srv)
	if err != nil {
		log.Fatalln(err)
	}

	if err := service.Run(); err != nil {
		log.Printf("Run: %s\n", err)
	}
	srv.Stop()
}

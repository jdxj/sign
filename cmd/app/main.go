package main

import (
	"fmt"
	"log"

	"github.com/asim/go-micro/plugins/registry/etcd/v4"
	"github.com/urfave/cli/v2"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"

	"github.com/jdxj/sign/internal/app/api"
	impl "github.com/jdxj/sign/internal/app/service"
	"github.com/jdxj/sign/internal/pkg/config"
	"github.com/jdxj/sign/internal/pkg/logger"
	"github.com/jdxj/sign/internal/pkg/util"
	pb "github.com/jdxj/sign/internal/proto/app"
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
			logger.Init("")

			return service.Options().
				Registry.Init(
				registry.Addrs(root.Etcd.Endpoints...),
				registry.TLSConfig(
					util.NewTLSConfig(root.Etcd.Ca, root.Etcd.Cert, root.Etcd.Key),
				),
			)
		}),

		micro.BeforeStart(func() error {
			api.Init(service.Client(), root.APIServer.Key)
			return nil
		}),
	)

	srv := impl.New(root.APIServer)
	err := pb.RegisterAppServiceHandler(service.Server(), srv)
	if err != nil {
		log.Fatalln(err)
	}

	if err := service.Run(); err != nil {
		log.Printf("Run: %s", err)
	}

	srv.Stop()
}

package main

import (
	"errors"
	"log"

	"github.com/asim/go-micro/plugins/registry/etcd/v4"
	"github.com/urfave/cli/v2"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"

	impl "github.com/jdxj/sign/internal/notice/service"
	"github.com/jdxj/sign/internal/pkg/config"
	"github.com/jdxj/sign/internal/pkg/logger"
	"github.com/jdxj/sign/internal/pkg/util"
	pb "github.com/jdxj/sign/internal/proto/notice"
	"github.com/jdxj/sign/internal/proto/user"
)

var (
	ErrConfigNotFound = errors.New("config not found")
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
				return ErrConfigNotFound
			}

			root = config.ReadConfigs(path)
			logger.Init(root.Logger.Path, pb.ServiceName)

			return service.Options().
				Registry.Init(
				registry.Addrs(root.Etcd.Endpoints...),
				registry.TLSConfig(
					util.NewTLSConfig(root.Etcd.Ca, root.Etcd.Cert, root.Etcd.Key),
				),
			)
		}),
	)

	impl.UserService = user.NewUserService(user.ServiceName, service.Client())

	err := pb.RegisterNoticeServiceHandler(service.Server(), impl.New(root.Bot))
	if err != nil {
		log.Fatalln(err)
	}

	if err := service.Run(); err != nil {
		log.Fatalln(err)
	}
}

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/asim/go-micro/plugins/registry/etcd/v4"
	"github.com/urfave/cli/v2"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"

	"github.com/jdxj/sign/internal/pkg/config"
	"github.com/jdxj/sign/internal/pkg/util"
	testPB "github.com/jdxj/sign/internal/proto/test-grpc"
)

func main() {
	service := micro.NewService(
		micro.Name("test-grpc-client"),
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

			return service.Options().
				Registry.Init(
				registry.Addrs(root.Etcd.Endpoints...),
				registry.TLSConfig(
					util.NewTLSConfig(root.Etcd.Ca, root.Etcd.Cert, root.Etcd.Key),
				),
			)
		}),
	)

	client := testPB.NewTestRPCService("test-grpc", service.Client())
	rsp, err := client.Hello(context.Background(), &testPB.HelloReq{
		Name: "abc",
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v\n", rsp)
}

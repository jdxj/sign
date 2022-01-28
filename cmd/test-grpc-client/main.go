package main

import (
	"context"
	"fmt"
	"log"

	"go-micro.dev/v4"

	"github.com/jdxj/sign/internal/pkg/config"
	"github.com/jdxj/sign/internal/pkg/rpc"
	test_grpc "github.com/jdxj/sign/internal/proto/test-grpc"

	"github.com/asim/go-micro/plugins/registry/etcd/v4"
	"go-micro.dev/v4/registry"
)

func main() {
	root := config.ReadConfigs("/Users/ing/workspace/sign/configs/conf.yaml")
	service := micro.NewService(
		micro.Registry(
			etcd.NewRegistry(
				registry.Addrs(root.Etcd.Endpoints...),
				registry.TLSConfig(rpc.NewTLSConfig(root.Etcd.Ca, root.Etcd.Cert, root.Etcd.Key)),
			),
		),
	)
	service.Init()
	client := test_grpc.NewTestRPCService("test-grpc", service.Client())
	rsp, err := client.Hello(context.Background(), &test_grpc.HelloReq{
		Name: "abc",
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v\n", rsp)
}

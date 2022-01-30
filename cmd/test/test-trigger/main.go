package main

import (
	"context"
	"fmt"
	"log"

	"github.com/asim/go-micro/plugins/registry/etcd/v4"
	"github.com/urfave/cli/v2"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"

	"github.com/jdxj/sign/configs"
	"github.com/jdxj/sign/internal/pkg/util"
	"github.com/jdxj/sign/internal/proto/trigger"
)

func main() {
	service := micro.NewService(
		micro.Registry(etcd.NewRegistry()))
	service.Init(
		micro.Action(func(cli *cli.Context) error {
			return service.Options().
				Registry.Init(
				registry.Addrs(configs.EtcdEndpoint),
				registry.TLSConfig(
					util.NewTLSConfig(
						configs.Ca,
						configs.Cert,
						configs.Key)))
		}))
	triggerService := trigger.NewTriggerService("trigger", service.Client())

	rsp, err := triggerService.CreateTrigger(context.Background(), &trigger.CreateTriggerRequest{
		Trigger: &trigger.Trigger{
			Spec: "0 8 * * *",
		},
	})
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	fmt.Printf("%s\n", rsp)

	getRsp, err := triggerService.GetTriggers(context.Background(), &trigger.GetTriggersRequest{
		Offset: 0,
		Limit:  10,
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("getRsp: %+v\n", getRsp)
}

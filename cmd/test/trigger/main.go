package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/asim/go-micro/plugins/registry/etcd/v4"
	"github.com/urfave/cli/v2"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"

	"github.com/jdxj/sign/internal/pkg/util"
	pb "github.com/jdxj/sign/internal/proto/trigger"
)

func main() {
	service := micro.NewService(
		micro.Name("test-trigger"),
		micro.Registry(etcd.NewRegistry()))
	service.Init(
		micro.Action(func(cli *cli.Context) error {
			return service.Options().
				Registry.Init(
				registry.Addrs(""),
				registry.TLSConfig(
					util.NewTLSConfig("", "", "")))
		}))

	triggerService := pb.NewTriggerService(pb.ServiceName, service.Client())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	testCreateTrigger := func() {
		ctRsp, err := triggerService.CreateTrigger(ctx, &pb.CreateTriggerRequest{
			Trigger: &pb.Trigger{
				Spec: "0 8 * * *",
			},
		})
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Printf("ctRsp: %+v\n", ctRsp)
	}
	testCreateTrigger()

	testGetTriggers := func() {
		gtRsp, err := triggerService.GetTriggers(ctx, &pb.GetTriggersRequest{
			Offset: 0,
			Limit:  1000,
		})
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Printf("gtRsp: %+v\n", gtRsp)
	}
	testGetTriggers()
}

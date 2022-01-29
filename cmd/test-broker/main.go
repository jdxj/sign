package main

import (
	"fmt"
	"log"
	"time"

	"github.com/asim/go-micro/plugins/broker/rabbitmq/v4"
	"go-micro.dev/v4/broker"
	"google.golang.org/protobuf/proto"

	"github.com/jdxj/sign/configs"
	testPB "github.com/jdxj/sign/internal/proto/test-grpc"
)

func main() {
	b := rabbitmq.NewBroker()
	err := b.Init(
		broker.Addrs(configs.RabbitMQEndpoint),
	)
	if err != nil {
		log.Printf("%s\n", err)
	}

	err = b.Connect()
	if err != nil {
		log.Printf("%s\n", err)
	}

	for i := 0; i < 1000; i++ {
		time.Sleep(time.Second)

		req := &testPB.WorldReq{
			Name: fmt.Sprintf("%d", i),
		}
		d, err := proto.Marshal(req)
		if err != nil {
			log.Printf("Marshal: %s\n", err)
			continue
		}

		m := &broker.Message{
			Body: d,
		}
		err = b.Publish("task-dispatch", m)
		if err != nil {
			log.Printf("Publish: %s\n", err)
		}
	}
}

package main

import (
	"log"
	"time"

	"github.com/asim/go-micro/plugins/broker/rabbitmq/v4"
	"go-micro.dev/v4/broker"
	"google.golang.org/protobuf/proto"

	"github.com/jdxj/sign/configs"
	testPB "github.com/jdxj/sign/internal/proto/test-grpc"
)

func main() {
	go testSub("1")
	go testSub("2")
	time.Sleep(time.Hour)
}

func testSub(id string) {
	b := rabbitmq.NewBroker()
	err := b.Init(
		broker.Addrs(configs.RabbitMQEndpoint),
	)
	if err != nil {
		log.Printf("%s\n", err)
		return
	}

	err = b.Connect()
	if err != nil {
		log.Printf("%s\n", err)
		return
	}

	_, err = b.Subscribe("task-dispatch", eventHandlerD(id), broker.Queue("task-consumer"))
	if err != nil {
		log.Printf("Subscribe: %s\n", err)
	}
	time.Sleep(time.Hour)
}

func eventHandlerD(id string) broker.Handler {
	return func(e broker.Event) error {
		log.Printf("id: %s", id)
		return eventHandler(e)
	}
}

func eventHandler(e broker.Event) error {
	req := &testPB.WorldReq{}
	err := proto.Unmarshal(e.Message().Body, req)
	if err != nil {
		log.Printf("Unmarshal: %s\n", err)
		return err
	}
	//fmt.Printf("name: %s\n", req.Name)
	log.Printf("name: %s\n", req.Name)
	return nil
}

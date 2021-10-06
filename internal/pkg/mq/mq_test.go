package mq

import (
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"google.golang.org/protobuf/proto"

	"github.com/jdxj/sign/internal/pkg/config"
	"github.com/jdxj/sign/internal/pkg/logger"
	"github.com/jdxj/sign/internal/proto/crontab"
)

func TestMain(t *testing.M) {
	logger.Init("./mq.log")
	conf := config.Rabbit{
		Host: "127.0.0.1",
		Port: 5672,
		User: "guest",
		Pass: "guest",
	}
	InitRabbit(conf)
	os.Exit(t.Run())
}

func TestNewTaskQueue(t *testing.T) {
	tq1, err := NewTaskQueue()
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	tq2, err := NewTaskQueue()
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	tq3, err := NewTaskQueue()
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	go func() {
		var i int
		for {
			i++
			err := tq1.Publish([]byte(strconv.Itoa(i)))
			if err != nil {
				fmt.Printf("tq1 publist: %s\n", err)
			} else {
				fmt.Printf("push ok: %d\n", i)
			}

			time.Sleep(time.Second)
		}
	}()

	go func() {
		msgChan, err := tq2.Consume()
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		for msg := range msgChan {
			fmt.Printf("tq2 receive: %s\n", msg)
		}
	}()

	go func() {
		msgChan, err := tq3.Consume()
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		for msg := range msgChan {
			fmt.Printf("tq3 receive: %s\n", msg)
		}
		fmt.Printf("bad\n")
	}()

	time.Sleep(10 * time.Second)
	tq1.Stop()
	tq2.Stop()
	tq3.Stop()
}

func TestTaskQueue_Consume(t *testing.T) {
	tq, err := NewTaskQueue()
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	defer tq.Stop()

	dataChan, err := tq.Consume()
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	for v := range dataChan {
		task := &crontab.Task{}
		err := proto.Unmarshal(v, task)
		if err != nil {
			t.Fatalf("%s\n", err)
		} else {
			fmt.Printf("%+v\n", task)
		}
	}
}

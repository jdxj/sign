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

	"github.com/jdxj/sign/internal/pkg/config"
	"github.com/jdxj/sign/internal/pkg/util"
	pb "github.com/jdxj/sign/internal/proto/task"
)

func main() {
	service := micro.NewService(
		micro.Name("test-task"),
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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	taskService := pb.NewTaskService(pb.ServiceName, service.Client())
	testCreateTask := func() {
		ctRsp, err := taskService.CreateTask(ctx, &pb.CreateTaskRequest{Task: &pb.Task{
			Description: "test create task",
			UserId:      1,
			Kind:        pb.Kind_MOCK.String(),
			Spec:        "* * * * *",
			Param:       []byte("abc"),
		}})
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Printf("ctRsp: %+v\n", ctRsp)
	}
	testCreateTask()

	testGetTask := func() {
		gtRsp, err := taskService.GetTask(ctx, &pb.GetTaskRequest{TaskId: 5})
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Printf("gtRsp: %+v\n", gtRsp)
	}
	testGetTask()

	testGetTasks := func() {
		gtsRsp, err := taskService.GetTasks(ctx, &pb.GetTasksRequest{
			TaskId:      0,
			Description: "test",
			UserId:      0,
			Kind:        "",
			Spec:        "",
			CreatedAt:   nil,
			Offset:      0,
			Limit:       1000,
		})
		if err != nil {
			log.Println(err)
			return
		}
		for _, v := range gtsRsp.GetTasks() {
			fmt.Printf("%+v\n", v)
		}
	}
	testGetTasks()

	testUpdateTask := func() {
		_, err := taskService.UpdateTask(ctx, &pb.UpdateTaskRequest{Task: &pb.Task{
			TaskId:      1,
			Description: "def",
			UserId:      0,
			Kind:        "",
			Spec:        "",
			Param:       nil,
			CreatedAt:   nil,
		}})
		if err != nil {
			log.Println(err)
			return
		}
	}
	testUpdateTask()

	testDispatchTasks := func() {
		spec := "* * * * *"
		_, err := taskService.DispatchTasks(ctx, &pb.DispatchTasksRequest{Spec: spec})
		if err != nil {
			log.Println(err)
			return
		}
	}
	testDispatchTasks()
}

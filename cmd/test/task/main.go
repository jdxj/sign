package main

import (
	"context"
	"errors"
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

var (
	TaskService       pb.TaskService
	ErrConfigNotFound = errors.New("config not found")
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
				return ErrConfigNotFound
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

	TaskService = pb.NewTaskService(pb.ServiceName, service.Client())

	//testGetTask := func() {
	//	gtRsp, err := taskService.GetTask(ctx, &pb.GetTaskRequest{TaskId: 5})
	//	if err != nil {
	//		log.Println(err)
	//		return
	//	}
	//	fmt.Printf("gtRsp: %+v\n", gtRsp)
	//}
	//testGetTask()

	//testGetTasks := func() {
	//	gtsRsp, err := taskService.GetTasks(ctx, &pb.GetTasksRequest{
	//		TaskId:      0,
	//		Description: "test",
	//		UserId:      0,
	//		Kind:        "",
	//		Spec:        "",
	//		CreatedAt:   nil,
	//		Offset:      0,
	//		Limit:       1000,
	//	})
	//	if err != nil {
	//		log.Println(err)
	//		return
	//	}
	//	for _, v := range gtsRsp.GetTasks() {
	//		fmt.Printf("%+v\n", v)
	//	}
	//}
	//testGetTasks()

	//testUpdateTask := func() {
	//	_, err := taskService.UpdateTask(ctx, &pb.UpdateTaskRequest{Task: &pb.Task{
	//		TaskId:      2,
	//		Description: "test update bili task",
	//		UserId:      1,
	//		Kind:        "",
	//		Spec:        "",
	//		Param:       configs.GetBiLiParam(),
	//		CreatedAt:   nil,
	//	}})
	//	if err != nil {
	//		log.Println(err)
	//		return
	//	}
	//}
	//testUpdateTask()

	//testDispatchTasks := func() {
	//	spec := "* * * * *"
	//	_, err := taskService.DispatchTasks(ctx, &pb.DispatchTasksRequest{Spec: spec})
	//	if err != nil {
	//		log.Println(err)
	//		return
	//	}
	//}
	//testDispatchTasks()

	testCreateTask(ctx)
	//testDeleteTask(ctx)
}

func testCreateTask(ctx context.Context) {
	ctRsp, err := TaskService.CreateTask(ctx, &pb.CreateTaskRequest{Task: &pb.Task{
		Description: "test create custom message task",
		UserId:      1,
		Kind:        pb.Kind_CUSTOM_MESSAGE.String(),
		Spec:        "* * * * *",
		Param:       nil,
	}})
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("ctRsp: %+v\n", ctRsp)
}

//func testDeleteTask(ctx context.Context) {
//	_, err := TaskService.DeleteTask(ctx, &pb.DeleteTaskRequest{
//		TaskId: 2,
//		UserId: 1,
//	})
//	if err != nil {
//		log.Println(err)
//	}
//}

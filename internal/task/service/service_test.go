package service

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/jdxj/sign/internal/pkg/config"
	"github.com/jdxj/sign/internal/pkg/db"
	pb "github.com/jdxj/sign/internal/proto/task"
)

func TestMain(t *testing.M) {
	_ = db.InitGorm(config.DB{
		User:   "root",
		Pass:   "123456",
		Host:   "127.0.0.1",
		Port:   3306,
		Dbname: "sign",
	})
	db.Debug()
	os.Exit(t.Run())
}

func TestService_GetTasks(t *testing.T) {
	req := &pb.GetTasksRequest{
		UserId: 2,
		Offset: 0,
		Limit:  1000,
	}
	rsp := &pb.GetTasksResponse{}

	s := New(config.Secret{})
	err := s.GetTasks(context.Background(), req, rsp)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	for _, v := range rsp.GetTasks() {
		fmt.Printf("%+v\n", v)
	}
}

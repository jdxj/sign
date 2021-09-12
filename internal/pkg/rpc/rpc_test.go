package rpc

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	clientV3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"

	"github.com/jdxj/sign/internal/crontab/service"
	"github.com/jdxj/sign/internal/pkg/logger"
	"github.com/jdxj/sign/internal/proto/crontab"
)

func TestMain(t *testing.M) {
	logger.Init("./rpc.log")
	Init("127.0.0.1:2379")
	os.Exit(t.Run())
}

func TestServer_Register(t *testing.T) {
	server, err := NewServer("hello",
		"127.0.0.1:2379", 49152)
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	server.Register(func(registrar grpc.ServiceRegistrar) {
		crontab.RegisterTestServiceServer(registrar, &service.Service{})
	})
	server.Serve()
}

func TestEtcdGet(t *testing.T) {
	client, err := clientV3.New(clientV3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
	})
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	rsp, err := client.Get(context.Background(), "a", clientV3.WithPrefix())
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	for _, kv := range rsp.Kvs {
		fmt.Printf("%+v\n", kv)
	}
}

func TestSignScheme(t *testing.T) {
	cc, err := grpc.Dial(DSN("crontab"), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	client := crontab.NewTestServiceClient(cc)
	rsp, err := client.Hello(context.Background(), &crontab.TestReq{
		Nickname: "jdxj-man",
	})
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("rsp: %+v\n", rsp)
	time.Sleep(time.Hour)
}

func TestNewClient(t *testing.T) {
	var client crontab.TestServiceClient
	NewClient("crontab", func(cc *grpc.ClientConn) {
		client = crontab.NewTestServiceClient(cc)
	})
	rsp, err := client.Hello(context.Background(), &crontab.TestReq{
		Nickname: "jdxj-man",
	})
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("rsp: %+v\n", rsp)
	time.Sleep(time.Hour)
}

func TestGetListenAddr(t *testing.T) {
	addr, err := GetListenAddr(49152, "release")
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("addr: %s\n", addr)
}

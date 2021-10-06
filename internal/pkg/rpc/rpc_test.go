package rpc

import (
	"context"
	"fmt"
	"os"
	"testing"

	clientV3 "go.etcd.io/etcd/client/v3"

	"github.com/jdxj/sign/internal/pkg/logger"
)

func TestMain(t *testing.M) {
	logger.Init("./rpc.log")
	Init("127.0.0.1:2379")
	os.Exit(t.Run())
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

func TestGetListenAddr(t *testing.T) {
	addr, err := GetListenAddr(49152, "release")
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("addr: %s\n", addr)
}

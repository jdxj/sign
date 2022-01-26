package main

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc/resolver"

	"github.com/jdxj/sign/internal/pkg/rpc"
	testPB "github.com/jdxj/sign/internal/proto/test-grpc"
)

func main() {
	resolver.Register(&rpc.LocalBuilder{})

	cc := testPB.NewTestRPCClient(rpc.NewConn(testPB.ServiceName))

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	rsp, err := cc.Hello(ctx, &testPB.HelloReq{
		Name: "abc",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("rsp: %+v\n", rsp)

	cc2 := testPB.NewTestMultiRPCClient(rpc.NewConn(testPB.MServiceName))

	ctx2, cancel2 := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel2()

	rsp2, err2 := cc2.World(ctx2, &testPB.WorldReq{
		Name: "abc",
	})
	if err2 != nil {
		panic(fmt.Errorf("z: %s", err2))
	}
	fmt.Printf("rsp2: %+v\n", rsp2)

	for i := 0; i < 10; i++ {
		cc2.World(ctx2, &testPB.WorldReq{
			Name: "abc",
		})
		time.Sleep(time.Second)
	}
}

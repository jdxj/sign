package bwh

import (
	"fmt"
	"testing"
	"time"

	"google.golang.org/protobuf/proto"

	pb "github.com/jdxj/sign/internal/proto/task"
)

func TestNetworkMeter_Execute(t *testing.T) {
	nm := &NetworkMeter{}
	param := &pb.BWH{
		VeId:   "",
		ApiKey: "",
	}
	body, _ := proto.Marshal(param)
	res, err := nm.Execute(body)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("res-%s\n", res)
}

func TestDate(t *testing.T) {
	tt := time.Date(2022, 2, 22, 13, 0, 0, 0, time.Local)
	t2 := tt.AddDate(0, 0, -30)
	fmt.Printf("t: %s\n", t2)
}

func TestCPUMeter_Execute(t *testing.T) {
	cm := &CPUMeter{}
	param := &pb.BWH{
		VeId:   "",
		ApiKey: "",
	}
	body, _ := proto.Marshal(param)
	_, err := cm.Execute(body)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
}

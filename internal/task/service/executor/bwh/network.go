package bwh

import (
	"fmt"

	kwv "github.com/jdxj/kiwivm-sdk-go"
	"google.golang.org/protobuf/proto"

	pb "github.com/jdxj/sign/internal/proto/task"
)

const (
	msgParseParamFailed = "解析参数失败"
	msgGetNetworkFailed = "获取流量使用数据失败"
	msgGetCPUFailed     = "获取CPU使用数据失败"
	msgCPULimit         = "CPU使用可能超过100%"
)

type NetworkMeter struct {
}

func (nm *NetworkMeter) Kind() string {
	return pb.Kind_BWH_NETWORK.String()
}

func (nm *NetworkMeter) Execute(body []byte) (string, error) {
	param := &pb.BWH{}
	err := proto.Unmarshal(body, param)
	if err != nil {
		return msgParseParamFailed, err
	}

	client := kwv.NewClient(param.GetVeId(), param.GetApiKey())
	rsp, err := client.GetRawUsageStats()
	if err != nil {
		return msgGetNetworkFailed, err
	}

	var count int64
	// 一天有288个5分钟
	for i := len(rsp.Data) - 1; i >= 0 && i >= len(rsp.Data)-288; i-- {
		count += rsp.Data[i].NetworkInBytes + rsp.Data[i].NetworkOutBytes
	}
	return fmt.Sprintf("BWH流量使用: %.3fMB",
		float64(count)/float64(1000*1000)), nil
}

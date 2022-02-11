package bwh

import (
	kwv "github.com/jdxj/kiwivm-sdk-go"
	"google.golang.org/protobuf/proto"

	pb "github.com/jdxj/sign/internal/proto/task"
)

type CPUMeter struct {
}

func (cm *CPUMeter) Kind() string {
	return pb.Kind_BWH_CPU.String()
}

func (cm *CPUMeter) Execute(body []byte) (string, error) {
	param := &pb.BWH{}
	err := proto.Unmarshal(body, param)
	if err != nil {
		return msgParseParamFailed, err
	}
	client := kwv.NewClient(param.GetVeId(), param.GetApiKey())
	rsp, err := client.GetRawUsageStats()
	if err != nil {
		return msgGetCPUFailed, err
	}

	// 一小时有12个5分钟
	var count int64
	for i := len(rsp.Data) - 1; i >= 0 && i >= len(rsp.Data)-12; i-- {
		count += rsp.Data[i].CpuUsage
	}
	if count < 100 {
		// 无需通知
		return "", nil
	}
	return msgCPULimit, nil
}

package bwh

import (
	"fmt"
	"time"

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

const (
	// Byte to Gigabyte
	gb float64 = 1 << 30
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

	var (
		now   = time.Now()
		day   = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
		month = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local)

		dayCount   int64
		monthCount int64
	)
	for i := len(rsp.Data) - 1; i >= 0; i-- {
		v := rsp.Data[i]
		t := time.Unix(v.Timestamp, 0)

		if t.After(day) {
			dayCount += v.NetworkInBytes + v.NetworkOutBytes
		}
		if t.After(month) {
			monthCount += v.NetworkInBytes + v.NetworkOutBytes
		} else {
			// rsp.Data 是按照时间排序的,
			// 一个月的数据查完后直接退出.
			break
		}
	}
	return fmt.Sprintf(`[BWH流量使用]
当天用量: %.3fGB
当月用量: %.3fGB`,
		float64(dayCount)/gb,
		float64(monthCount)/gb), nil
}

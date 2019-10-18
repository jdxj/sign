package pic

import (
	"encoding/json"
)

// 访问签到链接前需要获取一些数据
// json 格式

type Conf struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

type Data struct {
	Status int `json:"status"`
	Msg    Msg `json:"msg"`
}

type Msg struct {
	CycleTime string `json:"cycle_time"`
}

func CycleTime(data []byte) string {
	conf := Conf{}
	err := json.Unmarshal(data, &conf)
	if err != nil {
		return ""
	}

	return conf.Data.Msg.CycleTime
}

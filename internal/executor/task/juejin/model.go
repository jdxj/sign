package juejin

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/jdxj/sign/internal/pkg/util"
)

const (
	domain = ".juejin.cn"
	home   = "https://juejin.cn"

	apiPrefix   = "https://api.juejin.cn"
	signInURL   = apiPrefix + "/growth_api/v1/check_in"
	countsURL   = apiPrefix + "/growth_api/v1/get_counts"
	pointURL    = apiPrefix + "/growth_api/v1/get_cur_point"
	calendarURL = apiPrefix + "/growth_api/v1/get_coder_calendar"
)

const (
	msgJueJinExecFailed = "掘金任务执行失败"
)

var (
	ErrUnknownMistake = errors.New("unknown mistake")
)

type response struct {
	ErrNo  int         `json:"err_no"`
	ErrMsg string      `json:"err_msg"`
	Data   interface{} `json:"data"`
}

func execute(key, url string, data fmt.Stringer) (string, error) {
	jar := util.NewJar(key, domain, home)
	client := &http.Client{Jar: jar}

	rsp := &response{
		Data: data,
	}
	err := util.ParseBody(client, url, rsp)
	if err != nil {
		return msgJueJinExecFailed, fmt.Errorf("%w, stage: %s", err, crontab.Stage_Auth)
	}
	if rsp.ErrNo != 0 {
		return msgJueJinExecFailed, fmt.Errorf("%w: %s, stage: %s",
			ErrUnknownMistake, rsp.ErrMsg, crontab.Stage_Auth)
	}
	return data.String(), nil
}

type checkIn struct {
	IncrPoint int `json:"incr_point"`
	SumPoint  int `json:"sum_point"`
}

func (ci *checkIn) String() string {
	format := `掘金签到成功
增加点数: %d
总计点数: %d`
	return fmt.Sprintf(format, ci.IncrPoint, ci.SumPoint)
}

type counts struct {
	ContCount int `json:"cont_count"`
	SumCount  int `json:"sum_count"`
}

func (c *counts) String() string {
	format := `掘金签到天数统计
连续签到天数: %d
累计签到天数: %d`
	return fmt.Sprintf(format, c.ContCount, c.SumCount)
}

type ore int

func (o *ore) String() string {
	format := `掘金矿石数统计
当前矿石数: %d`
	return fmt.Sprintf(format, *o)
}

type jokes struct {
	Aphorism    string `json:"aphorism"`
	ShouldOrNot string `json:"should_or_not"`
}

func (j *jokes) String() string {
	format := `码农日历
格言: %s
宜忌: %s`
	return fmt.Sprintf(format, j.Aphorism, j.ShouldOrNot)
}

package bili

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jdxj/sign/internal/bot"
	"github.com/jdxj/sign/pkg/task/common"
)

const (
	Domain    = ".bilibili.com"
	SignURL   = "https://www.bilibili.com/"
	AuthURL   = "https://api.bilibili.com/x/member/web/account"
	VerifyURL = "https://api.bilibili.com/x/member/web/coin/log?jsonp=jsonp"

	BiURL = "https://account.bilibili.com/site/getCoin"
)

type AuthResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		MID    int    `json:"mid"`
		Uname  string `json:"uname"`
		UserID string `json:"user_id"`
	} `json:"data"`
}

type VerifyResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		List []struct {
			Time   string `json:"time"`
			Delta  int    `json:"delta"`
			Reason string `json:"reason"`
		} `json:"list"`
		Count int `json:"count"`
	} `json:"data"`
}

func Auth(cookies string) (*http.Client, error) {
	jar := common.NewJar(cookies, Domain, SignURL)
	client := &http.Client{Jar: jar}

	authResp := &AuthResp{}
	err := common.ParseBody(client, AuthURL, authResp)
	if err != nil {
		return client, err
	}

	if authResp.Code != 0 {
		return client, common.ErrorAuthFailed
	}
	return client, nil
}

func SignIn(task *common.Task) bool {
	err := common.ParseBody(task.Client, SignURL, nil)
	if err != nil {
		text := fmt.Sprintf("访问失败, id: %s, type: %s, err: %s",
			task.ID, common.TypeMap[task.Type], err)
		bot.Send(text)
		return false
	}

	err = verify(task.Client)
	if err != nil {
		text := fmt.Sprintf("验证失败, id: %s, type: %s, err: %s",
			task.ID, common.TypeMap[task.Type], err)
		bot.Send(text)
		return false
	}

	text := fmt.Sprintf("签到成功: id: %s, type: %s",
		task.ID, common.TypeMap[task.Type])
	bot.Send(text)
	return true
}

func verify(client *http.Client) error {
	verifyResp := &VerifyResp{}
	err := common.ParseBody(client, VerifyURL, verifyResp)
	if err != nil {
		return err
	}

	if verifyResp.Code != 0 {
		return common.ErrorAuthFailed
	}
	list := verifyResp.Data.List
	if len(list) <= 0 {
		return common.ErrorLogNotFound
	}

	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return err
	}
	now := time.Now().In(loc)
	last, err := time.ParseInLocation("2006-01-02 15:04:05", list[0].Time, loc)
	if err != nil {
		return err
	}

	if now.YearDay() != last.YearDay() {
		return common.ErrorSignInFailed
	}
	return nil
}

type BiResp struct {
	Code   int  `json:"code"`
	Status bool `json:"status"`
	Data   struct {
		Money int `json:"money"`
	} `json:"data"`
}

func QueryBi(task *common.Task) bool {
	biResp := &BiResp{}
	err := common.ParseBody(task.Client, BiURL, biResp)
	if err != nil {
		text := fmt.Sprintf("查询失败, id: %s, type: %s, err: %s",
			task.ID, common.TypeMap[task.Type], err)
		bot.Send(text)
		return false
	}

	if biResp.Code != 0 {
		text := fmt.Sprintf("查询失败, id: %s, type: %s, err: %s",
			task.ID, common.TypeMap[task.Type], common.ErrorAuthFailed)
		bot.Send(text)
		return false
	}

	text := fmt.Sprintf("查询成功, id: %s, type: %s, money: %d",
		task.ID, common.TypeMap[task.Type], biResp.Data.Money)
	bot.Send(text)
	return true
}

package bili

import (
	"net/http"
	"time"

	"github.com/jdxj/sign/pkg/task/common"
)

const (
	Domain    = ".bilibili.com"
	SignURL   = "https://www.bilibili.com/"
	AuthURL   = "https://api.bilibili.com/x/member/web/account"
	VerifyURL = "https://api.bilibili.com/x/member/web/coin/log?jsonp=jsonp"
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

func SignIn(client *http.Client) error {
	err := common.ParseBody(client, SignURL, nil)
	if err != nil {
		return err
	}

	return verify(client)
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

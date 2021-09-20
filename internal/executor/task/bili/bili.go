package bili

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jdxj/sign/internal/task/common"
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

func SignIn(c *http.Client) (err error) {
	for i := 0; i < common.RetryNumber; i++ {
		if err = signIn(c); err == nil {
			return
		}
		time.Sleep(common.RetryInterval)
	}
	return
}

func signIn(c *http.Client) error {
	err := common.ParseBody(c, SignURL, nil)
	if err != nil {
		return fmt.Errorf("stage: %s, err: %w", common.SignIn, err)
	}

	err = verify(c)
	if err != nil {
		return fmt.Errorf("stage: %s, err: %w", common.Verify, err)
	}
	return nil
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

func QueryBi(c *http.Client) (msg string, err error) {
	for i := 0; i < common.RetryNumber; i++ {
		if msg, err = queryBi(c); err == nil {
			return
		}
		time.Sleep(common.RetryInterval)
	}
	return
}

func queryBi(c *http.Client) (string, error) {
	biResp := &BiResp{}
	err := common.ParseBody(c, BiURL, biResp)
	if err != nil {
		return "", fmt.Errorf("stage: %s, err: %w", common.Query, err)
	}

	if biResp.Code != 0 {
		return "", fmt.Errorf("stage: %s, err: %w", common.Verify, err)
	}

	msg := fmt.Sprintf("硬币: %d", biResp.Data.Money)
	return msg, nil
}

package stg

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/jdxj/sign/internal/task/common"
)

const (
	Domain    = ".studygolang.com"
	HomeURL   = "https://studygolang.com/"
	AuthURL   = "https://studygolang.com/balance"
	SignURL   = "https://studygolang.com/mission/daily/redeem"
	VerifyURL = "https://studygolang.com/balance"
)

var (
	regAuth   *regexp.Regexp
	regVerify *regexp.Regexp
)

func init() {
	regAuth = regexp.MustCompile(`每日登录奖励`)
	regVerify = regexp.MustCompile(`202\d-\d{2}-\d{2}`)
}

func Auth(cookies string) (*http.Client, error) {
	jar := common.NewJar(cookies, Domain, HomeURL)
	client := &http.Client{Jar: jar}

	body, err := common.ParseRawBody(client, AuthURL)
	if err != nil {
		return client, err
	}
	target := regAuth.FindString(string(body))
	if target == "" {
		return client, common.ErrorAuthFailed
	}
	return client, nil
}

func SignIn(c *http.Client) error {
	err := accessSignURL(c)
	if err != nil {
		return fmt.Errorf("stage: %s, err: %w", common.SignIn, err)
	}

	err = verify(c)
	if err != nil {
		return fmt.Errorf("stage: %s, err: %w", common.Verify, err)
	}
	return nil
}

func accessSignURL(c *http.Client) error {
	return common.ParseBody(c, SignURL, nil)
}

func verify(c *http.Client) error {
	body, err := common.ParseRawBody(c, VerifyURL)
	if err != nil {
		return err
	}
	date := regVerify.FindString(string(body))
	return common.VerifyDate(date)
}

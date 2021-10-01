package stg

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"

	"github.com/jdxj/sign/internal/executor/task"
	"github.com/jdxj/sign/internal/proto/crontab"
)

const (
	domain    = ".studygolang.com"
	home      = "https://studygolang.com/"
	authURL   = "https://studygolang.com/balance"
	signURL   = "https://studygolang.com/mission/daily/redeem"
	verifyURL = "https://studygolang.com/balance"
)

var (
	regAuth   *regexp.Regexp
	regVerify *regexp.Regexp
)

var (
	ErrTargetNotFound = errors.New("target not found")
)

func init() {
	regAuth = regexp.MustCompile(`每日登录奖励`)
	regVerify = regexp.MustCompile(`202\d-\d{2}-\d{2}`)
}

type SignIn struct {
}

func (si *SignIn) Domain() crontab.Domain {
	return crontab.Domain_STG
}

func (si *SignIn) Kind() crontab.Kind {
	return crontab.Kind_STGSign
}

func (si *SignIn) Execute(key string) (string, error) {
	c, err := auth(key)
	if err != nil {
		return "", err
	}

	err = signIn(c)
	if err != nil {
		return "", err
	}

	err = verify(c)
	if err != nil {
		return "", err
	}
	return "Go语言中文网签到成功", nil
}

func auth(cookies string) (*http.Client, error) {
	jar := task.NewJar(cookies, domain, home)
	client := &http.Client{Jar: jar}

	body, err := task.ParseRawBody(client, authURL)
	if err != nil {
		return client, err
	}
	target := regAuth.FindString(string(body))
	if target == "" {
		return client, fmt.Errorf("%w, stage: %s",
			ErrTargetNotFound, crontab.Stage_Auth)
	}
	return client, nil
}

func signIn(c *http.Client) error {
	err := task.ParseBody(c, signURL, nil)
	if err != nil {
		return fmt.Errorf("%w, stage: %s",
			err, crontab.Stage_SignIn)
	}
	return nil
}

func verify(c *http.Client) error {
	body, err := task.ParseRawBody(c, verifyURL)
	if err != nil {
		return fmt.Errorf("%w, stage: %s",
			err, crontab.Stage_Verify)
	}
	date := regVerify.FindString(string(body))
	err = task.VerifyDate(date)
	if err != nil {
		return fmt.Errorf("%w, stage: %s",
			err, crontab.Stage_Verify)
	}
	return nil
}

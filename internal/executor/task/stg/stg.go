package stg

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"regexp"

	"github.com/jdxj/sign/internal/pkg/util"
)

const (
	authURL   = "https://studygolang.com/balance"
	signURL   = "https://studygolang.com/mission/daily/redeem"
	verifyURL = "https://studygolang.com/balance"
	loginURL  = "https://studygolang.com/account/login"
)

const (
	msgSTGSignInFailed = "Go语言中文网签到失败"
)

var (
	regAuth   *regexp.Regexp
	regVerify *regexp.Regexp
)

var (
	ErrLoginFailed    = errors.New("login failed")
	ErrTargetNotFound = errors.New("target not found")
)

func init() {
	regAuth = regexp.MustCompile(`每日登录奖励`)
	regVerify = regexp.MustCompile(`202\d-\d{2}-\d{2}`)
}

type SignIn struct{}

func (si *SignIn) Domain() crontab.Domain {
	return crontab.Domain_STG
}

func (si *SignIn) Kind() crontab.Kind {
	return crontab.Kind_STGSign
}

func (si *SignIn) Execute(key string) (string, error) {
	c, err := auth(key)
	if err != nil {
		return msgSTGSignInFailed, err
	}

	err = signIn(c)
	if err != nil {
		return msgSTGSignInFailed, err
	}

	err = verify(c)
	if err != nil {
		return msgSTGSignInFailed, err
	}
	return "Go语言中文网签到成功", nil
}

func auth(key string) (*http.Client, error) {
	d := util.ConvertStringToMap(key)
	f := url.Values{}
	for k, v := range d {
		f.Set(k, v)
	}
	client, rsp, err := util.PostForm(loginURL, f)
	if err != nil {
		return nil, err
	}
	_ = rsp.Body.Close()
	if rsp.StatusCode != http.StatusOK {
		return nil, ErrLoginFailed
	}

	body, err := util.ParseRawBody(client, authURL)
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
	err := util.ParseBody(c, signURL, nil)
	if err != nil {
		return fmt.Errorf("%w, stage: %s",
			err, crontab.Stage_SignIn)
	}
	return nil
}

func verify(c *http.Client) error {
	body, err := util.ParseRawBody(c, verifyURL)
	if err != nil {
		return fmt.Errorf("%w, stage: %s",
			err, crontab.Stage_Verify)
	}
	date := regVerify.FindString(string(body))
	err = util.VerifyDate(date)
	if err != nil {
		return fmt.Errorf("%w, stage: %s",
			err, crontab.Stage_Verify)
	}
	return nil
}

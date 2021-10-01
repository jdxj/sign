package v2ex

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"

	"github.com/jdxj/sign/internal/executor/task"
	"github.com/jdxj/sign/internal/proto/crontab"
)

const (
	domain    = ".v2ex.com"
	home      = "https://www.v2ex.com/"
	authURL   = "https://www.v2ex.com/balance"
	tokenURL  = "https://www.v2ex.com/mission/daily"
	signURL   = "https://www.v2ex.com/mission/daily/redeem?once=%s"
	verifyURL = "https://www.v2ex.com/balance"
)

var (
	ErrTargetNotFound = errors.New("target not found")
	ErrTokenNotFound  = errors.New("token not found")
)

var (
	regAuth   *regexp.Regexp
	regVerify *regexp.Regexp
	regToken  *regexp.Regexp
)

func init() {
	regAuth = regexp.MustCompile(`每日登录奖励`)
	regVerify = regexp.MustCompile(`202\d-\d{2}-\d{2}`)
	regToken = regexp.MustCompile(`once=(.+)'`)
}

type SignIn struct {
}

func (si *SignIn) Domain() crontab.Domain {
	return crontab.Domain_V2EX
}

func (si *SignIn) Kind() crontab.Kind {
	return crontab.Kind_V2EXSign
}

func (si *SignIn) Execute(key string) (string, error) {
	c, err := auth(key)
	if err != nil {
		return "", err
	}

	token, err := getSignToken(c)
	if err != nil {
		return "", err
	}

	err = signIn(c, token)
	if err != nil {
		return "", err
	}

	err = verify(c)
	if err != nil {
		return "", err
	}
	return "V2ex签到成功", nil
}

func auth(cookies string) (*http.Client, error) {
	jar := task.NewJar(cookies, domain, home)
	client := &http.Client{Jar: jar}

	body, err := task.ParseRawBody(client, authURL)
	if err != nil {
		return client, fmt.Errorf("%w, stage: %s", err, crontab.Stage_Auth)
	}

	target := regAuth.FindString(string(body))
	if target == "" {
		err = fmt.Errorf("%w, stage: %s", ErrTargetNotFound, crontab.Stage_Auth)
	}
	return client, err
}

func getSignToken(c *http.Client) (string, error) {
	body, err := task.ParseRawBody(c, tokenURL)
	if err != nil {
		return "", fmt.Errorf("%w, stage: %s", err, crontab.Stage_Query)
	}

	matched := regToken.FindStringSubmatch(string(body))
	if len(matched) != 2 {
		return "", fmt.Errorf("%w, stage: %s", ErrTokenNotFound, crontab.Stage_Query)
	}
	return matched[1], nil
}

func signIn(c *http.Client, token string) error {
	u := fmt.Sprintf(signURL, token)
	err := task.ParseBody(c, u, nil)
	if err != nil {
		return fmt.Errorf("%w, stage: %s", err, crontab.Stage_SignIn)
	}
	return nil
}

func verify(c *http.Client) error {
	body, err := task.ParseRawBody(c, verifyURL)
	if err != nil {
		return fmt.Errorf("%w, stage: %s", err, crontab.Stage_Verify)
	}
	date := regVerify.FindString(string(body))
	err = task.VerifyDate(date)
	if err != nil {
		err = fmt.Errorf("%w, stage: %s", err, crontab.Stage_Verify)
	}
	return err
}

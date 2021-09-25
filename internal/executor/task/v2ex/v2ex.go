package v2ex

import (
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
		return client, fmt.Errorf("stage: %s, error: %w",
			crontab.Stage_Auth, err)
	}
	target := regAuth.FindString(string(body))
	if target == "" {
		return client, fmt.Errorf("stage: %s, error: %s",
			crontab.Stage_Auth, "target not found")
	}
	return client, nil
}

func getSignToken(c *http.Client) (string, error) {
	body, err := task.ParseRawBody(c, tokenURL)
	if err != nil {
		return "", fmt.Errorf("stage: %s, error: %w",
			crontab.Stage_Query, err)
	}

	matched := regToken.FindStringSubmatch(string(body))
	if len(matched) != 2 {
		return "", fmt.Errorf("stage: %s, error: %s",
			crontab.Stage_Query, "sign token not found")
	}
	return matched[1], nil
}

func signIn(c *http.Client, token string) error {
	u := fmt.Sprintf(signURL, token)
	err := task.ParseBody(c, u, nil)
	if err != nil {
		return fmt.Errorf("stage: %s, error: %w",
			crontab.Stage_SignIn, err)
	}
	return nil
}

func verify(c *http.Client) error {
	body, err := task.ParseRawBody(c, verifyURL)
	if err != nil {
		return fmt.Errorf("stage: %s, error: %w",
			crontab.Stage_Verify, err)
	}
	date := regVerify.FindString(string(body))
	err = task.VerifyDate(date)
	if err != nil {
		return fmt.Errorf("stage: %s, error: %w",
			crontab.Stage_Verify, err)
	}
	return nil
}

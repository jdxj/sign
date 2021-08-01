package v2ex

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/jdxj/sign/internal/task/common"
)

const (
	Domain    = ".v2ex.com"
	Home      = "https://www.v2ex.com/"
	AuthURL   = "https://www.v2ex.com/balance"
	TokenURL  = "https://www.v2ex.com/mission/daily"
	SignURL   = "https://www.v2ex.com/mission/daily/redeem?once=%s"
	VerifyURL = "https://www.v2ex.com/balance"
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

func Auth(cookies string) (*http.Client, error) {
	jar := common.NewJar(cookies, Domain, Home)
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
	st, err := getSignToken(c)
	if err != nil {
		return fmt.Errorf("stage: %s, err: %w", common.GetToken, err)
	}

	err = accessSignURL(c, st)
	if err != nil {
		return fmt.Errorf("stage: %s, err: %w", common.SignIn, err)
	}

	err = verify(c)
	if err != nil {
		return fmt.Errorf("stage: %s, err: %w", common.Verify, err)
	}
	return nil
}

func getSignToken(c *http.Client) (string, error) {
	body, err := common.ParseRawBody(c, TokenURL)
	if err != nil {
		return "", err
	}

	matched := regToken.FindStringSubmatch(string(body))
	if len(matched) != 2 {
		return "", fmt.Errorf("sign token not found: %v", matched)
	}
	return matched[1], nil
}

func accessSignURL(c *http.Client, token string) error {
	u := fmt.Sprintf(SignURL, token)
	return common.ParseBody(c, u, nil)
}

func verify(c *http.Client) error {
	body, err := common.ParseRawBody(c, VerifyURL)
	if err != nil {
		return err
	}
	date := regVerify.FindString(string(body))
	return common.VerifyDate(date)
}

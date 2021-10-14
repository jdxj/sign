package hpi

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/jdxj/sign/internal/executor/task"
	"github.com/jdxj/sign/internal/proto/crontab"
)

const (
	domain       = ".ld246.com"
	home         = "https://ld246.com/"
	authURL      = "https://ld246.com/notifications/unread/count?_=%d"
	signTokenURL = "https://ld246.com/activity/checkin"
	signURL      = "https://ld246.com/activity/daily-checkin?token=%s"
	verifyURL    = "https://ld246.com/member/%s/points?p=1&pjax=true"
)

const (
	msgHPISignInFailed = "黑客派签到失败"
)

var (
	regSignToken *regexp.Regexp
	regVerify    *regexp.Regexp
	regUserName  *regexp.Regexp
)

var (
	ErrGetToken         = errors.New("get token failed")
	ErrTokenNotFound    = errors.New("token not found")
	ErrUserNameNotFound = errors.New("user name not found")
)

func init() {
	regSignToken = regexp.MustCompile(`csrfToken: '(.+)'`)
	regVerify = regexp.MustCompile(`202\d-\d{2}-\d{2}`)
	regUserName = regexp.MustCompile(`currentUserName: '(.+)'`)
}

type SignIn struct{}

func (si *SignIn) Domain() crontab.Domain {
	return crontab.Domain_HPI
}

func (si *SignIn) Kind() crontab.Kind {
	return crontab.Kind_HPISign
}

func (si *SignIn) Execute(key string) (string, error) {
	c, err := auth(key)
	if err != nil {
		return msgHPISignInFailed, err
	}

	token, userName, err := getSignToken(c)
	if err != nil {
		return msgHPISignInFailed, err
	}

	err = signIn(c, token)
	if err != nil {
		return msgHPISignInFailed, err
	}

	err = verify(c, userName)
	if err != nil {
		return msgHPISignInFailed, err
	}
	return "黑客派签到成功", nil
}

func auth(cookies string) (*http.Client, error) {
	jar := task.NewJar(cookies, domain, home)
	client := &http.Client{Jar: jar}
	authResp := make(map[string]interface{})

	param := time.Now().UnixNano() / 1000000
	u := fmt.Sprintf(authURL, param)
	err := task.ParseBody(client, u, &authResp)
	if err != nil {
		return client, err
	}
	return client, nil
}

func getSignToken(client *http.Client) (string, string, error) {
	body, err := task.ParseRawBody(client, signTokenURL)
	if err != nil {
		return "", "", fmt.Errorf("%w: %s, stage: %s",
			ErrGetToken, err, crontab.Stage_Auth)
	}

	matched := regSignToken.FindStringSubmatch(string(body))
	if len(matched) != 2 {
		return "", "", fmt.Errorf("%w, stage: %s",
			ErrTokenNotFound, crontab.Stage_Auth)
	}
	token := matched[1]

	matched = regUserName.FindStringSubmatch(string(body))
	if len(matched) != 2 {
		return token, "", fmt.Errorf("%w, stage: %s",
			ErrUserNameNotFound, crontab.Stage_Auth)
	}
	userName := matched[1]
	return token, userName, nil
}

func signIn(client *http.Client, token string) error {
	u := fmt.Sprintf(signURL, token)
	header := map[string]string{
		"Referer": signTokenURL,
	}
	return task.ParseBodyHeader(client, u, header)
}

func verify(client *http.Client, id string) error {
	u := fmt.Sprintf(verifyURL, id)
	d, err := task.ParseRawBody(client, u)
	if err != nil {
		return err
	}
	date := regVerify.FindString(string(d))
	return task.VerifyDate(date)
}

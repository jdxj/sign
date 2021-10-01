package bili

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/jdxj/sign/internal/executor/task"
	"github.com/jdxj/sign/internal/proto/crontab"
)

var (
	ErrLogNotFound = errors.New("log not found")
	ErrSignIn      = errors.New("sign in failed")
)

type SignIn struct {
}

func (si *SignIn) Domain() crontab.Domain {
	return crontab.Domain_BILI
}

func (si *SignIn) Kind() crontab.Kind {
	return crontab.Kind_BILISignIn
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

	return "B站签到成功", nil
}

type authResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		MID    int    `json:"mid"`
		Uname  string `json:"uname"`
		UserID string `json:"user_id"`
	} `json:"data"`
}

func auth(cookies string) (*http.Client, error) {
	jar := task.NewJar(cookies, domain, signURL)
	client := &http.Client{Jar: jar}

	authResp := &authResp{}
	err := task.ParseBody(client, authURL, authResp)
	if err != nil {
		return client, fmt.Errorf("%w, stage: %s", err, crontab.Stage_Auth)
	}

	if authResp.Code != 0 {
		return client, fmt.Errorf("%w, stage: %s", ErrInvalidCookie, crontab.Stage_Auth)
	}
	return client, nil
}

func signIn(c *http.Client) error {
	err := task.ParseBody(c, signURL, nil)
	if err != nil {
		err = fmt.Errorf("%w, stage: %s", err, crontab.Stage_SignIn)
	}
	return err
}

type verifyResp struct {
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

func verify(client *http.Client) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("%w, stage: %s", err, crontab.Stage_Verify)
		}
	}()

	verifyResp := &verifyResp{}
	err = task.ParseBody(client, verifyURL, verifyResp)
	if err != nil {
		return
	}
	if verifyResp.Code != 0 {
		return ErrInvalidCookie
	}

	list := verifyResp.Data.List
	if len(list) <= 0 {
		return ErrLogNotFound
	}

	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return
	}
	now := time.Now().In(loc)
	last, err := time.ParseInLocation("2006-01-02 15:04:05", list[0].Time, loc)
	if err != nil {
		return
	}

	if now.YearDay() != last.YearDay() {
		err = ErrSignIn
	}
	return
}

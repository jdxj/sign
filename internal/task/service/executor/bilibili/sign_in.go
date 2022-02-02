package bilibili

import (
	"fmt"
	"net/http"
	"time"

	"google.golang.org/protobuf/proto"

	"github.com/jdxj/sign/internal/pkg/util"
	pb "github.com/jdxj/sign/internal/proto/task"
)

type SignIn struct{}

func (si *SignIn) Kind() string {
	return pb.Kind_BILIBILI_SIGN_IN.String()
}

func (si *SignIn) Execute(body []byte) (string, error) {
	param := &pb.BiLiBiLi{}
	err := proto.Unmarshal(body, param)
	if err != nil {
		return msgParseParam, err
	}

	c, err := auth(param.GetCookie())
	if err != nil {
		return msgSignInFailed, err
	}
	err = signIn(c)
	if err != nil {
		return msgSignInFailed, err
	}
	err = verify(c)
	if err != nil {
		return msgSignInFailed, err
	}
	return "B站签到成功", nil
}

func auth(cookies string) (*http.Client, error) {
	jar := util.NewJar(cookies, domain, signURL)
	client := &http.Client{Jar: jar}
	authResp := &authResp{}
	err := util.ParseBody(client, authURL, authResp)
	if err != nil {
		return client, fmt.Errorf("%w, stage: %s", err, stageAuth)
	}
	if authResp.Code != 0 {
		return client, fmt.Errorf("%w, stage: %s", ErrInvalidCookie, stageAuth)
	}
	return client, nil
}

func signIn(c *http.Client) error {
	err := util.ParseBody(c, signURL, nil)
	if err != nil {
		err = fmt.Errorf("%w, stage: %s", err, stageSignIn)
	}
	return err
}

func verify(client *http.Client) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("%w, stage: %s", err, stageVerify)
		}
	}()

	verifyResp := &verifyResp{}
	err = util.ParseBody(client, verifyURL, verifyResp)
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

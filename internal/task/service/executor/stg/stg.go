package stg

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"regexp"

	"google.golang.org/protobuf/proto"

	"github.com/jdxj/sign/internal/pkg/util"
	pb "github.com/jdxj/sign/internal/proto/task"
)

const (
	authURL   = "https://studygolang.com/balance"
	signURL   = "https://studygolang.com/mission/daily/redeem"
	verifyURL = "https://studygolang.com/balance"
	loginURL  = "https://studygolang.com/account/login"
)

const (
	msgParseParamFailed = "解析参数失败"
	msgSTGSignInFailed  = "Go语言中文网签到失败"
)

const (
	stageAuth   = "auth"
	stageSignIn = "sign-in"
	stageVerify = "verify"
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

func (si *SignIn) Kind() string {
	return pb.Kind_STG_SIGN_IN.String()
}

func (si *SignIn) Execute(body []byte) (string, error) {
	param := &pb.STG{}
	err := proto.Unmarshal(body, param)
	if err != nil {
		return msgParseParamFailed, err
	}

	c, err := auth(param.GetUsername(), param.GetPasswd())
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

func auth(username, passwd string) (*http.Client, error) {
	f := url.Values{}
	f.Set("username", username)
	f.Set("passwd", passwd)

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
			ErrTargetNotFound, stageAuth)
	}
	return client, nil
}

func signIn(c *http.Client) error {
	err := util.ParseBody(c, signURL, nil)
	if err != nil {
		return fmt.Errorf("%w, stage: %s",
			err, stageSignIn)
	}
	return nil
}

func verify(c *http.Client) error {
	body, err := util.ParseRawBody(c, verifyURL)
	if err != nil {
		return fmt.Errorf("%w, stage: %s",
			err, stageVerify)
	}
	date := regVerify.FindString(string(body))
	err = util.VerifyDate(date)
	if err != nil {
		return fmt.Errorf("%w, stage: %s",
			err, stageVerify)
	}
	return nil
}

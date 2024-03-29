package v2ex

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"

	"google.golang.org/protobuf/proto"

	"github.com/jdxj/sign/internal/pkg/util"
	pb "github.com/jdxj/sign/internal/proto/task"
)

const (
	domain    = ".v2ex.com"
	home      = "https://www.v2ex.com/"
	authURL   = "https://www.v2ex.com/balance"
	tokenURL  = "https://www.v2ex.com/mission/daily"
	signURL   = "https://www.v2ex.com/mission/daily/redeem?once=%s"
	verifyURL = "https://www.v2ex.com/balance"
)

const (
	stageAuth   = "auth"
	stageSignIn = "sign-in"
	stageQuery  = "query"
	stageVerify = "verify"
)

const (
	msgParseParamFailed = "解析参数失败"
	msgV2exSignInFailed = "v2ex签到失败"
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

type SignIn struct{}

func (si *SignIn) Kind() string {
	return pb.Kind_V2EX_SIGN_IN.String()
}

func (si *SignIn) Execute(body []byte) (string, error) {
	param := &pb.V2Ex{}
	err := proto.Unmarshal(body, param)
	if err != nil {
		return msgParseParamFailed, err
	}

	c, err := auth(param.GetCookie())
	if err != nil {
		return msgV2exSignInFailed, err
	}

	token, err := getSignToken(c)
	if err != nil {
		return msgV2exSignInFailed, err
	}

	err = signIn(c, token)
	if err != nil {
		return msgV2exSignInFailed, err
	}

	err = verify(c)
	if err != nil {
		return msgV2exSignInFailed, err
	}
	return "v2ex签到成功", nil
}

func auth(cookies string) (*http.Client, error) {
	jar := util.NewJar(cookies, domain, home)
	client := &http.Client{Jar: jar}

	body, err := util.ParseRawBody(client, authURL)
	if err != nil {
		return client, fmt.Errorf("%w, stage: %s", err, stageAuth)
	}

	target := regAuth.FindString(string(body))
	if target == "" {
		err = fmt.Errorf("%w, stage: %s", ErrTargetNotFound, stageAuth)
	}
	return client, err
}

func getSignToken(c *http.Client) (string, error) {
	body, err := util.ParseRawBody(c, tokenURL)
	if err != nil {
		return "", fmt.Errorf("%w, stage: %s", err, stageQuery)
	}

	matched := regToken.FindStringSubmatch(string(body))
	if len(matched) != 2 {
		return "", fmt.Errorf("%w, stage: %s", ErrTokenNotFound, stageQuery)
	}
	return matched[1], nil
}

func signIn(c *http.Client, token string) error {
	u := fmt.Sprintf(signURL, token)
	err := util.ParseBody(c, u, nil)
	if err != nil {
		return fmt.Errorf("%w, stage: %s", err, stageSignIn)
	}
	return nil
}

func verify(c *http.Client) error {
	body, err := util.ParseRawBody(c, verifyURL)
	if err != nil {
		return fmt.Errorf("%w, stage: %s", err, stageVerify)
	}
	date := regVerify.FindString(string(body))
	err = util.VerifyDate(date)
	if err != nil {
		err = fmt.Errorf("%w, stage: %s", err, stageVerify)
	}
	return err
}

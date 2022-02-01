package juejin

import (
	"fmt"
	"net/http"

	"github.com/jdxj/sign/internal/pkg/util"
	pb "github.com/jdxj/sign/internal/proto/task"
)

type SignIn struct{}

func (si *SignIn) Kind() string {
	return pb.Kind_JUEJIN_SIGN_IN.String()
}

func (si *SignIn) Execute(body []byte) (string, error) {
	param, err := parseParam(body)
	if err != nil {
		return msgParseParamFailed, err
	}

	jar := util.NewJar(param.GetCookie(), domain, home)
	client := &http.Client{Jar: jar}

	rsp := &response{
		Data: &checkIn{},
	}
	err = util.ParseBodyPost(client, signInURL, nil, rsp)
	if err != nil {
		return msgJueJinExecFailed, fmt.Errorf("%w, stage: %s", err, stageAuth)
	}
	if rsp.ErrNo != 0 {
		return msgJueJinExecFailed, fmt.Errorf("%w: %s, stage: %s",
			ErrUnknownMistake, rsp.ErrMsg, stageAuth)
	}
	return fmt.Sprintf("%s", rsp.Data), nil
}

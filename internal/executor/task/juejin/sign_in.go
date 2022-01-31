package juejin

import (
	"fmt"
	"net/http"

	"github.com/jdxj/sign/internal/pkg/util"
)

type SignIn struct{}

func (si *SignIn) Domain() crontab.Domain {
	return crontab.Domain_JueJin
}

func (si *SignIn) Kind() crontab.Kind {
	return crontab.Kind_JueJinSign
}

func (si *SignIn) Execute(key string) (string, error) {
	jar := util.NewJar(key, domain, home)
	client := &http.Client{Jar: jar}

	rsp := &response{
		Data: &checkIn{},
	}
	err := util.ParseBodyPost(client, signInURL, nil, rsp)
	if err != nil {
		return msgJueJinExecFailed, fmt.Errorf("%w, stage: %s", err, crontab.Stage_Auth)
	}
	if rsp.ErrNo != 0 {
		return msgJueJinExecFailed, fmt.Errorf("%w: %s, stage: %s",
			ErrUnknownMistake, rsp.ErrMsg, crontab.Stage_Auth)
	}
	return fmt.Sprintf("%s", rsp.Data), nil
}

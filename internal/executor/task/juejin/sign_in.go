package juejin

import (
	"fmt"
	"net/http"

	"github.com/jdxj/sign/internal/executor/task"
	"github.com/jdxj/sign/internal/proto/crontab"
)

type SignIn struct {
}

func (si *SignIn) Domain() crontab.Domain {
	return crontab.Domain_JueJin
}

func (si *SignIn) Kind() crontab.Kind {
	return crontab.Kind_JueJinSign
}

func (si *SignIn) Execute(key string) (string, error) {
	jar := task.NewJar(key, domain, home)
	client := &http.Client{Jar: jar}

	rsp := &response{
		Data: &checkIn{},
	}
	err := task.ParseBodyPost(client, signInURL, nil, rsp)
	if err != nil {
		return "", fmt.Errorf("%w, stage: %s", err, crontab.Stage_Auth)
	}
	if rsp.ErrNo != 0 {
		return "", fmt.Errorf("%s, stage: %s", rsp.ErrMsg, crontab.Stage_Auth)
	}
	return fmt.Sprintf("%s", rsp.Data), nil
}

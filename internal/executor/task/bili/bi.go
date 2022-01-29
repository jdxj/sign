package bili

import (
	"fmt"
	"net/http"

	"github.com/jdxj/sign/internal/executor/task"
)

type Bi struct{}

func (bi *Bi) Domain() crontab.Domain {
	return crontab.Domain_BILI
}

func (bi *Bi) Kind() crontab.Kind {
	return crontab.Kind_BILIBCount
}

func (bi *Bi) Execute(key string) (string, error) {
	c, err := auth(key)
	if err != nil {
		return msgGetBiFailed, err
	}
	return queryBi(c)
}

func queryBi(c *http.Client) (string, error) {
	biResp := &biResp{}
	err := task.ParseBody(c, biURL, biResp)
	if err != nil {
		return msgGetBiFailed, fmt.Errorf("stage: %s, error: %w",
			crontab.Stage_Query, err)
	}

	if biResp.Code != 0 {
		return msgGetBiFailed, fmt.Errorf("%w, stage: %s",
			ErrInvalidCookie, crontab.Stage_Query)
	}

	msg := fmt.Sprintf("硬币: %d", biResp.Data.Money)
	return msg, nil
}

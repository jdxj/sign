package bili

import (
	"fmt"
	"net/http"

	"github.com/jdxj/sign/internal/executor/task"
	"github.com/jdxj/sign/internal/proto/crontab"
)

type Bi struct {
}

func (bi *Bi) Domain() crontab.Domain {
	return crontab.Domain_BILI
}

func (bi *Bi) Kind() crontab.Kind {
	return crontab.Kind_BILIBCount
}

func (bi *Bi) Execute(key string) (string, error) {
	c, err := auth(key)
	if err != nil {
		return "", err
	}

	return queryBi(c)
}

type biResp struct {
	Code   int  `json:"code"`
	Status bool `json:"status"`
	Data   struct {
		Money int `json:"money"`
	} `json:"data"`
}

func queryBi(c *http.Client) (string, error) {
	biResp := &biResp{}
	err := task.ParseBody(c, biURL, biResp)
	if err != nil {
		return "", fmt.Errorf("stage: %s, error: %w",
			crontab.Stage_Query, err)
	}

	if biResp.Code != 0 {
		return "", fmt.Errorf("stage: %s, error: %s",
			crontab.Stage_Query, "Cookies 可能失效")
	}

	msg := fmt.Sprintf("硬币: %d", biResp.Data.Money)
	return msg, nil
}

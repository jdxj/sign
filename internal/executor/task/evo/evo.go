package evo

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jdxj/sign/internal/executor/task"
	"github.com/jdxj/sign/internal/proto/crontab"
)

type Updater struct{}

func (u *Updater) Domain() crontab.Domain {
	return crontab.Domain_Evo
}

func (u *Updater) Kind() crontab.Kind {
	return crontab.Kind_Raphael
}

func (u *Updater) Execute(key string) (string, error) {
	url := fmt.Sprintf(buildURL, key)
	bi := &buildInfo{}
	err := task.ParseBody(&http.Client{}, url, bi)
	if err != nil {
		return msgEvoUpdateFailed, err
	}
	updateTime := time.Unix(bi.Datetime, 0)
	if time.Since(updateTime) <= 24*time.Hour {
		return bi.String(), nil
	}
	return msgEvoUpdateFailed, ErrUpdateNotFound
}

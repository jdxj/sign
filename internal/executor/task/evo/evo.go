package evo

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jdxj/sign/internal/executor/task"
	"github.com/jdxj/sign/internal/proto/crontab"
)

type Updater struct {
}

func (u *Updater) Domain() crontab.Domain {
	// todo: proto
	return 0
}

func (u *Updater) Kind() crontab.Kind {
	// todo: proto
	return 0
}

func (u *Updater) Execute(key string) (string, error) {
	url := fmt.Sprintf(buildURL, key)
	bi := &buildInfo{}
	err := task.ParseBody(&http.Client{}, url, bi)
	if err != nil {
		return "", err
	}

	updateTime := time.Unix(bi.Datetime, 0)
	if time.Since(updateTime) <= 24*time.Hour {
		return bi.String(), nil
	}
	return "", fmt.Errorf("update not found")
}

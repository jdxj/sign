package executor

import (
	"github.com/jdxj/sign/internal/executor/task/bili"
	"github.com/jdxj/sign/internal/proto/crontab"
)

func init() {
	biSI := &bili.SignIn{}
	agents[biSI.Kind()] = biSI
}

var (
	agents = make(map[crontab.Kind]Agent)
)

type Agent interface {
	Domain() crontab.Domain
	Kind() crontab.Kind
	Execute(key string) (string, error)
}

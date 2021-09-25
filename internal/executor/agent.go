package executor

import (
	"github.com/jdxj/sign/internal/executor/task/bili"
	"github.com/jdxj/sign/internal/executor/task/hpi"
	"github.com/jdxj/sign/internal/executor/task/stg"
	"github.com/jdxj/sign/internal/executor/task/v2ex"
	"github.com/jdxj/sign/internal/proto/crontab"
)

func init() {
	biSI := &bili.SignIn{}
	agents[biSI.Kind()] = biSI

	hpiSI := &hpi.SignIn{}
	agents[hpiSI.Kind()] = hpiSI

	stgSI := &stg.SignIn{}
	agents[stgSI.Kind()] = stgSI

	v2exSI := &v2ex.SignIn{}
	agents[v2exSI.Kind()] = v2exSI
}

var (
	agents = make(map[crontab.Kind]Agent)
)

type Agent interface {
	Domain() crontab.Domain
	Kind() crontab.Kind
	Execute(key string) (string, error)
}

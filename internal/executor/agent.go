package executor

import (
	"github.com/jdxj/sign/internal/executor/task/bili"
	"github.com/jdxj/sign/internal/executor/task/evo"
	"github.com/jdxj/sign/internal/executor/task/github"
	"github.com/jdxj/sign/internal/executor/task/hpi"
	"github.com/jdxj/sign/internal/executor/task/stg"
	"github.com/jdxj/sign/internal/executor/task/v2ex"
	"github.com/jdxj/sign/internal/proto/crontab"
)

func init() {
	biSI := &bili.SignIn{}
	agents[biSI.Kind()] = biSI

	bi := &bili.Bi{}
	agents[bi.Kind()] = bi

	hpiSI := &hpi.SignIn{}
	agents[hpiSI.Kind()] = hpiSI

	stgSI := &stg.SignIn{}
	agents[stgSI.Kind()] = stgSI

	v2exSI := &v2ex.SignIn{}
	agents[v2exSI.Kind()] = v2exSI

	updater := &evo.Updater{}
	agents[updater.Kind()] = updater

	release := &github.Release{}
	agents[release.Kind()] = release
}

var (
	agents = make(map[crontab.Kind]Agent)
)

type Agent interface {
	Domain() crontab.Domain
	Kind() crontab.Kind
	Execute(key string) (string, error)
}

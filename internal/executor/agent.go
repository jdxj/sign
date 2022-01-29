package executor

import (
	"net/http"

	"github.com/jdxj/sign/internal/executor/task/bili"
	"github.com/jdxj/sign/internal/executor/task/evo"
	"github.com/jdxj/sign/internal/executor/task/github"
	"github.com/jdxj/sign/internal/executor/task/hpi"
	"github.com/jdxj/sign/internal/executor/task/juejin"
	"github.com/jdxj/sign/internal/executor/task/nokey"
	"github.com/jdxj/sign/internal/executor/task/stg"
	"github.com/jdxj/sign/internal/executor/task/v2ex"

	"net/url"
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

	jueSI := &juejin.SignIn{}
	agents[jueSI.Kind()] = jueSI
	jueCo := &juejin.Count{}
	agents[jueCo.Kind()] = jueCo
	juePo := &juejin.Point{}
	agents[juePo.Kind()] = juePo
	jueCa := &juejin.Calendar{}
	agents[jueCa.Kind()] = jueCa

	cm := &nokey.CustomMessage{}
	agents[cm.Kind()] = cm
	url.Values{}
	req := http.Request{}
	req.ParseForm()
}

var (
	agents = make(map[crontab.Kind]Agent)
)

type Agent interface {
	Domain() crontab.Domain
	Kind() crontab.Kind
	Execute(key string) (string, error)
}

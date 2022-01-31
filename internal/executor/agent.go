package executor

import (
	"net/http"
	"net/url"

	bilibili2 "github.com/jdxj/sign/internal/task/service/executor/bilibili"
	custom_msg "github.com/jdxj/sign/internal/task/service/executor/custom-message"
	"github.com/jdxj/sign/internal/task/service/executor/evo"
	"github.com/jdxj/sign/internal/task/service/executor/github"
	"github.com/jdxj/sign/internal/task/service/executor/hpi"
	juejin2 "github.com/jdxj/sign/internal/task/service/executor/juejin"
	"github.com/jdxj/sign/internal/task/service/executor/stg"
	"github.com/jdxj/sign/internal/task/service/executor/v2ex"
)

func init() {
	biSI := &bilibili2.SignIn{}
	agents[biSI.Kind()] = biSI

	bi := &bilibili2.Bi{}
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

	jueSI := &juejin2.SignIn{}
	agents[jueSI.Kind()] = jueSI
	jueCo := &juejin2.Count{}
	agents[jueCo.Kind()] = jueCo
	juePo := &juejin2.Point{}
	agents[juePo.Kind()] = juePo
	jueCa := &juejin2.Calendar{}
	agents[jueCa.Kind()] = jueCa

	cm := &custom_msg.CustomMessage{}
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

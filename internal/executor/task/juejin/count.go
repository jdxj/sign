package juejin

import (
	"github.com/jdxj/sign/internal/proto/crontab"
)

type Count struct{}

func (cou *Count) Domain() crontab.Domain {
	return crontab.Domain_JueJin
}

func (cou *Count) Kind() crontab.Kind {
	return crontab.Kind_JueJinCount
}

func (cou *Count) Execute(key string) (string, error) {
	return execute(key, countsURL, &counts{})
}

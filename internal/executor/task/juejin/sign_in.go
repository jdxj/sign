package juejin

import (
	"github.com/jdxj/sign/internal/pkg/util"
	"github.com/jdxj/sign/internal/proto/crontab"
)

type SignIn struct {
}

func (si *SignIn) Domain() crontab.Domain {
	// todo: build proto
	return 701
}

func (si *SignIn) Kind() crontab.Kind {
	// todo: build proto
	return 702
}

const (
	authURL   = ""
	signInURL = ""
)

type response struct {
}

func (si *SignIn) Execute(key string) (string, error) {
	util.GetJson()
}

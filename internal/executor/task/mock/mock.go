package mock

import (
	"github.com/jdxj/sign/internal/pkg/logger"
)

type PrintAgent struct {
}

func (pa *PrintAgent) Domain() int {
	return 1000
}

func (pa *PrintAgent) Kind() int {
	return 1001
}

func (pa *PrintAgent) SignIn(key string) error {
	logger.Debugf("mock, key: %s", key)
	return nil
}

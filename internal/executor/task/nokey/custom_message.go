package nokey

import (
	"github.com/jdxj/sign/internal/proto/crontab"
)

type CustomMessage struct{}

func (cm *CustomMessage) Domain() crontab.Domain {
	return crontab.Domain_NoKey
}

func (cm *CustomMessage) Kind() crontab.Kind {
	return crontab.Kind_CustomMessage
}

func (cm *CustomMessage) Execute(key string) (string, error) {
	return key, nil
}

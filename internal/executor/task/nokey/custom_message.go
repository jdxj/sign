package nokey

import (
	"github.com/jdxj/sign/internal/proto/crontab"
)

type CustomMessage struct{}

func (cm *CustomMessage) Domain() crontab.Domain {
	// todo: compile proto
	return 801
}

func (cm *CustomMessage) Kind() crontab.Kind {
	// todo: compile proto
	return 802
}

func (cm *CustomMessage) Execute(key string) (string, error) {
	return key, nil
}

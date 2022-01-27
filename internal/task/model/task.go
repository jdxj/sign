package model

import (
	"github.com/jdxj/sign/internal/proto/task"
)

type Executor interface {
	Kind() string
	Inputs() []*task.Input
	Execute([]byte) (string, error)
}

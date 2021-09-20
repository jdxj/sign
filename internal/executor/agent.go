package executor

import (
	"errors"

	"github.com/jdxj/sign/internal/executor/task/mock"
)

var (
	ErrAlreadyRegistered = errors.New("already registered")
	ErrInvalidKind       = errors.New("invalid kind")
)

var (
	agents = registerDefaultAgent()
)

func RegisterAgent(agent Agent) {
	if agent.Kind() == 0 {
		panic(ErrInvalidKind)
	}
	_, ok := agents[agent.Kind()]
	if ok {
		panic(ErrAlreadyRegistered)
	}
	agents[agent.Kind()] = agent
}

type Agent interface {
	Domain() int
	Kind() int
	SignIn(key string) error
}

func registerDefaultAgent() map[int]Agent {
	agents := make(map[int]Agent)

	// todo: test
	pa := &mock.PrintAgent{}
	agents[pa.Kind()] = pa

	return agents
}

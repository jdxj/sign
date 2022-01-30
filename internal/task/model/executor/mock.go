package executor

import (
	pb "github.com/jdxj/sign/internal/proto/task"
)

type mockExecutor struct {
}

func (m *mockExecutor) Kind() string {
	return pb.Kind_MOCK.String()
}

func (m *mockExecutor) Execute(param []byte) (string, error) {
	return "hello", nil
}

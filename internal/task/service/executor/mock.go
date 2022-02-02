package executor

import (
	pb "github.com/jdxj/sign/internal/proto/task"
)

type MockExecutor struct {
}

func (m *MockExecutor) Kind() string {
	return pb.Kind_MOCK.String()
}

func (m *MockExecutor) Execute(param []byte) (string, error) {
	return "hello", nil
}

package juejin

import (
	pb "github.com/jdxj/sign/internal/proto/task"
)

type Count struct{}

func (cou *Count) Kind() string {
	return pb.Kind_JUEJIN_COUNT.String()
}

func (cou *Count) Execute(body []byte) (string, error) {
	return execute(body, countsURL, &counts{})
}

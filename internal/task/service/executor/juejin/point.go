package juejin

import (
	pb "github.com/jdxj/sign/internal/proto/task"
)

type Point struct{}

func (p *Point) Kind() string {
	return pb.Kind_JUEJIN_POINT.String()
}

func (p *Point) Execute(body []byte) (string, error) {
	var o ore
	return execute(body, pointURL, &o)
}

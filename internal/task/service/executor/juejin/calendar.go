package juejin

import (
	pb "github.com/jdxj/sign/internal/proto/task"
)

type Calendar struct{}

func (cal *Calendar) Kind() string {
	return pb.Kind_JUEJIN_CALENDAR.String()
}

func (cal *Calendar) Execute(body []byte) (string, error) {
	return execute(body, calendarURL, &jokes{})
}

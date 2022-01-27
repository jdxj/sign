package service

import (
	"errors"

	"github.com/robfig/cron/v3"
	"github.com/streadway/amqp"

	"github.com/jdxj/sign/internal/proto/task"
)

const ()

var (
	ErrKindNotFound = errors.New("kind not found")
)

func NewService() *Service {
	srv := &Service{
		cronParser: cron.NewParser(
			cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow),
	}
	return srv
}

type Service struct {
	task.UnimplementedTaskServiceServer

	cronParser cron.Parser
	mqChannel  *amqp.Channel
}

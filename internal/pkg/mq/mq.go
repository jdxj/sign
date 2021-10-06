package mq

import (
	"fmt"

	"github.com/streadway/amqp"

	"github.com/jdxj/sign/internal/pkg/config"
)

var (
	connRabbit *amqp.Connection
)

func InitRabbit(conf config.Rabbit) {
	var (
		dsn = fmt.Sprintf("amqp://%s:%s@%s:%d", conf.User, conf.Pass, conf.Host, conf.Port)
		err error
	)
	connRabbit, err = amqp.Dial(dsn)
	if err != nil {
		panic(err)
	}
}

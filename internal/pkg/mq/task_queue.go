package mq

import (
	"sync"

	"github.com/streadway/amqp"

	"github.com/jdxj/sign/internal/pkg/logger"
)

const (
	taskExchange     = "task-exchange"
	taskExchangeKind = "topic"
	taskQueue        = "task-Queue"
	taskRoutingKey   = "task"
)

func NewTaskQueue() (*TaskQueue, error) {
	channel, err := Conn.Channel()
	if err != nil {
		return nil, err
	}
	tq := &TaskQueue{
		channel: channel,
		wg:      &sync.WaitGroup{},
	}
	return tq, tq.newQueue()

}

type TaskQueue struct {
	channel *amqp.Channel

	wg *sync.WaitGroup
}

func (tq *TaskQueue) newQueue() error {
	err := tq.channel.ExchangeDeclare(taskExchange, taskExchangeKind, true, false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	queue, err := tq.channel.QueueDeclare(taskQueue, true, false, false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	err = tq.channel.QueueBind(queue.Name, taskRoutingKey, taskExchange, false, nil)
	return err
}

func (tq *TaskQueue) Stop() {
	_ = tq.channel.Close()
	tq.wg.Wait()
}

func (tq *TaskQueue) Publish(body []byte) error {
	msg := amqp.Publishing{
		DeliveryMode: 2,
		Body:         body,
	}
	return tq.channel.Publish(taskExchange, taskRoutingKey, false, false, msg)
}

func (tq *TaskQueue) Consume() (<-chan []byte, error) {
	deliveryChan, err := tq.channel.Consume(taskQueue, "", true, false, false,
		false, nil)
	if err != nil {
		return nil, err
	}

	bodyChan := make(chan []byte)
	tq.wg.Add(1)
	go func() {
		defer tq.wg.Done()

		for msg := range deliveryChan {
			bodyChan <- msg.Body
		}
		close(bodyChan)
		logger.Debugf("task chan stopped")
	}()
	return bodyChan, nil
}

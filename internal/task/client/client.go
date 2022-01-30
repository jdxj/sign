package client

import (
	"github.com/panjf2000/ants/v2"
	"go-micro.dev/v4/broker"

	"github.com/jdxj/sign/internal/proto/notice"
	"github.com/jdxj/sign/internal/proto/trigger"
)

var (
	TriggerService trigger.TriggerService
	NoticeService  notice.NoticeService
)

var (
	GPool *ants.Pool
	MQ    broker.Broker
)

package main

import (
	"os"

	"github.com/spf13/pflag"

	"github.com/jdxj/sign/internal/pkg/config"
	"github.com/jdxj/sign/internal/pkg/db"
	"github.com/jdxj/sign/internal/pkg/logger"
	"github.com/jdxj/sign/internal/pkg/mq"
	"github.com/jdxj/sign/internal/pkg/util"
	"github.com/jdxj/sign/internal/trigger"
)

const (
	serviceName = "trigger"
)

func main() {
	flagSet := pflag.NewFlagSet(serviceName, pflag.ExitOnError)
	file := flagSet.StringP("file", "f", "config.yaml", "configure path")
	_ = flagSet.Parse(os.Args) // 忽略 err, 因为使用了 ExitOnError

	root := config.ReadConfigs(*file)
	logger.Init(root.Logger.Path+serviceName+".log",
		logger.WithMode(root.Logger.Mode))

	dbConf := root.DB
	db.InitGorm(dbConf)

	rabbitConf := root.Rabbit
	mq.InitRabbit(rabbitConf)

	trg := trigger.New(dbConf)
	trg.Start()
	util.Hold()
	trg.Stop()
}

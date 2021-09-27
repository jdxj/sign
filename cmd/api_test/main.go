package main

import (
	"github.com/jdxj/sign/internal/pkg/api"
	"github.com/jdxj/sign/internal/pkg/config"
	"github.com/jdxj/sign/internal/pkg/logger"
	"github.com/jdxj/sign/internal/pkg/util"
)

func main() {
	logger.Init("./api.log")
	conf := config.APIServer{
		Host: "127.0.0.1",
		Port: 8080,
		Key:  "123",
	}

	s := api.NewServer(conf, api.NewRouter())
	s.Start()
	util.Hold()
	s.Stop()
}

package apiserver

import (
	"github.com/jdxj/sign/internal/pkg/config"
)

var (
	jwtKey string
)

func Init(conf config.APIServer) {
	jwtKey = conf.Key
}

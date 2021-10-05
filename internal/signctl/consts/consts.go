package consts

import (
	"errors"
)

// flag name
const (
	Host     = "host"
	Token    = "token"
	Nickname = "nickname"
	Password = "password"
)

// api path
const (
	prefix  = "/api"
	version = "/v1"

	CreateUser = prefix + "/user"
	AuthUser   = prefix + "/token"
)

var (
	ErrPostJson = errors.New("post json failed")
)

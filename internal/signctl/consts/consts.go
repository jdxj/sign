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
	Describe = "describe"
	Domain   = "domain"
	Key      = "key"
	SecretID = "secret-id"
)

// api path
const (
	prefix  = "/api"
	version = "/v1"

	CreateUser = prefix + "/user"
	AuthUser   = prefix + "/token"

	CreateSecret = prefix + version + "/secret"
	UpdateSecret = CreateSecret
	DeleteSecret = CreateSecret
)

var (
	ErrSendJson     = errors.New("send json failed")
	ErrInvalidParam = errors.New("invalid param")
)

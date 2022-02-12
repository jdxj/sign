package consts

import (
	"errors"
)

// flag name
const (
	Host  = "host"
	Token = "token"
	Debug = "debug"

	Nickname = "nickname"
	Password = "password"
	Mail     = "mail"
	Telegram = "telegram"

	Description = "description"

	Kind      = "kind"
	Spec      = "spec"
	Param     = "param"
	TaskID    = "task-id"
	CreatedAt = "created_at"
	PageID    = "page-id"
	PageSize  = "page-size"
)

// api path
const (
	prefix  = "/api"
	version = "/v1"

	users    = prefix + version + "/users"
	ApiUser  = users + "/user"
	ApiToken = users + "/token"

	ApiTasks = prefix + version + "/tasks"
)

var (
	ErrSendJson     = errors.New("send json failed")
	ErrInvalidParam = errors.New("invalid param")
)

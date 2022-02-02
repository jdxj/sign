package consts

import (
	"errors"
)

// flag name
const (
	Host  = "host"
	Token = "token"

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

	session      = prefix + version + "/session"
	SessionLogin = session + "/login"

	user       = prefix + version + "/user"
	UserCreate = user + "/sign-up"
	UserUpdate = user + "/update"

	task       = prefix + version + "/task"
	TaskCreate = task + "/create"
	TaskGet    = task + "/get"
	TaskList   = task + "/list"
	TaskUpdate = task + "/update"
	TaskDelete = task + "/delete"
)

var (
	ErrSendJson     = errors.New("send json failed")
	ErrInvalidParam = errors.New("invalid param")
)

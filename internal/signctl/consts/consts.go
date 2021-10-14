package consts

import (
	"errors"
)

// flag name
const (
	Host     = "host"
	Token    = "token"
	Format   = "format"
	Nickname = "nickname"
	Password = "password"
	Describe = "describe"
	Domain   = "domain"
	Kind     = "kind"
	Spec     = "spec"
	Key      = "key"
	SecretID = "secret-id"
	TaskID   = "task-id"
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
	GetSecrets   = prefix + version + "/secrets"

	CreateTask = prefix + version + "/task"
	DeleteTask = CreateTask
	GetTasks   = prefix + version + "/tasks"
)

var (
	ErrSendJson     = errors.New("send json failed")
	ErrInvalidParam = errors.New("invalid param")
)

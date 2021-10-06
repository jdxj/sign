package model

import (
	"encoding/json"
)

type Request struct {
	Token string      `json:"token"`
	Data  interface{} `json:"data"`
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (rsp *Response) String() string {
	data, _ := json.MarshalIndent(rsp, "", "  ")
	return string(data)
}

type CreateUserReq struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

type CreateUserRsp struct {
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
}

type AuthReq struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

type AuthRsp struct {
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
}

type CreateSecretReq struct {
	Describe string `json:"describe"`
	Domain   int    `json:"domain"`
	Key      string `json:"key"`
}

type CreateSecretRsp struct {
	SecretID int64 `json:"secret_id"`
}

type UpdateSecretReq struct {
	SecretID int64  `json:"secret_id"`
	Describe string `json:"describe"`
	Domain   int    `json:"domain"`
	Key      string `json:"key"`
}

type DeleteSecretReq struct {
	SecretID int64 `json:"secret_id"`
}

type GetSecretsReq struct {
	SecretIDs []int64 `json:"secret_ids"`
	Domains   []int   `json:"domains"`
}

type Secret struct {
	SecretID int64  `json:"secret_id"`
	Describe string `json:"describe"`
	Domain   int    `json:"domain"`
	Key      string `json:"key"`
}

type GetSecretsRsp struct {
	List []*Secret `json:"list"`
}

type CreateTaskReq struct {
	Describe string `json:"describe"`
	Kind     int    `json:"kind"`
	Spec     string `json:"spec"`
	SecretID int64  `json:"secret_id"`
}

type CreateTaskRsp struct {
	TaskID int64 `json:"task_id"`
}

type DeleteTaskReq struct {
	TaskID int64 `json:"task_id"`
}

type GetTasksReq struct {
	Kinds     []int   `json:"kinds"`
	SecretIDs []int64 `json:"secret_ids"`
}

type Task struct {
	TaskID   int64  `json:"task_id"`
	Describe string `json:"describe"`
	Kind     int    `json:"kind"`
	Spec     string `json:"spec"`
	SecretID int64  `json:"secret_id"`
}

type GetTasksRsp struct {
	List []*Task `json:"list"`
}

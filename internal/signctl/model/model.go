package model

import (
	"encoding/json"
)

type Request struct {
	Token string      `json:"token"`
	Data  interface{} `json:"data"`
}

type Response struct {
	Code int         `json:"code"`
	Desc string      `json:"desc"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (rsp *Response) String() string {
	data, _ := json.MarshalIndent(rsp, "", "  ")
	return string(data)
}

type CreateUserReq struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	Mail     string `json:"mail"`
	Telegram int64  `json:"telegram"`
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

type UpdateTaskReq struct {
	Desc  string          `json:"desc"`
	Spec  string          `json:"spec"`
	Param json.RawMessage `json:"param"`
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
	Describe string          `json:"describe"`
	Kind     string          `json:"kind"`
	Spec     string          `json:"spec"`
	Param    json.RawMessage `json:"param"`
}

type CreateTaskRsp struct {
	TaskID int64 `json:"task_id"`
}

type DeleteTaskReq struct {
	TaskID int64 `json:"task_id"`
}

type GetTasksReq struct {
	Desc      string `json:"desc"`
	Kind      string `json:"kind"`
	Spec      string `json:"spec"`
	CreatedAt int64  `json:"created_at"`
	PageID    int64  `json:"page_id"`
	PageSize  int64  `json:"page_size"`
}

type Task struct {
	TaskID   int64  `json:"task_id"`
	Describe string `json:"describe"`
	Kind     int    `json:"kind"`
	Spec     string `json:"spec"`
	Param    string `json:"param"`
}

type GetTasksRsp struct {
	Count int64   `json:"count"`
	List  []*Task `json:"tasks"`
}

type UpdateUserReq struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	Mail     string `json:"mail"`
	Telegram int64  `json:"telegram"`
}

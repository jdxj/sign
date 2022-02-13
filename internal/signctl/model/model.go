package model

import (
	"encoding/json"
	"fmt"
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

type UpdateTaskReq struct {
	Desc  string          `json:"desc"`
	Spec  string          `json:"spec"`
	Param json.RawMessage `json:"param"`
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

type GetTasksReq struct {
	Desc      string `json:"desc"`
	Kind      string `json:"kind"`
	Spec      string `json:"spec"`
	CreatedAt int64  `json:"created_at"`
	PageID    int64  `json:"page_id"`
	PageSize  int64  `json:"page_size"`
}

type Task struct {
	TaskID   int64           `json:"task_id"`
	Describe string          `json:"describe"`
	Kind     string          `json:"kind"`
	Spec     string          `json:"spec"`
	Param    json.RawMessage `json:"param"`
}

type GetTasksRsp struct {
	Count    int64   `json:"count"`
	PageID   int64   `json:"-"`
	PageSize int64   `json:"-"`
	List     []*Task `json:"tasks"`
}

func (gtr *GetTasksRsp) String() string {
	result := fmt.Sprintf("\n%-7s %-15s %-15s %-15s %-15s\n",
		"task id", "description", "kind", "spec", "param")
	for _, v := range gtr.List {
		param := v.Param
		if len(param) > 48 {
			param = append(param[:45], "..."...)
		}
		result = fmt.Sprintf("%s%7d %-15s %-15s %-15s %-15s\n",
			result, v.TaskID, v.Describe, v.Kind, v.Spec, param)
	}
	result = fmt.Sprintf("%s%5s %7s %9s\n", result, "count", "page id", "page size")
	result = fmt.Sprintf("%s%5d %7d %9d\n", result, gtr.Count, gtr.PageID, gtr.PageSize)
	return result
}

type UpdateUserReq struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	Mail     string `json:"mail"`
	Telegram int64  `json:"telegram"`
}

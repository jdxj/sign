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
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (rsp *Response) String() string {
	data, _ := json.MarshalIndent(rsp, "", "  ")
	return fmt.Sprintf("%s", data)
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

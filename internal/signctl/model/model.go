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

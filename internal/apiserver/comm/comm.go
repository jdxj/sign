package comm

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/jdxj/sign/internal/pkg/code"
)

type Request struct {
	Token string          `json:"token"`
	Data  json.RawMessage `json:"data"`
}

func NewResponse(code int, msg string, data interface{}) *Response {
	return &Response{
		Code:    code,
		Message: msg,
		Data:    data,
	}
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type AddTaskReq struct {
	ID     string `json:"id"`
	Domain int    `json:"domain"`
	Type   []int  `json:"type"`
	Key    string `json:"key"`
}

func UnmarshalRequest(ctx *gin.Context, req interface{}) error {
	data := ctx.GetString("data")
	err := json.Unmarshal([]byte(data), req)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			NewResponse(
				code.ErrBindReqFailed,
				"unmarshal request failed",
				err))
	}
	return err
}

type Claim struct {
	UserID   int64
	Nickname string
	jwt.StandardClaims
}

func GenerateToken(claim *Claim) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	// todo: 从配置中获取
	key := []byte("20210910")
	return token.SignedString(key)
}

func CheckToken(sign string) (*Claim, error) {
	token, err := jwt.ParseWithClaims(sign, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		key := []byte("20210910")
		return key, nil
	})
	if err != nil {
		return nil, err
	}

	claim, ok := token.Claims.(*Claim)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("parse token failed")
	}
	return claim, nil
}

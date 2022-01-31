package api

import (
	"github.com/gin-gonic/gin"

	ser "github.com/jdxj/sign/internal/pkg/sign-error"
)

type handler func(*Request) (interface{}, error)
type checker func(*Request) error

type Request struct {
	Token string      `json:"token"`
	Data  interface{} `json:"data"`

	claim *Claim
}

func checkToken(req *Request) (err error) {
	req.claim, err = NewClaimFromToken(req.Token)
	if err != nil {
		return ser.New(ser.ErrAuthFailed, "无效的 token")
	}
	return
}

func Process(ctx *gin.Context, d interface{}, handle handler, checkers ...checker) {
	req := &Request{
		Data: d,
	}
	if err := ctx.Bind(req); err != nil {
		Respond(ctx, nil, ser.New(ser.ErrBindReqFailed, err.Error()))
		return
	}

	for _, check := range checkers {
		if err := check(req); err != nil {
			Respond(ctx, nil, err)
			return
		}
	}

	data, err := handle(req)
	Respond(ctx, data, err)
}

func ProcessCheckToken(ctx *gin.Context, d interface{}, handle handler) {
	Process(ctx, d, handle, checkToken)
}

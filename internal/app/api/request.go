package api

import (
	"github.com/gin-gonic/gin"

	"github.com/jdxj/sign/internal/pkg/logger"
	ser "github.com/jdxj/sign/internal/pkg/sign-error"
)

const (
	SignClaimKey = "sign-claim"
)

type handler func(*Request) (interface{}, error)
type checker func(*Request) error

// Request
// Deprecated
type Request struct {
	Token string      `json:"token"`
	Data  interface{} `json:"data"`

	Claim *SignClaim `json:"-"`
}

func checkToken(req *Request) (err error) {
	req.Claim, err = NewSignClaimFromToken(req.Token)
	if err != nil {
		return ser.New(ser.ErrUnknown, "无效的 token")
	}
	return
}

// Process
// Deprecated
func Process(ctx *gin.Context, d interface{}, handle handler, checkers ...checker) {
	req := &Request{
		Data: d,
	}
	if err := ctx.Bind(req); err != nil {
		logger.Errorf("Bind: %s", err)
		Respond(ctx, nil, ser.New(ser.ErrBindRequest, err.Error()))
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

// ProcessCheckToken
// Deprecated
func ProcessCheckToken(ctx *gin.Context, d interface{}, handle handler) {
	Process(ctx, d, handle, checkToken)
}

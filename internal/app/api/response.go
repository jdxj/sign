package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	ser "github.com/jdxj/sign/internal/pkg/sign-error"
)

type Response struct {
	Code        int         `json:"code"`
	Description string      `json:"description"`
	Data        interface{} `json:"data"`
}

func Respond(ctx *gin.Context, data interface{}, err error) {
	rsp := &Response{
		Data: data,
	}

	if err != nil {
		var se *ser.SignError
		if errors.As(err, &se) {
			rsp.Code = se.Code
			rsp.Description = fmt.Sprintf("%s - %s", se.CodeDesc, se.Description)
		} else {
			rsp.Code = ser.ErrUnknown
			rsp.Description = err.Error()
		}
	}
	ctx.JSON(http.StatusOK, rsp)
}

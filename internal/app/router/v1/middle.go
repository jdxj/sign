package v1

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/jdxj/sign/internal/app/api"
	ser "github.com/jdxj/sign/internal/pkg/sign-error"
)

const (
	headerAuthorization = "Authorization"
)

func authnBearer(ctx *gin.Context) {
	bearer := ctx.Request.Header.Get(headerAuthorization)
	if bearer == "" {
		err := ser.New(ser.ErrAuth, "http header '%s' not found", headerAuthorization)
		api.Respond(ctx, nil, err)
		ctx.Abort()
		return
	}

	var token string
	_, _ = fmt.Sscanf(bearer, "bearer %s", &token)
	claim, err := api.NewSignClaimFromToken(token)
	if err != nil {
		err := ser.Wrap(ser.ErrAuth, err, "NewSignClaimFromToken")
		api.Respond(ctx, nil, err)
		ctx.Abort()
		return
	}
	ctx.Set(api.SignClaimKey, claim)
}

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
		api.Abort(
			ctx,
			nil,
			ser.New(ser.ErrAuth, "http header '%s' not found", headerAuthorization),
		)
		return
	}

	var token string
	_, _ = fmt.Sscanf(bearer, "bearer %s", &token)
	claim, err := api.NewSignClaimFromToken(token)
	if err != nil {
		api.Abort(
			ctx,
			nil,
			ser.Wrap(ser.ErrAuth, err, "NewSignClaimFromToken"),
		)
		return
	}
	ctx.Set(api.SignClaimKey, claim)
}

func getUserID(ctx *gin.Context) int64 {
	v, exists := ctx.Get(api.SignClaimKey)
	if exists {
		return v.(*api.SignClaim).UserID
	}
	return 0
}

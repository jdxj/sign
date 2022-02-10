package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/jdxj/sign/internal/app/api"
)

func getUserID(ctx *gin.Context) int64 {
	v, exists := ctx.Get(api.SignClaimKey)
	if exists {
		return v.(*api.SignClaim).UserID
	}
	return 0
}

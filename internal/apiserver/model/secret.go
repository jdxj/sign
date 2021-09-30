package model

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jdxj/sign/internal/apiserver"
	"github.com/jdxj/sign/internal/proto/crontab"
	"github.com/jdxj/sign/internal/proto/secret"
)

type CreateSecretReq struct {
	UserID int64          `json:"user_id"`
	Domain crontab.Domain `json:"domain"`
	Key    string         `json:"key"`
}

type CreateSecretRsp struct {
	SecretID int64 `json:"secret_id"`
}

func CreateSecret(ctx *gin.Context) {
	req := &CreateSecretReq{}
	apiserver.Handle(ctx, req, func(tCtx context.Context) (interface{}, error) {
		return createSecret(tCtx, req)
	})
}

func createSecret(ctx context.Context, req *CreateSecretReq) (*CreateSecretRsp, error) {
	secretRsp, err := apiserver.SecretClient.CreateSecret(ctx, &secret.CreateSecretReq{
		UserID: req.UserID,
		Domain: req.Domain,
		Key:    req.Key,
	})
	if err != nil {
		return nil, err
	}
	rsp := &CreateSecretRsp{
		SecretID: secretRsp.SecretID,
	}
	return rsp, nil
}

type UpdateSecretReq struct {
	SecretID int64          `json:"secret_id"`
	Domain   crontab.Domain `json:"domain"`
	Key      string         `json:"key"`
}

type UpdateSecretRsp struct {
}

func UpdateSecret(ctx *gin.Context) {
	req := &UpdateSecretReq{}
	apiserver.Handle(ctx, req, func(tCtx context.Context) (interface{}, error) {
		return nil, updateSecret(tCtx, req)
	})
}

func updateSecret(ctx context.Context, req *UpdateSecretReq) error {
	_, err := apiserver.SecretClient.UpdateSecret(ctx, &secret.UpdateSecretReq{
		SecretID: req.SecretID,
		Domain:   req.Domain,
		Key:      req.Key,
	})
	return err
}

func GetSecret(ctx *gin.Context) {
	// todo: get secret or secrets?
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "todo: get secret or secrets?",
	})
}

type DeleteSecretReq struct {
	SecretID int64 `json:"secret_id"`
}

func DeleteSecret(ctx *gin.Context) {
	req := &DeleteSecretReq{}
	apiserver.Handle(ctx, req, func(tCtx context.Context) (interface{}, error) {
		return nil, deleteSecret(tCtx, req)
	})
}

func deleteSecret(ctx context.Context, req *DeleteSecretReq) error {
	_, err := apiserver.SecretClient.DeleteSecret(ctx, &secret.DeleteSecretReq{
		SecretID: req.SecretID,
	})
	return err
}

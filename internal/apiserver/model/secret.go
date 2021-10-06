package model

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/jdxj/sign/internal/apiserver"
	"github.com/jdxj/sign/internal/proto/crontab"
	"github.com/jdxj/sign/internal/proto/secret"
)

type CreateSecretReq struct {
	Describe string         `json:"describe"`
	Domain   crontab.Domain `json:"domain"`
	Key      string         `json:"key"`
}

type CreateSecretRsp struct {
	SecretID int64 `json:"secret_id"`
}

func CreateSecret(ctx *gin.Context) {
	req := &CreateSecretReq{}
	value, _ := ctx.Get(apiserver.KeyClaim)
	apiserver.Handle(ctx, req, func(tCtx context.Context) (interface{}, error) {
		return createSecret(tCtx, req, value.(*apiserver.Claim).UserID)
	})
}

func createSecret(ctx context.Context, req *CreateSecretReq, userID int64) (*CreateSecretRsp, error) {
	secretRsp, err := apiserver.SecretClient.CreateSecret(ctx, &secret.CreateSecretReq{
		UserID:   userID,
		Describe: req.Describe,
		Domain:   req.Domain,
		Key:      req.Key,
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
	Describe string         `json:"describe"`
	Domain   crontab.Domain `json:"domain"`
	Key      string         `json:"key"`
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
		Describe: req.Describe,
		Domain:   req.Domain,
		Key:      req.Key,
	})
	return err
}

type GetSecretsReq struct {
	SecretIDs []int64          `json:"secret_ids"`
	Domains   []crontab.Domain `json:"domains"`
}

type Secret struct {
	SecretID int64          `json:"secret_id"`
	Describe string         `json:"describe"`
	Domain   crontab.Domain `json:"domain"`
	Key      string         `json:"key"`
}

type GetSecretsRsp struct {
	List []*Secret `json:"list"`
}

func GetSecrets(ctx *gin.Context) {
	req := &GetSecretsReq{}
	value, _ := ctx.Get(apiserver.KeyClaim)
	apiserver.Handle(ctx, req, func(tCtx context.Context) (interface{}, error) {
		return getSecrets(tCtx, req, value.(*apiserver.Claim).UserID)
	})
}

func getSecrets(ctx context.Context, req *GetSecretsReq, userID int64) (*GetSecretsRsp, error) {
	secretList, err := apiserver.SecretClient.GetSecretList(ctx, &secret.GetSecretListReq{
		SecretIDs: req.SecretIDs,
		Domains:   req.Domains,
		UserID:    userID,
	})
	if err != nil {
		return nil, err
	}

	rsp := &GetSecretsRsp{}
	for _, v := range secretList.List {
		s := &Secret{
			SecretID: v.SecretID,
			Describe: v.Describe,
			Domain:   v.Domain,
			Key:      v.Key,
		}
		rsp.List = append(rsp.List, s)
	}
	return rsp, nil
}

type DeleteSecretReq struct {
	SecretID int64 `json:"secret_id"`
}

func DeleteSecret(ctx *gin.Context) {
	req := &DeleteSecretReq{}
	value, _ := ctx.Get(apiserver.KeyClaim)
	apiserver.Handle(ctx, req, func(tCtx context.Context) (interface{}, error) {
		return nil, deleteSecret(tCtx, req, value.(*apiserver.Claim).UserID)
	})
}

func deleteSecret(ctx context.Context, req *DeleteSecretReq, userID int64) error {
	_, err := apiserver.SecretClient.DeleteSecret(ctx, &secret.DeleteSecretReq{
		SecretID: req.SecretID,
		UserID:   userID,
	})
	return err
}

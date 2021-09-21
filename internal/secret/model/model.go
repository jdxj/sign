package model

import (
	"github.com/jdxj/sign/internal/pkg/util"
	"github.com/jdxj/sign/internal/proto/crontab"
	secretPb "github.com/jdxj/sign/internal/proto/secret"
	secretDao "github.com/jdxj/sign/internal/secret/dao/secret"
)

func CreateSecret(key string, req *secretPb.CreateSecretReq) (*secretPb.CreateSecretRsp, error) {
	sec := &secretDao.Secret{
		UserID: req.UserID,
		Domain: int32(req.Domain),
		Key:    util.Encrypt(key, req.Key),
	}
	secretID, err := secretDao.Insert(sec)
	rsp := &secretPb.CreateSecretRsp{
		SecretID: secretID,
	}
	return rsp, err
}

func GetSecret(key string, req *secretPb.GetSecretReq) (*secretPb.GetSecretRsp, error) {
	where := map[string]interface{}{
		"secret_id = ?": req.SecretID,
	}

	secret, err := secretDao.FindOne(where)
	if err != nil {
		return nil, err
	}
	rsp := &secretPb.GetSecretRsp{
		SecretID: secret.SecretID,
		Domain:   crontab.Domain(secret.Domain),
		Key:      util.Decrypt(key, secret.Key),
	}
	return rsp, nil
}

func UpdateSecret(req *secretPb.UpdateSecretReq) error {
	where := map[string]interface{}{
		"secret_id = ?": req.SecretID,
	}

	data := map[string]interface{}{}
	if req.Domain != 0 {
		data["domain"] = req.Domain
	}
	if req.Key != "" {
		data["key"] = req.Key
	}
	return secretDao.Update(where, data)
}

func DeleteSecret(req *secretPb.DeleteSecretReq) error {
	where := map[string]interface{}{
		"secret_id = ?": req.SecretID,
	}
	return secretDao.Delete(where)
}

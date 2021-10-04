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
		"secret_id": req.SecretID,
	}
	secret, err := secretDao.FindOne(where)
	if err != nil {
		return nil, err
	}

	record := &secretPb.SecretRecord{
		SecretID: secret.SecretID,
		Describe: secret.Describe,
		Domain:   crontab.Domain(secret.Domain),
		Key:      util.Decrypt(key, secret.Key),
	}
	rsp := &secretPb.GetSecretRsp{Record: record}
	return rsp, nil
}

func GetSecretList(key string, req *secretPb.GetSecretListReq) (*secretPb.GetSecretListRsp, error) {
	where := map[string]interface{}{
		"user_id = ?": req.UserID,
	}
	secrets, err := secretDao.Find(where)
	if err != nil {
		return nil, err
	}

	rsp := &secretPb.GetSecretListRsp{}
	for _, v := range secrets {
		record := &secretPb.SecretRecord{
			SecretID: v.SecretID,
			Describe: v.Describe,
			Domain:   crontab.Domain(v.Domain),
			Key:      util.Decrypt(key, v.Key),
		}
		rsp.List = append(rsp.List, record)
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

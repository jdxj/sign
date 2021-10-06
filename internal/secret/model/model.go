package model

import (
	"errors"
	"fmt"

	"github.com/jdxj/sign/internal/pkg/util"
	"github.com/jdxj/sign/internal/proto/crontab"
	secretPb "github.com/jdxj/sign/internal/proto/secret"
	secretDao "github.com/jdxj/sign/internal/secret/dao/secret"
)

var (
	ErrDomainNotFound = errors.New("domain not found")
)

func CreateSecret(key string, req *secretPb.CreateSecretReq) (*secretPb.CreateSecretRsp, error) {
	if _, ok := crontab.Domain_name[int32(req.Domain)]; !ok {
		return nil, fmt.Errorf("%w: %d", ErrDomainNotFound, req.Domain)
	}

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

func GetSecretList(key string, req *secretPb.GetSecretListReq) (*secretPb.GetSecretListRsp, error) {
	where := map[string]interface{}{
		"user_id = ?": req.UserID,
	}
	if len(req.SecretIDs) != 0 {
		where["secret_id IN ?"] = req.SecretIDs
	}
	if len(req.Domains) != 0 {
		where["domain IN ?"] = req.Domains
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
	if _, ok := crontab.Domain_name[int32(req.Domain)]; !ok {
		return fmt.Errorf("%w: %d", ErrDomainNotFound, req.Domain)
	}

	where := map[string]interface{}{
		"secret_id = ?": req.SecretID,
	}

	data := map[string]interface{}{
		"describe": req.Describe,
		"domain":   req.Domain,
		"key":      req.Key,
	}
	return secretDao.Update(where, data)
}

func DeleteSecret(req *secretPb.DeleteSecretReq) error {
	where := map[string]interface{}{
		"secret_id = ?": req.SecretID,
		"user_id = ?":   req.UserID,
	}
	return secretDao.Delete(where)
}

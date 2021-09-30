package secret

import (
	"github.com/jdxj/sign/internal/pkg/db"
)

// todo: 添加 describe 字段

type Secret struct {
	SecretID int64 `gorm:"primaryKey"`
	UserID   int64
	Domain   int32
	Key      string
}

const TableName = "secret"

func Insert(sec *Secret) (int64, error) {
	query := db.Gorm.Table(TableName)
	return sec.SecretID, query.Create(sec).Error
}

func FindOne(where map[string]interface{}) (Secret, error) {
	query := db.Gorm.Table(TableName)
	for cond, param := range where {
		query = query.Where(cond, param)
	}

	s := Secret{}
	return s, query.First(&s).Error
}

func Find(where map[string]interface{}) ([]Secret, error) {
	query := db.Gorm.Table(TableName)
	for cond, param := range where {
		query = query.Where(cond, param)
	}

	var rows []Secret
	return rows, query.Find(&rows).Error
}

func Update(where, data map[string]interface{}) error {
	query := db.Gorm.Table(TableName)
	for cond, param := range where {
		query = query.Where(cond, param)
	}
	return query.Updates(data).Error
}

func Delete(where map[string]interface{}) error {
	query := db.Gorm.Table(TableName)
	for cond, param := range where {
		query = query.Where(cond, param)
	}
	return query.Delete(nil).Error
}

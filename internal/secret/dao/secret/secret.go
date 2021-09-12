package secret

import (
	"github.com/jdxj/sign/internal/pkg/db"
)

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

func Find(where map[string]interface{}) ([]Secret, error) {
	query := db.Gorm.Table(TableName)
	for k, v := range where {
		query = query.Where(k, v)
	}

	var rows []Secret
	return rows, query.Find(&rows).Error
}

func Update(where, data map[string]interface{}) error {
	query := db.Gorm.Table(TableName)
	for k, v := range where {
		query = query.Where(k, v)
	}
	return query.Updates(data).Error
}

func Delete(where map[string]interface{}) error {
	query := db.Gorm.Table(TableName)
	for k, v := range where {
		query = query.Where(k, v)
	}
	return query.Delete(nil).Error
}

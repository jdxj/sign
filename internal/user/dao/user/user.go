package user

import (
	"github.com/jdxj/sign/internal/pkg/db"
)

type User struct {
	UserID   int64 `gorm:"primaryKey"`
	Nickname string
	Password string
	Salt     string
}

const (
	TableName = "user"
)

func Insert(data *User) error {
	query := db.Gorm.Table(TableName)
	return query.Create(data).Error
}

func FindOne(where map[string]interface{}) (User, error) {
	query := db.Gorm.Table(TableName)
	for cond, param := range where {
		query = query.Where(cond, param)
	}

	var row User
	return row, query.First(&row).Error
}

func Find(where map[string]interface{}) ([]User, error) {
	query := db.Gorm.Table(TableName)
	for cond, param := range where {
		query = query.Where(cond, param)
	}

	var rows []User
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

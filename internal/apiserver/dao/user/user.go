package user

import (
	"github.com/jdxj/sign/internal/pkg/db"
)

var (
	TableName = "user"
)

type User struct {
	UserID   int64
	Nickname string
	Password string
}

func Find(nickname, password string) (*User, error) {
	user := &User{}
	db := db.Gorm.Table(TableName).
		Where("nickname = ? AND password = ?", nickname, password)
	return user, db.First(user).Error
}

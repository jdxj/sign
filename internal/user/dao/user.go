package dao

import (
	"time"
)

type User struct {
	UserID    int64 `gorm:"primaryKey"`
	Nickname  string
	Password  string
	Salt      string
	Mail      string
	Telegram  int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (u *User) TableName() string {
	return TableName
}

const (
	TableName = "user"
)

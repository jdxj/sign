package dao

import (
	"time"
)

type Task struct {
	TaskID    int64 `gorm:"primaryKey"`
	Describe  string
	UserID    int64
	Kind      string
	Spec      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Param     []byte
}

const (
	TableName = "task"
)

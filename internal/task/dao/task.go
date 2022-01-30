package dao

import (
	"time"
)

type Task struct {
	TaskID      int64 `gorm:"primaryKey"`
	Description string
	UserID      int64
	Kind        string
	Spec        string
	Param       []byte
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

func (t *Task) TableName() string {
	return "task"
}

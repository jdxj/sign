package task

import (
	"github.com/jdxj/sign/internal/pkg/db"
)

type Task struct {
	TaskID   int64 `gorm:"primaryKey"`
	Describe string
	UserID   int64
	SecretID int64
	Kind     int
	Spec     string
}

const (
	TableName = "task"
)

func Insert(data *Task) error {
	query := db.Gorm.Table(TableName)
	return query.Create(data).Error
}

func Find(where map[string]interface{}) ([]Task, error) {
	query := db.Gorm.Table(TableName)
	for cond, param := range where {
		query = query.Where(cond, param)
	}

	var rows []Task
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

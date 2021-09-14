package specification

import (
	"errors"

	"github.com/go-sql-driver/mysql"

	"github.com/jdxj/sign/internal/pkg/db"
)

type Specification struct {
	SpecID int64 `gorm:"primaryKey"`
	Spec   string
}

const (
	TableName = "specification"
)

func Insert(data *Specification) error {
	query := db.Gorm.Table(TableName)
	err := query.Create(data).Error
	if err == nil {
		return nil
	}
	var mysqlError *mysql.MySQLError
	if !errors.As(err, &mysqlError) || mysqlError.Number != 1062 {
		return err
	}

	row, err := FindOne(map[string]interface{}{"spec = ?": data.Spec})
	data.SpecID = row.SpecID
	return err
}

func FindOne(where map[string]interface{}) (Specification, error) {
	query := db.Gorm.Table(TableName)
	for cond, param := range where {
		query = query.Where(cond, param)
	}

	var row Specification
	return row, query.First(&row).Error
}

func Find(where map[string]interface{}) ([]Specification, error) {
	query := db.Gorm.Table(TableName)
	for cond, param := range where {
		query = query.Where(cond, param)
	}

	var rows []Specification
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

package db

import (
	"context"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/jdxj/sign/internal/pkg/config"
)

var (
	gormDB *gorm.DB
)

func InitGorm(conf config.DB) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.User, conf.Pass, conf.Host, conf.Port, conf.Dbname)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	if err := sqlDB.Ping(); err != nil {
		return err
	}
	gormDB = db
	log.Printf(" connected to db")
	return nil
}

func WithCtx(ctx context.Context) *gorm.DB {
	return gormDB.WithContext(ctx)
}

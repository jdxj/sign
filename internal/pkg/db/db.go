package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/jdxj/sign/internal/pkg/config"
)

var (
	Conn *gorm.DB
)

func InitGorm(conf config.DB) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.User, conf.Pass, conf.Host, conf.Port, conf.Dbname)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	Conn = db
}

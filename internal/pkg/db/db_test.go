package db

import (
	"testing"

	"github.com/jdxj/sign/internal/pkg/config"
)

type User struct {
	UserID   uint64
	Nickname string
	Password string
}

func TestInitGorm(t *testing.T) {
	conf := config.DB{
		User:   "root",
		Pass:   "123456",
		Host:   "127.0.0.1",
		Port:   3306,
		Dbname: "sign",
	}
	_ = InitGorm(conf)
}

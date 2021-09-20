package apiserver

import (
	"os"
	"testing"
	"time"

	"github.com/jdxj/sign/internal/pkg/config"
	"github.com/jdxj/sign/internal/pkg/db"
	"github.com/jdxj/sign/internal/pkg/logger"
)

func TestMain(t *testing.M) {
	logger.Init("./test.log")
	db.InitGorm(config.DB{
		User:   "root",
		Pass:   "123456",
		Host:   "127.0.0.1",
		Port:   3306,
		Dbname: "sign",
	})
	os.Exit(t.Run())
}

func TestStart(t *testing.T) {
	Start(config.APIServer{Port: "8080"})
	time.Sleep(time.Hour)
}
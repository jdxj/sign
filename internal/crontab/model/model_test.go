package model

import (
	"fmt"
	"os"
	"testing"

	"github.com/jdxj/sign/internal/pkg/config"
	"github.com/jdxj/sign/internal/pkg/db"
	"github.com/jdxj/sign/internal/proto/crontab"
)

func TestMain(t *testing.M) {
	db.InitGorm(config.DB{
		User:   "root",
		Pass:   "123456",
		Host:   "127.0.0.1",
		Port:   3306,
		Dbname: "sign",
	})
	os.Exit(t.Run())
}

func TestGetTasks(t *testing.T) {
	rsp, err := GetTasks(&crontab.GetTasksReq{UserID: 1})
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	for _, v := range rsp.List {
		fmt.Printf("%+v\n", v)
	}
}

package task

import (
	"fmt"
	"os"
	"testing"

	"github.com/jdxj/sign/internal/pkg/config"
	"github.com/jdxj/sign/internal/pkg/db"
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

func TestFind(t *testing.T) {
	where := map[string]interface{}{
		"user_id = ?": 1,
	}
	rows, err := Find(where)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	for _, row := range rows {
		fmt.Printf("%+v\n", row)
	}
}

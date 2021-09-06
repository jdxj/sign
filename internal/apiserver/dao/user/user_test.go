package user

import (
	"fmt"
	"os"
	"testing"

	"github.com/jdxj/sign/internal/pkg/config"
	"github.com/jdxj/sign/internal/pkg/db"
)

func TestMain(t *testing.M) {
	conf := config.DB{
		User:   "root",
		Pass:   "123456",
		Host:   "127.0.0.1",
		Port:   3306,
		Dbname: "sign",
	}
	err := db.InitGorm(conf)
	if err != nil {
		panic(err)
	}
	os.Exit(t.Run())
}

func TestFind(t *testing.T) {
	name := "jdxj"
	pass := "jdxj1"
	user, err := Find(name, pass)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", *user)
}

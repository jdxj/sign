package specification

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

func TestInsert(t *testing.T) {
	data := &Specification{
		Spec: "abc",
	}
	err := Insert(data)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%d\n", data.SpecID)
	fmt.Printf("%+v\n", data)
}

func TestFind(t *testing.T) {
	rsp, err := Find(nil)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	for _, v := range rsp {
		fmt.Printf("%+v\n", v)
	}
}

package secret

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
	for _, v := range rows {
		fmt.Printf("%+v\n", v)
	}

}

func TestInsert(t *testing.T) {
	sec := &Secret{
		SecretID: 0,
		UserID:   1,
		Domain:   101,
		Key:      "def",
	}
	id, err := Insert(sec)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("secret id: %d\n", id)
}

func TestUpdate(t *testing.T) {
	where := map[string]interface{}{
		"secret_id = ?": 2,
	}
	data := map[string]interface{}{}
	err := Update(where, data)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
}

func TestDelete(t *testing.T) {
	where := map[string]interface{}{
		"secret_id = ?": 1,
	}
	err := Delete(where)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
}

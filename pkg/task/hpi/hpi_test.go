package hpi

import (
	"fmt"
	"testing"
	"time"

	"github.com/jdxj/sign/pkg/task/common"
)

func TestMSecond(t *testing.T) {
	fmt.Println(time.Now().UnixNano())
	fmt.Println(1627303261118)
}

var tmp = "symphony=ac"

func TestAuth(t *testing.T) {
	client, err := Auth(tmp)
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}

	task := &common.Task{
		ID:     "test",
		Type:   202,
		Client: client,
	}
	SignIn(task)
}

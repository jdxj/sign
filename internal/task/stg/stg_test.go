package stg

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/jdxj/sign/internal/task/common"
)

func TestBalance(t *testing.T) {
	d, err := common.ParseRawBody(http.DefaultClient, VerifyURL)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%s\n", d)
}

var (
	tmp = ""
)

func TestAuth(t *testing.T) {
	c, err := Auth("a=b")
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	_ = c
}

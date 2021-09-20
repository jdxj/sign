package v2ex

import (
	"testing"
)

var (
	tmp = ``
)

func TestAuth(t *testing.T) {
	client, err := Auth(tmp)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	err = SignIn(client)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
}

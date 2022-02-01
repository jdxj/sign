package stg

import (
	"fmt"
	"testing"
)

func TestSignIn_Execute(t *testing.T) {
	si := &SignIn{}
	msg, err := si.Execute(nil)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("msg: %s\n", msg)
}

func TestAuth(t *testing.T) {
	_, err := auth("", "")
	if err != nil {
		t.Fatalf("%s\n", err)
	}
}

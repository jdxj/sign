package bilibili

import (
	"fmt"
	"testing"
)

func TestSignIn_Execute(t *testing.T) {
	si := &SignIn{}

	msg, err := si.Execute(key)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("msg: %s\n", msg)
}

func TestBi_Execute(t *testing.T) {
	bi := &Bi{}
	msg, err := bi.Execute(key)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("msg: %s\n", msg)
}

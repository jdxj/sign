package hpi

import (
	"fmt"
	"regexp"
	"testing"
)

func TestFindSignToken(t *testing.T) {
	str := ``
	reg := regexp.MustCompile(`csrfToken: '(.+)'`)
	res := reg.FindStringSubmatch(str)
	fmt.Println(res)
}

func TestRegDate(t *testing.T) {
	str := ``
	reg := regexp.MustCompile(str)
	res := reg.FindAllString(str, -1)
	for _, d := range res {
		fmt.Println(d)
	}
}

func TestSignIn_Execute(t *testing.T) {
	si := &SignIn{}
	msg, err := si.Execute(key)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("msg: %s\n", msg)
}

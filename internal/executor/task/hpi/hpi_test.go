package hpi

import (
	"encoding/json"
	"fmt"
	"regexp"
	"testing"

	"github.com/jdxj/sign/internal/pkg/util"
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
	msg, err := si.Execute(loginKey)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("msg: %s\n", msg)
}

type A struct {
	NameAbc      string
	def          string
	NameAbc2_Abc string
}

func TestJsonEncoding(t *testing.T) {
	a := A{
		NameAbc:      "123",
		def:          "456",
		NameAbc2_Abc: "abc",
	}
	d, err := json.Marshal(a)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%s\n", d)

	key := "userName=abc;userPassword=def"
	req := &loginReq{}
	err = util.PopulateStruct(key, req)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", req)
}

func TestLogin(t *testing.T) {
	token, err := login(loginKey)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("token: %s\n", token)
}

package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"path"
	"testing"
)

func TestPassword(t *testing.T) {
	for i := 0; i < 100; i++ {
		s := WithSalt("abc", "def")
		fmt.Println(s)
	}
}

func TestSalt(t *testing.T) {
	for i := 0; i < 100; i++ {
		s := Salt()
		fmt.Println(s)
	}
}

func TestLen(t *testing.T) {
	str := "abc"
	str2 := "你好"
	fmt.Printf("%s: strLen: %d, byteLen: %d\n", str, len(str), len([]byte(str)))
	fmt.Printf("%s: strLen: %d, byteLen: %d\n", str2, len(str2), len([]byte(str2)))
}

func TestEncrypt(t *testing.T) {
	key := "abc"
	text := "def"
	res := Encrypt([]byte(key), []byte(text))
	plaintext := Decrypt([]byte(key), res)
	if !bytes.Equal([]byte(text), plaintext) {
		t.Fatalf("err encrypt or decrypt")
	}
}

func TestValues(t *testing.T) {
	v := url.Values{}
	res := v.Encode()
	fmt.Printf("res: %s\n", res)
	fmt.Printf("len: %d\n", len(v))
}

func TestParseQuery(t *testing.T) {
	u := "https://example.com?abc=123&ghi=789"
	uu, err := url.Parse(u)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("raw query: %s\n", uu.RawQuery)

	vs := uu.Query()
	for k, v := range vs {
		fmt.Printf("k: %s, v: %v\n", k, v)
	}

	vs.Add("def", "456")
	uu.RawQuery = vs.Encode()
	fmt.Printf("res: %s\n", uu.String())
}

func TestJoin(t *testing.T) {
	u := "https://example.com?abc=123&ghi=789"
	uu, err := url.Parse(u)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	uu.Path = path.Join(uu.Path, "")
	fmt.Printf("res: %s\n", uu.String())
}

func TestMarshal(t *testing.T) {
	d, err := json.Marshal(nil)
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	fmt.Printf("d: %v\n", d)
}

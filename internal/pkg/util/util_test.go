package util

import (
	"fmt"
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
	res := Encrypt(key, "def")
	plaintext := Decrypt(key, res)
	if text != plaintext {
		t.Fatalf("err encrypt or decrypt")
	}
}

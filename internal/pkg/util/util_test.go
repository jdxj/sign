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

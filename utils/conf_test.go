package utils

import (
	"fmt"
	"testing"
)

func TestConfAll(t *testing.T) {
	kvs := ConfAll("bilibili.com")
	for _, kv := range kvs {
		fmt.Println("key:", kv.K)
		fmt.Println("value:", kv.V)
	}
}

func TestEmptyFunc(t *testing.T) {
	EmptyFunc()
}
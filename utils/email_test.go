package utils

import (
	"fmt"
	"testing"
)

func TestSendEmail(t *testing.T) {
	err := SendEmail("test", "hello")
	if err != nil {
		fmt.Println(err)
		return
	}
}

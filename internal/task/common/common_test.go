package common

import (
	"testing"
)

func TestVerifyDate(t *testing.T) {
	err := VerifyDate("2021-07-27")
	if err != nil {
		t.Fatalf("%s\n", err)
	}
}

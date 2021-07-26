package toucher

import (
	"fmt"
	"testing"
	"time"
)

var key1 = ""
var key2 = ""

func TestYearDay(t *testing.T) {
	now := time.Now()
	fmt.Printf("now: %d\n", now.YearDay())

	befor := now.Add(-24 * time.Hour)
	fmt.Printf("befor: %d\n", befor.YearDay())
}

func TestValidatorBili(t *testing.T) {
	bili, err := NewBili("abc", key1)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	err = bili.SignIn()
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	err = bili.Verify()
	if err != nil {
		t.Fatalf("%s\n", err)
	}
}

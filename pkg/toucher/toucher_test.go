package toucher

import (
	"fmt"
	"testing"
	"time"
)

var key1 = ""
var key2 = ""

func TestParseCookies(t *testing.T) {
	cookies := ParseCookies(key2, ".bilibili.com")
	for _, v := range cookies {
		fmt.Printf("%#v\n", v)
	}
}

func TestBilibili_Auth(t *testing.T) {
	bili := NewBilibili()
	jar, err := bili.Auth(key1)
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}
	fmt.Printf("%#v\n", jar)

	jar, err = bili.Auth(key2)
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}

	fmt.Printf("%#v\n", jar)
}

func TestYearDay(t *testing.T) {
	now := time.Now()
	fmt.Printf("now: %d\n", now.YearDay())

	befor := now.Add(-24 * time.Hour)
	fmt.Printf("befor: %d\n", befor.YearDay())
}

func TestBilibili_Verify(t *testing.T) {
	bili := NewBilibili()
	jar, err := bili.Auth(key1)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("--\n")
	err = bili.Verify(jar)
	if err != nil {
		t.Fatal(err)
	}
}

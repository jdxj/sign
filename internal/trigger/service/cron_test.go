package service

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestRandomSchedule_Next(t *testing.T) {
	rs, err := newRandomSchedule("")
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	for i := 0; i < 10; i++ {
		tt := rs.Next(time.Now())
		fmt.Printf("%s\n", tt)
	}

	fmt.Println()
	rs, err = newRandomSchedule("@randomly 9:00:00 21:00:01")
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	for i := 0; i < 10; i++ {
		tt := rs.Next(time.Now())
		fmt.Printf("%s\n", tt)
	}
}

func TestSplit(t *testing.T) {
	res := strings.Split("", " ")
	fmt.Printf("len: %d\n", len(res))
	for _, v := range res {
		fmt.Printf("%s\n", v)
	}
}

func TestSub(t *testing.T) {
	t1 := time.Time{}
	t2 := time.Time{}
	fmt.Printf("%f\n", t1.Sub(t2).Seconds())

	fmt.Printf("h: %d\n", t1.Hour())
	fmt.Printf("m: %d\n", t1.Minute())
	fmt.Printf("s: %d\n", t1.Second())
	fmt.Printf("t3: %s\n", t1.Add(24*time.Hour))
	fmt.Printf("t4: %s\n", t1.Add(24*time.Hour-time.Second))
	fmt.Printf("t5: %s\n", t1.In(time.Local))
}

func TestParseRanges(t *testing.T) {
	ts := []string{
		//"20:60",
		"20:59:00",
		"8:59:00",
	}
	res, err := parseRanges(ts)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%v\n", res)
}

func TestRandomInt(t *testing.T) {
	zero, err := time.ParseInLocation(randomLayout, "00:00:00", time.UTC)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%s\n", zero)
}

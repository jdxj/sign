package service

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestRandomSchedule_Next(t *testing.T) {
	rs := &randomSchedule{}
	for i := 0; i < 100; i++ {
		fmt.Printf("%s\n", rs.Next(time.Now()))
	}
}

func TestSplit(t *testing.T) {
	res := strings.Split("", " ")
	fmt.Printf("len: %d\n", len(res))
	for _, v := range res {
		fmt.Printf("%s\n", v)
	}
}

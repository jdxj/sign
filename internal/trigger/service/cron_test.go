package service

import (
	"fmt"
	"testing"
	"time"
)

func TestRandomSchedule_Next(t *testing.T) {
	rs := &randomSchedule{}
	for i := 0; i < 100; i++ {
		fmt.Printf("%s\n", rs.Next(time.Now()))
	}
}

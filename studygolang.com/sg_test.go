package studygolang_com

import (
	"fmt"
	"testing"
	"time"
)

func TestTiming(t *testing.T) {
	timer := time.NewTimer(time.Second)
	var i int
	for {
		select {
		case <-timer.C:
			fmt.Println("time arrived!", i, time.Now().Unix())
			i++
			timer.Reset(time.Second)
		}

		if i == 10 {
			timer.Reset(2 * time.Second)
		}
	}
}

func TestNewDay(t *testing.T) {
	now := time.Now()
	fmt.Println(now)

	tomNow := now.Add(24 * time.Hour)
	newDay := time.Date(tomNow.Year(), tomNow.Month(), tomNow.Day(), 0, 0, 0, 0, tomNow.Location())
	fmt.Println(newDay)
}

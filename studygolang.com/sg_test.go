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

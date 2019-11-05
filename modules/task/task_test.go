package task

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	ticker := time.NewTicker(time.Minute)
	//timer := time.NewTimer(time.Minute)

	for {
		f := func() {
			fmt.Println("模拟执行任务:", rand.Int())
		}
		dur := rand.Int63n(int64(time.Minute))
		time.AfterFunc(time.Duration(dur), f)
		<-ticker.C
	}
}

func TestDur(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	dur := time.Duration(r.Int63n(86400) * 1e9)
	fmt.Println(dur)
}

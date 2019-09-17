package utils

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMyTimer(t *testing.T) {
	now := time.Now()
	new := now.Add(time.Minute)
	fmt.Println(now.Sub(new))
	fmt.Println(new.Sub(now))
}

func Test0Timer(t *testing.T) {
	timer := time.NewTimer(0)
	timer.Reset(10 * time.Second)

	fmt.Println(time.Now())
	select {
	case <-timer.C:
		fmt.Println("ok")
	}
	fmt.Println(time.Now())
}

func TestWaitGroup(t *testing.T) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		fmt.Println("kkk")
	}()
	wg.Wait()
}

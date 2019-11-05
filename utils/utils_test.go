package utils

import (
	"fmt"
	"strings"
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

func TestTrim(t *testing.T) {
	res := strings.ReplaceAll("123=456; 789=101", " ", "")
	fmt.Println(res)
}

func TestStrToCookies(t *testing.T) {
	//cooStr := "123=456; 789=101112"
	//cooStr := "123=456; =101112"
	cooStr := "=456"
	cookies, err := StrToCookies(cooStr, Pic58CookieDomain)
	if err != nil {
		panic(err)
	}

	for _, cookie := range cookies {
		fmt.Println(cookie)
	}
}

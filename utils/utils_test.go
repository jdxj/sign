package utils

import (
	"encoding/base64"
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
	// 带有引号的 cookie
	//cooStr := ``
	// 正常 cookie
	cooStr := ``
	cookies, err := StrToCookies(cooStr, V2exCookieDomain)
	if err != nil {
		panic(err)
	}

	for _, cookie := range cookies {
		fmt.Println(cookie)
	}
}

func TestBase64(t *testing.T) {
	raw := "X3V1aWQ9QzUyRTUyQkUtNUQwMi1EODk5LUI4QTEtNkVBNTlDNDhFNjg1Mzc2MjFpbmZvYzsgYnV2aWQzPTc2MEU3NUE3LTNCRTUtNDY2Qy05RDM2LUQxMDVERURBNTUwRDE5MDk0NmluZm9jOyBMSVZFX0JVVklEPUFVVE80NTE1NzEzMTA4Mzg1MzM5OyBzaWQ9ODltbjFoZWY7IERlZGVVc2VySUQ9OTg2MzQyMTE7IERlZGVVc2VySURfX2NrTWQ1PTcwY2I1NDc2ZWMzZjA5Nzc7IFNFU1NEQVRBPWFhYjljYWQxJTJDMTU3MzkwMjg1MCUyQzA2ODM1YmExOyBiaWxpX2pjdD1kNmM2YWJjZTM0ZmNmNmM2YmNkNGJlOTNjZTBkYzQ1NTsgYnBfdF9vZmZzZXRfOTg2MzQyMTE9MzE4MTQ1Nzc2ODU0NDc0Mjg5Cg=="
	data, err := base64.StdEncoding.DecodeString(raw)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", data)
}

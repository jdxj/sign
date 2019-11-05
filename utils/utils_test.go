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
	//cooStr := "123=456; 789=101112"
	//cooStr := "123=456; =101112"
	cooStr := "_uuid=C52E52BE-5D02-D899-B8A1-6EA59C48E68537621infoc; buvid3=760E75A7-3BE5-466C-9D36-D105DEDA550D190946infoc; LIVE_BUVID=AUTO4515713108385339; sid=89mn1hef; DedeUserID=98634211; DedeUserID__ckMd5=70cb5476ec3f0977; SESSDATA=aab9cad1%2C1573902850%2C06835ba1; bili_jct=d6c6abce34fcf6c6bcd4be93ce0dc455; bp_t_offset_98634211=318145776854474289"
	cookies, err := StrToCookies(cooStr, Pic58CookieDomain)
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

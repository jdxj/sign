package studygolang

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"testing"
	"time"
)

func newTouchStudyGolang() (*ToucherStudyGolang, error) {
	t := &ToucherStudyGolang{
		username:  "",
		password:  "",
		loginURL:  "https://studygolang.com/account/login",
		signURL:   "https://studygolang.com/mission/daily/redeem",
		verifyKey: ".balance_area",
		signKey:   ".c9",
		signValue: "每日登录奖励已领取",
		client:    &http.Client{},
		activeURL: "https://studygolang.com/user/jdxj",
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	t.client.Jar = jar
	return t, nil
}

func TestNewTouchStudyGolang(t *testing.T) {
	tou, _ := newTouchStudyGolang()

	if !tou.Boot() {
		fmt.Println("boot fail")
		return
	}
	if !tou.Login() {
		fmt.Println("login fail")
		return
	}
	if !tou.Sign() {
		fmt.Println("sign fail")
	} else {
		fmt.Println("sign success")
	}
}

func TestTimer(t *testing.T) {
	now := time.Now()
	today0AM := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	fmt.Printf("today0AM: %s\n", today0AM.Format(time.RFC1123))

	today21PM := today0AM.Add(21 * time.Hour)
	fmt.Printf("today21AM: %s\n", today21PM.Format(time.RFC1123))

	dur := today21PM.Sub(now)
	fmt.Println(dur)
}

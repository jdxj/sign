package studygolang

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"testing"
)

func newTouchStudyGolang() (*TouchStudyGolang, error) {
	t := &TouchStudyGolang{
		username:    "",
		password:    "",
		loginURL:    "https://studygolang.com/account/login",
		signURL:     "https://studygolang.com/mission/daily/redeem",
		verifyKey:   ".balance_area",
		verifyValue: "",
		signKey:     ".c9",
		signValue:   "每日登录奖励已领取",
		client:      &http.Client{},
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

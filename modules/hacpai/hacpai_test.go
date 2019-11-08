package hacpai

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"testing"
)

func newToucherHacPai() (*ToucherHacPai, error) {
	tou := &ToucherHacPai{
		username:   "985759262@qq.com",
		password:   "",
		loginURL:   "https://hacpai.com/api/v2/login",
		signRefURL: "https://hacpai.com/activity/checkin",
		signURL:    "https://hacpai.com/activity/daily-checkin",
		client:     &http.Client{},
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	tou.client.Jar = jar
	return tou, nil
}

func TestToucherHacPai(t *testing.T) {
	tou, err := newToucherHacPai()
	if err != nil {
		panic(err)
	}

	if !tou.Boot() {
		fmt.Println("boot fail")
	}
	if !tou.Login() {
		fmt.Println("login fail")
	}
	if !tou.Sign() {
		fmt.Println("sign fail")
	}
}

func TestToMd5(t *testing.T) {
	fmt.Println(toMd5(""))
}

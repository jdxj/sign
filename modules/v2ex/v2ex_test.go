package v2ex

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"testing"
)

// https://www.v2ex.com/mission/daily
func newToucherV2ex() (*ToucherV2ex, error) {
	tou := &ToucherV2ex{
		name:        "",
		cookies:     "",
		loginURL:    "https://www.v2ex.com/member/jdxj",
		signURL:     "https://www.v2ex.com/mission/daily",
		verifyKey:   "h1",
		verifyValue: "jdxj",
		client:      &http.Client{},
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	tou.client.Jar = jar
	return tou, nil
}

func TestToucherV2ex(t *testing.T) {
	tou, err := newToucherV2ex()
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

func TestParseOnce(t *testing.T) {
	res, err := parseOnce("location.href = '/mission/daily/redeem?once=53192';")
	if err != nil {
		panic(err)
	}

	fmt.Println("res:", res)
}

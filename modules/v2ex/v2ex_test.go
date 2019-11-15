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
		name:      "",
		cookies:   ``,
		loginURL:  "https://www.v2ex.com/balance",
		signURL:   "https://www.v2ex.com/mission/daily",
		verifyKey: ".balance_area,.bigger",
		client:    &http.Client{},
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
		return
	}
	if !tou.Login() {
		fmt.Println("login fail")
		return
	}
	if !tou.Sign() {
		fmt.Println("sign fail")
		return
	}
}

func TestParseOnce(t *testing.T) {
	res, err := parseOnce("location.href = '/mission/daily/redeem?once=53192';")
	if err != nil {
		panic(err)
	}

	fmt.Println("res:", res)
}

package bilibili

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"testing"
)

func newTouchBilibili() (*ToucherBilibili, error) {
	t := &ToucherBilibili{
		cookies:  "",
		loginURL: "https://api.bilibili.com/x/web-interface/nav/stat",
		//verifyKey:   "#big",
		verifyValue: "9",
		client:      &http.Client{},
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	t.client.Jar = jar
	return t, nil
}

func TestNewTouchBilibili(t *testing.T) {
	tou, _ := newTouchBilibili()

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

package bilibili

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"testing"
)

func newTouchBilibili() (*TouchBilibili, error) {
	t := &TouchBilibili{
		cookies:     "",
		loginURL:    "https://space.bilibili.com/98634211",
		verifyKey:   "title",
		verifyValue: "王者王尼玛的个人空间 - 哔哩哔哩 ( ゜- ゜)つロ 乾杯~ Bilibili",
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
}

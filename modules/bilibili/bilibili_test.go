package bilibili

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"testing"
)

func newTouchBilibili() (*ToucherBilibili, error) {
	t := &ToucherBilibili{
		cookies:  "_uuid=4CD3F1DB-697E-8F84-E407-9FA5FF1597F966217infoc; buvid3=1CF70DC8-4872-4BDD-B0DD-DDAB5B1C7D06155831infoc; LIVE_BUVID=AUTO8615730880685539; sid=il3fqh8q; INTVER=1; DedeUserID=98634211; DedeUserID__ckMd5=70cb5476ec3f0977; SESSDATA=d114ab6d%2C1576305092%2Cf6c81ab1; bili_jct=302e04a97651abcbf3380c511697c118",
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

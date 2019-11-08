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
		cookies:     `__cfduid=d86b54b1cc6be6976fa982ed32bac8f661573210702&PB3_SESSION=2|1:0|10:1573210702|11:PB3_SESSION|36:djJleDo2OC40LjIwMy4xMzE6NTYwOTIwNTk=|f44f759f19bd44a9ec873cecbbafadde3b7bb4034de559164ad35f6000699a65&V2EX_LANG=zhcn&_ga=GA1.2.378922632.1573210728&_gid=GA1.2.1769837915.1573210728&A2=2|1:0|10:1573211699|2:A2|48:M2Q5Y2VhZmEtNmIzZi00YjI0LWFmYmQtZDk3OWRkY2VhNDY0|4c6f74cd03036d51264d15319fd1f141a43535add6c28044e08067f4fed1dc57&V2EX_TAB=2|1:0|10:1573212567|8:V2EX_TAB|8:dGVjaA==|e08ca2dacbd2cb28ace115418edae66043ca9e238f798c48110ac3126e38f7ad`,
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
	//location.href = '/mission/daily/redeem?once=69089';
	if !tou.Sign() {
		fmt.Println("sign fail")
	}
}

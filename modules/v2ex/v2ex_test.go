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
		cookies:   `__cfduid=d86b54b1cc6be6976fa982ed32bac8f661573210702; _ga=GA1.2.378922632.1573210728; A2="2|1:0|10:1573211699|2:A2|48:M2Q5Y2VhZmEtNmIzZi00YjI0LWFmYmQtZDk3OWRkY2VhNDY0|4c6f74cd03036d51264d15319fd1f141a43535add6c28044e08067f4fed1dc57"; _gid=GA1.2.1940533009.1573608528; V2EX_REFERRER="2|1:0|10:1573629223|13:V2EX_REFERRER|8:amRoYW8=|e439d473513d39fdc0d3afb3879af01421ba22cd0542ee27018dc2553ec202ab"; PB3_SESSION="2|1:0|10:1573700825|11:PB3_SESSION|36:djJleDo2OC40LjIwMy4xMzE6MzUwNTE5MzA=|2510d668dd8221034234a5009e5698021dbf7ac3f681e4424c2f69bd81a20ed2"; V2EX_LANG=zhcn; V2EX_TAB="2|1:0|10:1573717300|8:V2EX_TAB|8:dGVjaA==|34b0195cf5999173c3228a9210fd1ce7c6950496f53c3fae5912c8ef05fd3a60"; _gat=1`,
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

package stg

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/jdxj/sign/internal/task/common"
)

func TestBalance(t *testing.T) {
	d, err := common.ParseRawBody(http.DefaultClient, VerifyURL)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%s\n", d)
}

var (
	tmp = ""
)

func TestAuth(t *testing.T) {
	c, err := Auth(tmp)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	URL, _ := url.Parse(HomeURL)
	req, err := http.NewRequest("", "", nil)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	cookies := c.Jar.Cookies(URL)
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}
	res := req.Header.Get("Cookie")
	fmt.Println(res)
}

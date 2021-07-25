package toucher

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

const (
	DomainBili = ".bilibili.com"
	urlBili    = "https://www.bilibili.com/"
	authBili   = "https://api.bilibili.com/x/member/web/account"
	verifyBili = "https://api.bilibili.com/x/member/web/coin/log?jsonp=jsonp"
)

// 验证过程的错误
var (
	ErrorAuthFailed   = errors.New("auth failed")
	ErrorLogNotFound  = errors.New("log not found")
	ErrorSignInFailed = errors.New("sign in failed")
)

// 其他错误

var (
	ErrorUnsupportedDomain = errors.New("unsupported domain")
)

// 签到通用流程:
//     1. 身份验证
//     2. 执行签到
//     3. 验证签到

type Toucher interface {
	Domain() string                              // 返回所签到的网站
	Auth(cookies string) (*cookiejar.Jar, error) // 身份验证
	SignIn(*cookiejar.Jar) error                 // 执行签到
	Verify(jar *cookiejar.Jar) error             // 验证签到
}

func ParseCookies(raw, domain string) []*http.Cookie {
	req, _ := http.NewRequest("", "", nil)
	req.Header.Add("Cookie", raw)
	cookies := req.Cookies()
	for _, cookie := range cookies {
		cookie.Path = "/"
		cookie.Domain = domain
		cookie.Expires = time.Now().Add(time.Hour * 24 * 365)
	}
	return cookies
}

// Empty 用于不想解析具体数据时使用
type Empty struct{}

func ParseBody(client *http.Client, u string, v interface{}) error {
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return err
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, v)
}

func NewJar(key, domain, u string) *cookiejar.Jar {
	cookies := ParseCookies(key, domain)
	URL, _ := url.Parse(u)
	jar, _ := cookiejar.New(nil)
	jar.SetCookies(URL, cookies)
	return jar
}

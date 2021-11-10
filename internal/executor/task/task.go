package task

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"

	"github.com/mitchellh/mapstructure"
)

var (
	ErrSignInFailed = errors.New("sign in failed")
)

// ResolveCookies 从字符构造 http.Cookie
func ResolveCookies(raw, domain string) []*http.Cookie {
	req, _ := http.NewRequestWithContext(context.Background(), "", "", nil)
	req.Header.Add("Cookie", raw)
	cookies := req.Cookies()
	for _, cookie := range cookies {
		cookie.Path = "/"
		cookie.Domain = domain
		cookie.Expires = time.Now().Add(time.Hour * 24 * 365)
	}
	return cookies
}

func NewRequest(method, url string, body io.Reader) (*http.Request, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		cancel()
		return nil, nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36")
	return req, cancel, nil
}

// ParseBody 解析 body 中的 json 数据到 v
func ParseBody(client *http.Client, u string, v interface{}) error {
	req, cancel, err := NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return err
	}
	defer cancel()

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if v == nil {
		return nil
	}
	return json.Unmarshal(body, v)
}

// ParseBodyHeader 用于访问一个 url, 并且可以指定 http.Request header
func ParseBodyHeader(client *http.Client, u string, header map[string]string) error {
	req, cancel, err := NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return err
	}
	defer cancel()

	for k, v := range header {
		req.Header.Add(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	return err
}

// ParseBodyPost 可以指定 http.Request body
func ParseBodyPost(client *http.Client, u string, reader io.Reader, v interface{}) error {
	req, cancel, err := NewRequest(http.MethodPost, u, reader)
	if err != nil {
		return err
	}
	defer cancel()

	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	d, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(d, v)
}

// ParseRawBody 读取 body 为 []byte
func ParseRawBody(client *http.Client, u string) ([]byte, error) {
	req, cancel, err := NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	defer cancel()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// PostForm 使用带 context 方式发送请求, 并返回 client 和 response.
// 返回 client 的意图是其中保存了 cookie, 以便后续发送请求可重用该 cookie.
func PostForm(url string, f url.Values) (*http.Client, *http.Response, error) {
	b := strings.NewReader(f.Encode())
	req, cancel, err := NewRequest(http.MethodPost, url, b)
	if err != nil {
		return nil, nil, err
	}
	defer cancel()

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	j, _ := cookiejar.New(nil)
	c := &http.Client{Jar: j}
	rsp, err := c.Do(req)
	return c, rsp, err
}

// NewJar 从给定 cookie 字符串构造 jar
func NewJar(key, domain, u string) *cookiejar.Jar {
	cookies := ResolveCookies(key, domain)
	URL, _ := url.Parse(u)
	jar, _ := cookiejar.New(nil)
	jar.SetCookies(URL, cookies)
	return jar
}

// VerifyDate 验证指定时间字符串与当前时间的日期是否相同
func VerifyDate(raw string) error {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return err
	}
	now := time.Now().In(loc)
	last, err := time.ParseInLocation("2006-01-02", raw, loc)
	if err != nil {
		return err
	}

	if now.YearDay() != last.YearDay() {
		return ErrSignInFailed
	}
	return nil
}

// ConvertStringToMap 将 'key1=value1;key2=value2' 转换成 map
func ConvertStringToMap(key string) map[string]string {
	res := make(map[string]string)
	header := map[string][]string{
		"Cookie": {key},
	}
	req := &http.Request{Header: header}
	for _, cookie := range req.Cookies() {
		res[cookie.Name] = cookie.Value
	}
	return res
}

// PopulateStruct 从 key 中解码数据到 s
func PopulateStruct(key string, s interface{}) error {
	data := ConvertStringToMap(key)
	return mapstructure.Decode(data, s)
}

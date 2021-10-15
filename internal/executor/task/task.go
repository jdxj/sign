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

// ParseBody 解析 body 中的 json 数据到 v
func ParseBody(client *http.Client, u string, v interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
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

	if v == nil {
		return nil
	}
	return json.Unmarshal(body, v)
}

// ParseBodyHeader 用于访问一个 url, 并且可以指定 http.Request header
func ParseBodyHeader(client *http.Client, u string, header map[string]string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return err
	}

	// 通用 User-Agent
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36")
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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u, reader)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.87 Safari/537.36")
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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
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

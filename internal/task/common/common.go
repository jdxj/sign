package common

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

// XxxDomain 用于 Auth

// B 站相关任务
const (
	BiliDomain = iota + 101
	BiliSign   // 签到
	BiliBCount // 获取 B 币数量
)

// 黑客派相关任务
const (
	HPIDomain = iota + 201
	HPISign   // 签到
)

// Go 语言中文网相关任务
const (
	STGDomain = iota + 301
	STGSign   // 签到
)

// v2ex 相关任务
const (
	V2exDomain = iota + 401
	V2exSign   // 签到
)

// 访问阶段定义
const (
	SignIn = "sign in"
	Verify = "verify"

	Query = "query"

	GetToken = "get token"
)

// TypeMap 任务字符串描述
var TypeMap = map[int]string{
	BiliSign:   "B站签到",
	BiliBCount: "B币查询",
	HPISign:    "黑客派签到",
	STGSign:    "Go语言中文网签到",
	V2exSign:   "v2ex签到",
}

// 验证过程的错误
var (
	ErrorAuthFailed   = errors.New("auth failed")
	ErrorLogNotFound  = errors.New("log not found")
	ErrorSignInFailed = errors.New("sign in failed")
	ErrorDateNotMatch = errors.New("date not match")
)

// 其他错误
var (
	ErrorUnsupportedDomain = errors.New("unsupported domain")
	ErrorUnsupportedType   = errors.New("unsupported type")
)

// 重试配置
var (
	RetryNumber   = 3
	RetryInterval = 100 * time.Millisecond
)

// ResolveCookies 从字符构造 http.Cookie
func ResolveCookies(raw, domain string) []*http.Cookie {
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

// ParseBody 解析 body 中的 json 数据到 v
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

	if v == nil {
		return nil
	}
	return json.Unmarshal(body, v)
}

// ParseBodyHeader 用于访问一个 url, 并且可以指定 http.Request header
func ParseBodyHeader(client *http.Client, u string, header map[string]string) error {
	req, err := http.NewRequest(http.MethodGet, u, nil)
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
	return nil
}

// ParseBodyPost 可以指定 http.Request body
func ParseBodyPost(client *http.Client, u string, reader io.Reader, v interface{}) error {
	req, err := http.NewRequest(http.MethodPost, u, reader)
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
	req, err := http.NewRequest(http.MethodGet, u, nil)
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
		return ErrorSignInFailed
	}
	return nil
}

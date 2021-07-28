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

const (
	UnknownTask = iota - 1
	BiliDomain  = iota + 100
	BiliSign
	BiliBCount

	biliEnd

	HPIDomain = iota + 300 - biliEnd
	HPISign
)

var (
	TypeMap = map[int]string{
		BiliSign:   "B站签到",
		BiliBCount: "B币查询",
		HPISign:    "黑客派签到",
	}
)

type Task struct {
	ID     string
	Type   int
	Client *http.Client
}

func NewPool() *Pool {
	return &Pool{
		num:   0,
		tasks: make(map[int]*Task),
	}
}

type Pool struct {
	num   int
	tasks map[int]*Task
}

func (p *Pool) AddTask(t *Task) {
	p.num++
	p.tasks[p.num] = t
}

func (p *Pool) DelTask(num int) {
	delete(p.tasks, num)
}

func (p *Pool) GetAll() map[int]*Task {
	return p.tasks
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
)

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

func NewJar(key, domain, u string) *cookiejar.Jar {
	cookies := ResolveCookies(key, domain)
	URL, _ := url.Parse(u)
	jar, _ := cookiejar.New(nil)
	jar.SetCookies(URL, cookies)
	return jar
}

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

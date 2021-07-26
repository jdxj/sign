package common

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
	UnknownTask = iota
	BiliSign
	BiliBCount
)

var (
	TypeMap = map[int]string{
		BiliSign: "b站签到",
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

func NewJar(key, domain, u string) *cookiejar.Jar {
	cookies := ResolveCookies(key, domain)
	URL, _ := url.Parse(u)
	jar, _ := cookiejar.New(nil)
	jar.SetCookies(URL, cookies)
	return jar
}

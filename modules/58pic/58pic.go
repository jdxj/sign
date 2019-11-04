package _8pic

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"gopkg.in/ini.v1"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"sign/modules"
	"sign/utils"
)

func New58Pic(cfgSec *ini.Section) (*Touch58pic, error) {
	if cfgSec == nil {
		return nil, fmt.Errorf("invalid cfg")
	}

	t := &Touch58pic{
		cfg: cfgSec,
	}

	if err := t.Init(); err != nil {
		return nil, err
	}
	return t, nil
}

type cfg struct {
	Cookies            string // 用 "key=value; key=value" 表示的字符串
	LoginURL           string // 用于验证是否登录成功所要抓取的网页
	VerifyKey          string // 指定要抓取得属性, 比如 class, li 等 html 标签或属性
	VerifyValue        string // 当要抓取的属性等于 VerifyValue 时, 判断为登录成功
	VerifyReverseValue string // 当要抓取的属性等于 VerifyValue 时, 判断为登录失败
	SignDataURL        string // 执行签到签获取签到数据的链接
	SignURL            string // 执行签到所要访问的链接
}

type Touch58pic struct {
	cfg    *cfg
	client *http.Client
}

func (tou *Touch58pic) Init() error {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return err
	}

	tou.client = &http.Client{
		Jar: jar,
	}
	return nil
}

// Login 58pic 的登录使用 cookie 方式
func (tou *Touch58pic) Login() ([]*http.Cookie, error) {
	cookies, err := utils.StrToCookies(tou.cfg.Cookies, utils.Pic58Cookie)
	if err != nil {
		return nil, err
	}

	cookieURL, err := url.Parse(utils.Pic58CookieURL)
	if err != nil {
		return nil, err
	}

	tou.client.Jar.SetCookies(cookieURL, cookies)
	resp, err := tou.client.Get(tou.cfg.LoginURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	mark := true
	doc.Find(".cs-ul3-li1").Each(func(i int, selection *goquery.Selection) {
		if selection.Text() == tou.cfg.VerifyReverseValue {
			mark = false
		}
	})
	if !mark {
		return nil, fmt.Errorf("login verification fail")
	}

	// todo: 官方的 jar 实现是否能合并 cookie?
	tou.client.Jar.SetCookies(cookieURL, resp.Cookies())
	return tou.client.Jar.Cookies(cookieURL), nil
}

// todo: 实现
// 在签到前需要获取些数据
func (tou *Touch58pic) Sign() bool {
	return false
}

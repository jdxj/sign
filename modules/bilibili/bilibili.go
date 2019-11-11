package bilibili

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"gopkg.in/ini.v1"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"sign/utils"
	"sign/utils/log"
)

func NewToucherBilibili(sec *ini.Section) (*ToucherBilibili, error) {
	if sec == nil {
		return nil, fmt.Errorf("invalid cfg")
	}

	t := &ToucherBilibili{
		name:        sec.Name(),
		cookies:     sec.Key("cookies").String(),
		loginURL:    sec.Key("loginURL").String(),
		verifyKey:   sec.Key("verifyKey").String(),
		verifyValue: sec.Key("verifyValue").String(),
		client:      &http.Client{},
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	t.client.Jar = jar
	return t, nil
}

type ToucherBilibili struct {
	name     string
	cookies  string
	loginURL string

	verifyKey   string
	verifyValue string

	client    *http.Client
	loginStat bool
}

func (tou *ToucherBilibili) Name() string {
	return tou.name
}

func (tou *ToucherBilibili) Boot() bool {
	cookies, err := utils.StrToCookies(tou.cookies, utils.BilibiliCookieDomain)
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_Bilibili, err)
		return false
	}

	cookieURL, err := url.Parse(utils.BilibiliCookieURL)
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_Bilibili, err)
		return false
	}

	tou.client.Jar.SetCookies(cookieURL, cookies)
	return true
}

func (tou *ToucherBilibili) Login() bool {
	req, err := http.NewRequest("GET", tou.loginURL, nil)
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_Bilibili, err)
		return false
	}

	// todo: 为所有请求生成 user-agent
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.70 Safari/537.36")
	resp, err := tou.client.Do(req)
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_Bilibili, err)
		return false
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_Bilibili, err)
		return false
	}

	var mark bool
	doc.Find(tou.verifyKey).Each(func(i int, selection *goquery.Selection) {
		log.MyLogger.Debug("%s verify value: %s", log.Log_Bilibili, selection.Text())
		if selection.Text() == tou.verifyValue {
			mark = true
		}
	})
	tou.loginStat = mark
	return mark
}

func (tou *ToucherBilibili) Sign() bool {
	_, err := tou.client.Get("http://bilibili.com")
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_Bilibili, err)
		return false
	}

	return true
}

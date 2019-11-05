package studygolang

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"gopkg.in/ini.v1"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"sign/utils"
)

func NewToucherStudyGolang(sec *ini.Section) (*ToucherStudyGolang, error) {
	if sec == nil {
		return nil, fmt.Errorf("invalid cfg")
	}

	t := &ToucherStudyGolang{
		name:        sec.Name(),
		username:    sec.Key("username").String(),
		password:    sec.Key("password").String(),
		loginURL:    sec.Key("loginURL").String(),
		signURL:     sec.Key("signURL").String(),
		verifyKey:   sec.Key("verifyKey").String(),
		verifyValue: sec.Key("verifyValue").String(),
		signKey:     sec.Key("signKey").String(),
		signValue:   sec.Key("signValue").String(),
		client:      &http.Client{},
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	t.client.Jar = jar
	return t, nil
}

type ToucherStudyGolang struct {
	name string

	username string
	password string

	loginURL    string
	signURL     string
	verifyKey   string
	verifyValue string
	signKey     string
	signValue   string

	client *http.Client

	bootStat bool
}

func (tou *ToucherStudyGolang) Name() string {
	return tou.name
}

func (tou *ToucherStudyGolang) Boot() bool {
	val := url.Values{
		"redirect_uri": []string{"https://studygolang.com/"},
		"username":     []string{tou.username},
		"passwd":       []string{tou.password},
		"remember_me":  []string{"1"},
	}
	resp, err := tou.client.PostForm(tou.loginURL, val)
	if err != nil {
		utils.MyLogger.Error("%s", err)
		return false
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		utils.MyLogger.Error("%s", err)
		return false
	}

	var mark bool
	doc.Find(tou.verifyKey).Each(func(i int, selection *goquery.Selection) {
		mark = true
	})

	tou.bootStat = mark
	return mark
}

func (tou *ToucherStudyGolang) Login() bool {
	return tou.bootStat
}

func (tou *ToucherStudyGolang) Sign() bool {
	resp, err := tou.client.Get(tou.signURL)
	if err != nil {
		utils.MyLogger.Error("%s, %s", "[StudyGolang]", err)
		return false
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		utils.MyLogger.Error("%s, %s", "[StudyGolang]", err)
		return false
	}

	var mark bool
	doc.Find(tou.signKey).Each(func(i int, selection *goquery.Selection) {
		// 只要有一个相等就判为签到成功
		if selection.Text() == tou.signValue {
			mark = true
		}
	})
	return mark
}

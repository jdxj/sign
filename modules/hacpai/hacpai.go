package hacpai

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"sign/utils"
	"sign/utils/log"

	"github.com/PuerkitoBio/goquery"

	"gopkg.in/ini.v1"
)

func NewToucherHacPai(sec *ini.Section) (*ToucherHacPai, error) {
	if sec == nil {
		return nil, fmt.Errorf("invaild section config")
	}

	tou := &ToucherHacPai{
		//username: sec.Key("username").String(),
		//password: sec.Key("password").String(),
		cookies:     sec.Key("cookies").String(),
		loginURL:    sec.Key("loginURL").String(),
		signURL:     sec.Key("signURL").String(),
		verifyKey:   sec.Key("verifyKey").String(),
		verifyValue: sec.Key("verifyValue").String(),
		client:      &http.Client{},
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	tou.client.Jar = jar
	return tou, nil
}

type ToucherHacPai struct {
	//username string
	//password string
	cookies string

	loginURL string
	signURL  string

	verifyKey   string
	verifyValue string

	client *http.Client
}

func (tou *ToucherHacPai) Name() string {
	return ""
}

func (tou *ToucherHacPai) Boot() bool {
	cookies, err := utils.StrToCookies(tou.cookies, utils.HacPaiCookieDomain)
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_HacPai, err)
		return false
	}

	cookieURL, err := url.Parse(utils.HacPaiCookieURL)
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_HacPai, err)
		return false
	}

	tou.client.Jar.SetCookies(cookieURL, cookies)
	return true
}

func (tou *ToucherHacPai) Login() bool {
	resp, err := tou.client.Get(tou.loginURL)
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_HacPai, err)
		return false
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_HacPai, err)
		return false
	}

	var mark bool
	doc.Find(tou.verifyKey).Each(func(i int, selection *goquery.Selection) {
		fmt.Println("selection:", selection.Text())
		if selection.Text() == tou.verifyValue {
			mark = true
		}
	})
	return mark
}

// todo: 实现
func (tou *ToucherHacPai) Sign() bool {
	resp, err := tou.client.Get(tou.signURL)
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_HacPai, err)
		return false
	}
	defer resp.Body.Close()

	//doc, err := goquery.NewDocumentFromReader(resp.Body)
	//if err != nil {
	//	log.MyLogger.Error("%s %s", log.Log_HacPai, err)
	//	return false
	//}
	//
	//realSignURL, ok := doc.Find(".green").Attr("href")
	//if !ok {
	//	log.MyLogger.Error("%s %s", log.Log_HacPai, "real sign url not found")
	//	return false
	//}
	//fmt.Println("real sign url:", realSignURL)
	//
	//// todo: 假设 get 方法可以访问
	//resp, err = tou.client.Get(realSignURL)
	//if err != nil {
	//	log.MyLogger.Error("%s can not access real sign url: %s, %s", log.Log_HacPai, realSignURL, err)
	//	return false
	//}
	defer resp.Body.Close()

	//var mark bool
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.MyLogger.Error("%s can not read resp body after get real sign url: %s", log.Log_HacPai, err)
		return false
	}

	fmt.Printf("%s\n", data)

	return false
}

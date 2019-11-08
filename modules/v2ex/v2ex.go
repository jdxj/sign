package v2ex

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

func NewToucherV2ex(sec *ini.Section) (*ToucherV2ex, error) {
	if sec == nil {
		return nil, fmt.Errorf("invaild section config")
	}

	tou := &ToucherV2ex{
		name:        sec.Name(),
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

type ToucherV2ex struct {
	name    string
	cookies string

	loginURL    string
	signURL     string
	verifyKey   string
	verifyValue string

	client *http.Client
}

func (tou *ToucherV2ex) Name() string {
	return tou.name
}

// todo: 实现
func (tou *ToucherV2ex) Boot() bool {
	cookies, err := utils.StrToCookies(tou.cookies, utils.V2exCookieDomain)
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_V2ex, err)
		return false
	}

	cookieURL, err := url.Parse(utils.V2exCookieURL)
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_V2ex, err)
		return false
	}

	tou.client.Jar.SetCookies(cookieURL, cookies)
	return true
}

// todo: 实现
func (tou *ToucherV2ex) Login() bool {
	resp, err := tou.client.Get(tou.loginURL)
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_V2ex, err)
		return false
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_V2ex, err)
		return false
	}

	sel := doc.Find(tou.verifyKey)
	if sel == nil {
		log.MyLogger.Error("%s selection is nil", log.Log_V2ex)
		return false
	}
	if sel.Text() != tou.verifyValue {
		log.MyLogger.Error("%s user name not found", log.Log_V2ex)
		return false
	}
	return true
}

// todo: 实现
func (tou *ToucherV2ex) Sign() bool {
	resp, err := tou.client.Get(tou.signURL)
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_V2ex, err)
		return false
	}
	defer resp.Body.Close()

	//data, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	log.MyLogger.Error("%s %s", log.Log_V2ex, err)
	//	return false
	//}
	//
	//fmt.Printf("%s\n", data)
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_V2ex, err)
		return false
	}
	sel := doc.Find(".super")
	if sel == nil {
		log.MyLogger.Error("%s selection not found", log.Log_V2ex)
		return false
	}

	target, ok := sel.Attr("onclick")
	if !ok {
		log.MyLogger.Error("%s real sign url suffix not found", log.Log_V2ex)
		return false
	}
	fmt.Println(target)
	return false
}

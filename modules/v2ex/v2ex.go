package v2ex

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"sign/utils"
	"sign/utils/conf"
	"sign/utils/log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func NewV2exFromApi(conf *conf.V2exConf) (*ToucherV2ex, error) {
	if conf == nil {
		return nil, fmt.Errorf("invaild config")
	}

	tou := &ToucherV2ex{
		conf:      conf,
		loginURL:  "https://www.v2ex.com/balance",
		signURL:   "https://www.v2ex.com/mission/daily",
		verifyKey: ".balance_area,.bigger",
		client:    &http.Client{},
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	tou.client.Jar = jar
	return tou, nil
}

type ToucherV2ex struct {
	conf *conf.V2exConf

	loginURL  string
	signURL   string
	verifyKey string

	client *http.Client
}

func (tou *ToucherV2ex) Name() string {
	return tou.conf.Name
}

func (tou *ToucherV2ex) Email() string {
	return tou.conf.To
}

func (tou *ToucherV2ex) Boot() bool {
	cookies, err := utils.StrToCookies(tou.conf.Cookies, utils.V2exCookieDomain)
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

	if sel := doc.Find(tou.verifyKey); sel == nil {
		log.MyLogger.Error("%s selection is nil", log.Log_V2ex)
		return false
	}
	return true
}

func (tou *ToucherV2ex) Sign() bool {
	resp, err := tou.client.Get(tou.signURL)
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

	param, err := parseOnce(target)
	if err != nil {
		log.MyLogger.Debug("%s target is: %s", log.Log_V2ex, target)
		log.MyLogger.Error("%s %s", log.Log_V2ex, err)
		return false
	}

	realSignURL := "https://www.v2ex.com/mission/daily/redeem?once=" + param
	resp, err = tou.client.Get(realSignURL)
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_V2ex, err)
		return false
	}
	defer resp.Body.Close()

	doc, err = goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_V2ex, err)
		return false
	}

	sel = doc.Find(".gray")
	if sel == nil {
		log.MyLogger.Error("%s sign stat selection not found", log.Log_V2ex)
		return false
	}
	if sel.Text() == " &nbsp;每日登录奖励已领取" {
		return true
	}
	return true
}

// 类似: location.href = '/mission/daily/redeem?once=69089';
func parseOnce(str string) (string, error) {
	idx := strings.Index(str, "once")
	if idx < 0 {
		return "", fmt.Errorf("not found once param")
	}

	str = str[idx+5:]
	str = strings.Trim(str, ";")
	str = strings.Trim(str, "'")
	return str, nil
}
